name "python3"

default_version "3.8.10"

if ohai["platform"] != "windows"
  dependency "libffi"
  dependency "ncurses"
  dependency "zlib"
  dependency "openssl"
  dependency "pkg-config"
  dependency "bzip2"
  dependency "libsqlite3"
  dependency "liblzma"
  dependency "libyaml"

  source :url => "https://python.org/ftp/python/#{version}/Python-#{version}.tgz",
         :sha256 => "b37ac74d2cbad2590e7cd0dd2b3826c29afe89a734090a87bf8c03c45066cb65"

  relative_path "Python-#{version}"

  python_configure = ["./configure",
                      "--prefix=#{install_dir}/embedded",
                      "--with-ssl=#{install_dir}/embedded",
                      "--with-ensurepip=no"] # pip is installed separately by its own software def

  if mac_os_x?
    python_configure.push("--enable-ipv6",
                          "--with-universal-archs=intel",
                          "--enable-shared")
  elsif linux?
    python_configure.push("--enable-shared",
                          "--enable-ipv6")
  elsif aix?
    # something here...
  end

  python_configure.push("--with-dbmliborder=")

  build do
    ship_license "PSFL"

    env = case ohai["platform"]
          when "aix"
            aix_env
          else
            {
              "CFLAGS" => "-I#{install_dir}/embedded/include -O2 -g -pipe",
              "LDFLAGS" => "-Wl,-rpath,#{install_dir}/embedded/lib -L#{install_dir}/embedded/lib",
              "PKG_CONFIG" => "#{install_dir}/embedded/bin/pkg-config",
              "PKG_CONFIG_PATH" => "#{install_dir}/embedded/lib/pkgconfig"
            }
          end
    command python_configure.join(" "), :env => env
    command "make -j #{workers}", :env => env
    command "make install", :env => env
    delete "#{install_dir}/embedded/lib/python3.8/test"

    # There exists no configure flag to tell Python to not compile readline support :(
    major, minor, bugfix = version.split(".")
    block do
      FileUtils.rm_f(Dir.glob("#{install_dir}/embedded/lib/python#{major}.#{minor}/lib-dynload/readline.*"))
      FileUtils.rm_f(Dir.glob("#{install_dir}/embedded/lib/python#{major}.#{minor}/distutils/command/wininst-*.exe"))
    end
  end

else
  dependency "vc_redist_14"

  if windows_arch_i386?
    dependency "vc_ucrt_redist"

    source :url => "https://dd-agent-omnibus.s3.amazonaws.com/test-python-windows-#{version}-x86.zip",
            :sha256 => "b0c1b9873ed4521567bc0e778394eb7a8dbcf69c3775039d544539b4d1de981a"
  else

    # note that startring with 3.7.3 on Windows, the zip should be created without the built-in pip
    source :url => "https://dd-agent-omnibus.s3.amazonaws.com/test-python-windows-#{version}-amd64.zip",
         :sha256 => "9ee2e897542f9635b3f378998e495d1c6a815f4d6045a27ccf9d2c0ae221686a"

  end
  vcrt140_root = "#{Omnibus::Config.source_dir()}/vc_redist_140/expanded"
  build do
    command "XCOPY /YEHIR *.* \"#{windows_safe_path(python_3_embedded)}\""
    command "copy /y \"#{windows_safe_path(vcrt140_root)}\\*.dll\" \"#{windows_safe_path(python_3_embedded)}\""
  end
end
