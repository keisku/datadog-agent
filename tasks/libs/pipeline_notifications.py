import json
import os
import subprocess
from collections import defaultdict

from .common.gitlab import Gitlab


def get_failed_jobs(project_name, pipeline_id):
    gitlab = Gitlab()

    jobs = gitlab.all_jobs(project_name, pipeline_id)

    # Get instances of failed jobs
    failed_jobs = {job["name"]: [] for job in jobs if job["status"] == "failed"}

    # Group jobs per name
    for job in jobs:
        if job["name"] in failed_jobs:
            failed_jobs[job["name"]].append(job)

    # There, we now have the following map:
    # job name -> list of jobs with that name, including at least one failed job

    final_failed_jobs = []
    for job_name, jobs in failed_jobs.items():
        # We sort each list per creation date
        jobs.sort(key=lambda x: x["created_at"])
        # Check the final job in the list: it contains the current status of the job
        final_status = {
            "name": job_name,
            "id": jobs[-1]["id"],
            "stage": jobs[-1]["stage"],
            "status": jobs[-1]["status"],
            "allow_failure": jobs[-1]["allow_failure"],
            "url": jobs[-1]["web_url"],
            "retry_summary": [job["status"] for job in jobs],
        }
        final_failed_jobs.append(final_status)

    return final_failed_jobs


class Test:
    PACKAGE_PREFIX = "github.com/DataDog/datadog-agent/"

    def __init__(self, owners, name, package):
        self.name = name
        self.package = self.__removeprefix(package)
        self.owners = self.__get_owners(owners, package)

    def __removeprefix(self, package):
        return package[len(self.PACKAGE_PREFIX) :]

    def __get_owners(self, OWNERS, package):
        owners = OWNERS.of(self.__removeprefix(package))
        return [name for (kind, name) in owners if kind == "TEAM"]


def read_owners(owners_file):
    from codeowners import CodeOwners

    with open(owners_file, 'r') as f:
        return CodeOwners(f.read())


def get_failed_tests(project_name, job, owners_file=".github/CODEOWNERS"):
    gitlab = Gitlab()
    owners = read_owners(owners_file)
    test_output = gitlab.artifact(project_name, job["id"])
    for line in test_output.splitlines():
        try:
            json_test = json.loads(line)
            if "message" in json_test:
                continue  # Failed request
            if 'Test' in json_test and json_test["Action"] == "fail":
                yield Test(owners, json_test['Test'], json_test['Package'])
        except Exception as e:
            print("WARN: parsing '{}' failed: {}".format(line, e))


def find_job_owners(failed_jobs, owners_file=".gitlab/JOBOWNERS"):
    owners = read_owners(owners_file)
    owners_to_notify = defaultdict(list)

    for job in failed_jobs:
        # Exclude jobs that were retried and succeeded
        # Also exclude jobs allowed to fail
        if job["status"] == "failed" and not job["allow_failure"]:
            job_owners = owners.of(job["name"])
            # job_owners is a list of tuples containing the type of owner (eg. USERNAME, TEAM) and the name of the owner
            # eg. [('TEAM', '@DataDog/agent-platform')]

            for owner in job_owners:
                owners_to_notify[owner[1]].append(job)

    return owners_to_notify


def base_message(header):
    return """{header} pipeline <{pipeline_url}|{pipeline_id}> for {commit_ref_name} failed.
{commit_title} (<{commit_url}|{commit_short_sha}>) by {author}""".format(
        header=header,
        pipeline_url=os.getenv("CI_PIPELINE_URL"),
        pipeline_id=os.getenv("CI_PIPELINE_ID"),
        commit_ref_name=os.getenv("CI_COMMIT_REF_NAME"),
        commit_title=os.getenv("CI_COMMIT_TITLE"),
        commit_url="{project_url}/commit/{commit_sha}".format(
            project_url=os.getenv("CI_PROJECT_URL"), commit_sha=os.getenv("CI_COMMIT_SHA")
        ),
        commit_short_sha=os.getenv("CI_COMMIT_SHORT_SHA"),
        author=get_git_author(),
    )


def prepare_global_failure_message(header, failed_jobs):
    message = base_message(header)

    message += "\nFailed jobs:"
    for job in failed_jobs:
        # Exclude jobs that were retried and succeeded
        # Also exclude jobs allowed to fail
        if job["status"] == "failed" and not job["allow_failure"]:
            message += "\n - <{url}|{name}> (stage: {stage}, after {retries} retries)".format(
                url=job["url"], name=job["name"], stage=job["stage"], retries=len(job["retry_summary"]) - 1
            )

    return message


def prepare_team_failure_message(header, failed_jobs):
    message = base_message(header)

    message += "\nFailed jobs you own:"
    for job in failed_jobs:
        message += "\n - <{url}|{name}> (stage: {stage}, after {retries} retries)".format(
            url=job["url"], name=job["name"], stage=job["stage"], retries=len(job["retry_summary"]) - 1
        )

    return message


def prepare_test_failure_section(failed_tests):
    # failed_tests : dict[Test, list[Job]]

    section = "\nFailed unit tests you own:"
    for test, jobs in failed_tests:
        MAX_SHOW = 2
        job_list = ", ".join("<{}|{}>".format(job["url"], job["name"]) for job in jobs[:MAX_SHOW])
        if len(jobs) > MAX_SHOW:
            job_list += "and {} more".format(len(jobs) - MAX_SHOW)
        section += "\n - `{}` from package `{}`(in {})".format(test.name, test.package, job_list)

    return section


def get_git_author():
    return (
        subprocess.check_output(["git", "show", "-s", "--format='%an'", "HEAD"])
        .decode('utf-8')
        .strip()
        .replace("'", "")
    )


def send_slack_message(recipient, message):
    subprocess.run(["postmessage", recipient, message], check=True)