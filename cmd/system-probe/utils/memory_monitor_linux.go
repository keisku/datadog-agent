// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2020 Datadog, Inc.

// +build linux

package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"sort"
	"strconv"
	"strings"

	"github.com/alecthomas/units"

	"github.com/DataDog/datadog-agent/pkg/util"
	"github.com/DataDog/datadog-agent/pkg/util/log"
)

// MemoryMonitor monitors memory cgroup usage
type MemoryMonitor = util.MemoryController

const maxProfileCount = 10

func getActionCallback(action string) (func(), string, error) {
	switch action {
	case "gc":
		return runtime.GC, "garbage collector", nil
	case "log":
		return func() {}, "nothing", nil
	case "profile":
		return func() {
			tmpDir := os.TempDir()
			tmpFiles, err := ioutil.ReadDir(tmpDir)
			if err != nil {
				log.Errorf("Failed to list old memory profiles: %s", err)
			} else {
				var oldProfiles []os.FileInfo
				for _, tmpFile := range tmpFiles {
					if strings.HasPrefix(tmpFile.Name(), "memcg-pprof-heap") {
						oldProfiles = append(oldProfiles, tmpFile)
					}
				}

				sort.Slice(oldProfiles, func(i, j int) bool {
					return oldProfiles[i].ModTime().After(oldProfiles[j].ModTime())
				})

				for i := len(oldProfiles) - 1; i >= 0 && i >= maxProfileCount-1; i-- {
					os.Remove(filepath.Join(tmpDir, oldProfiles[i].Name()))
					oldProfiles = oldProfiles[:i]
				}
			}

			memProfile, err := ioutil.TempFile(tmpDir, "memcg-pprof-heap")
			if err != nil {
				log.Errorf("Failed to generate memory profile: %s", err)
				return
			}
			defer memProfile.Close()

			if err := pprof.WriteHeapProfile(memProfile); err != nil {
				log.Errorf("Failed to generate memory profile: %s", err)
				return
			}
		}, "heap profile", nil
	default:
		return nil, "", fmt.Errorf("unknown memory controller action '%s'", action)
	}
}

// NewMemoryMonitor instantiates a new memory monitor
func NewMemoryMonitor(pressureLevels map[string]string, thresholds map[string]string) (*MemoryMonitor, error) {
	var memoryMonitors []util.MemoryMonitor

	for pressureLevel, action := range pressureLevels {
		actionCallback, name, err := getActionCallback(action)
		if err != nil {
			return nil, err
		}

		log.Infof("New memory pressure monitor on level %s with action %s", pressureLevel, name)
		memoryMonitors = append(memoryMonitors, util.MemoryPressureMonitor(func() {
			log.Infof("Memory pressure reached level '%s', triggering %s", pressureLevel, name)
			actionCallback()
		}, pressureLevel))
	}

	for threshold, action := range thresholds {
		actionCallback, name, err := getActionCallback(action)
		if err != nil {
			return nil, err
		}

		monitorCallback := func() {
			log.Infof("Memory pressure above %s threshold, triggering %s", threshold, name)
			actionCallback()
		}

		var memoryMonitor util.MemoryMonitor
		threshold = strings.TrimSpace(threshold)
		if strings.HasSuffix(threshold, "%") {
			percentage, err := strconv.Atoi(strings.TrimSuffix(threshold, "%"))
			if err != nil {
				return nil, fmt.Errorf("invalid memory threshold '%s': %w", threshold, err)
			}

			memoryMonitor = util.MemoryPercentageThresholdMonitor(monitorCallback, uint64(percentage), false)
		} else {
			size, err := units.ParseBase2Bytes(strings.ToUpper(threshold))
			if err != nil {
				return nil, fmt.Errorf("invalid memory threshold '%s': %w", threshold, err)
			}

			memoryMonitor = util.MemoryThresholdMonitor(monitorCallback, uint64(size), false)
		}

		log.Infof("New memory threshold monitor on level %s with action %s", threshold, name)
		memoryMonitors = append(memoryMonitors, memoryMonitor)
	}

	return util.NewMemoryController(memoryMonitors...)
}