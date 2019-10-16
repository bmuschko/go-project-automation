# Solution

1. A properly installed binary should render help options similar to the ones below:

```
$ goreleaser -h

Deliver Go binaries as fast and easily as possible

USAGE:
  $ goreleaser [<flags>] <command> [<args> ...]

FLAGS:
  -h, --help                    Show context-sensitive help (also try --help-long and --help-man).
      --debug                   Enable debug mode
  -f, --config=.goreleaser.yml  Load configuration from file
  -v, --version                 Show application version.

COMMANDS:
  help [<command>...]
    Show help.

  init
    Generates a .goreleaser.yml file

  check
    Checks if configuration is valid

  release [<flags>] (default)
    Releases the current project
```

2. After running the command you should see the file `.goreleaser.yml` in the directory. The output of the command looks as such:

```
$ goreleaser init

   • Generating .goreleaser.yml file
   • config created; please edit accordingly to your needs file=.goreleaser.yml
```

3. The section `archive` should contain the following code:

```yaml
archive:
  name_template: '{{ .ProjectName }}-{{ .Version }}-{{ .Os }}-{{ .Arch }}{{ if .Arm}}v{{ .Arm }}{{ end }}'
  format_overrides:
    - goos: windows
      format: zip
  files:
    - LICENSE
```

4. The section `builds` should contain the following code:

```yaml
builds:
- env:
  ...
  ldflags: -s -w -X main.version={{.Version}}
```

5. First tag the release with the the `git` command e.g. `git tag -a v0.4 -m "Version 0.4"`. The command `goreleaser release --help` renders all options needed to perform the task. To skip validation and publishing, run the command `goreleaser release --skip-validate --skip-publish`. You will want to delete the existing `dist` directory beforehand. The output of the command looks as such:

```
$ goreleaser release --skip-validate --skip-publish

   • releasing using goreleaser 0.119.0...
   • loading config file       file=.goreleaser.yml
   • RUNNING BEFORE HOOKS
      • running go mod tidy
      • running go generate ./...
   • LOADING ENVIRONMENT VARIABLES
      • pipe skipped              error=publishing is disabled
   • GETTING AND VALIDATING GIT STATE
      • releasing v0.4, commit 8e611b1791c11a4a52369e645a4027363ccb266c
      • pipe skipped              error=validation is disabled
   • PARSING TAG
   • SETTING DEFAULTS
      • LOADING ENVIRONMENT VARIABLES
      • SNAPSHOTING
      • GITHUB/GITLAB/GITEA RELEASES
      • PROJECT NAME
      • BUILDING BINARIES
      • ARCHIVES
      • LINUX PACKAGES WITH NFPM
      • SNAPCRAFT PACKAGES
      • CALCULATING CHECKSUMS
      • SIGNING ARTIFACTS
      • DOCKER IMAGES
      • ARTIFACTORY
      • S3
      • BLOB
      • HOMEBREW TAP FORMULA
         • optimistically guessing `brew[0].installs`, double check
      • SCOOP MANIFEST
   • SNAPSHOTING
      • pipe skipped              error=not a snapshot
   • CHECKING ./DIST
   • WRITING EFFECTIVE CONFIG FILE
      • writing                   config=dist/config.yaml
   • GENERATING CHANGELOG
      • writing                   changelog=dist/CHANGELOG.md
   • BUILDING BINARIES
      • building                  binary=dist/lets-gopher-exercise_freebsd_arm_6/lets-gopher-exercise
      • building                  binary=dist/lets-gopher-exercise_linux_arm64/lets-gopher-exercise
      • building                  binary=dist/lets-gopher-exercise_linux_386/lets-gopher-exercise
      • building                  binary=dist/lets-gopher-exercise_linux_amd64/lets-gopher-exercise
      • building                  binary=dist/lets-gopher-exercise_netbsd_amd64/lets-gopher-exercise
      • building                  binary=dist/lets-gopher-exercise_windows_amd64/lets-gopher-exercise.exe
      • building                  binary=dist/lets-gopher-exercise_netbsd_arm_6/lets-gopher-exercise
      • building                  binary=dist/lets-gopher-exercise_netbsd_386/lets-gopher-exercise
      • building                  binary=dist/lets-gopher-exercise_darwin_amd64/lets-gopher-exercise
      • building                  binary=dist/lets-gopher-exercise_openbsd_arm_6/lets-gopher-exercise
      • building                  binary=dist/lets-gopher-exercise_openbsd_386/lets-gopher-exercise
      • building                  binary=dist/lets-gopher-exercise_darwin_386/lets-gopher-exercise
      • building                  binary=dist/lets-gopher-exercise_openbsd_amd64/lets-gopher-exercise
      • building                  binary=dist/lets-gopher-exercise_freebsd_386/lets-gopher-exercise
      • building                  binary=dist/lets-gopher-exercise_freebsd_amd64/lets-gopher-exercise
      • building                  binary=dist/lets-gopher-exercise_linux_arm_6/lets-gopher-exercise
      • building                  binary=dist/lets-gopher-exercise_windows_386/lets-gopher-exercise.exe
   • ARCHIVES
      • creating                  archive=dist/lets-gopher-exercise_0.4_netbsd_x86_64.tar.gz
      • creating                  archive=dist/lets-gopher-exercise_0.4_netbsd_i386.tar.gz
      • creating                  archive=dist/lets-gopher-exercise_0.4_freebsd_armv6.tar.gz
      • creating                  archive=dist/lets-gopher-exercise_0.4_Linux_arm64.tar.gz
      • creating                  archive=dist/lets-gopher-exercise_0.4_openbsd_x86_64.tar.gz
      • creating                  archive=dist/lets-gopher-exercise_0.4_openbsd_i386.tar.gz
      • creating                  archive=dist/lets-gopher-exercise_0.4_Darwin_i386.tar.gz
      • creating                  archive=dist/lets-gopher-exercise_0.4_Linux_i386.tar.gz
      • creating                  archive=dist/lets-gopher-exercise_0.4_freebsd_x86_64.tar.gz
      • creating                  archive=dist/lets-gopher-exercise_0.4_openbsd_armv6.tar.gz
      • creating                  archive=dist/lets-gopher-exercise_0.4_netbsd_armv6.tar.gz
      • creating                  archive=dist/lets-gopher-exercise_0.4_Linux_x86_64.tar.gz
      • creating                  archive=dist/lets-gopher-exercise_0.4_Windows_x86_64.tar.gz
      • creating                  archive=dist/lets-gopher-exercise_0.4_Darwin_x86_64.tar.gz
      • creating                  archive=dist/lets-gopher-exercise_0.4_Linux_armv6.tar.gz
      • creating                  archive=dist/lets-gopher-exercise_0.4_Windows_i386.tar.gz
      • creating                  archive=dist/lets-gopher-exercise_0.4_freebsd_i386.tar.gz
   • LINUX PACKAGES WITH NFPM
      • pipe skipped              error=no output formats configured
   • SNAPCRAFT PACKAGES
      • pipe skipped              error=no summary nor description were provided
   • CALCULATING CHECKSUMS
      • checksumming              file=lets-gopher-exercise_0.4_netbsd_i386.tar.gz
      • checksumming              file=lets-gopher-exercise_0.4_openbsd_armv6.tar.gz
      • checksumming              file=lets-gopher-exercise_0.4_freebsd_i386.tar.gz
      • checksumming              file=lets-gopher-exercise_0.4_Linux_armv6.tar.gz
      • checksumming              file=lets-gopher-exercise_0.4_freebsd_armv6.tar.gz
      • checksumming              file=lets-gopher-exercise_0.4_Linux_arm64.tar.gz
      • checksumming              file=lets-gopher-exercise_0.4_netbsd_x86_64.tar.gz
      • checksumming              file=lets-gopher-exercise_0.4_openbsd_i386.tar.gz
      • checksumming              file=lets-gopher-exercise_0.4_Darwin_i386.tar.gz
      • checksumming              file=lets-gopher-exercise_0.4_Linux_i386.tar.gz
      • checksumming              file=lets-gopher-exercise_0.4_Windows_x86_64.tar.gz
      • checksumming              file=lets-gopher-exercise_0.4_openbsd_x86_64.tar.gz
      • checksumming              file=lets-gopher-exercise_0.4_Windows_i386.tar.gz
      • checksumming              file=lets-gopher-exercise_0.4_Linux_x86_64.tar.gz
      • checksumming              file=lets-gopher-exercise_0.4_freebsd_x86_64.tar.gz
      • checksumming              file=lets-gopher-exercise_0.4_netbsd_armv6.tar.gz
      • checksumming              file=lets-gopher-exercise_0.4_Darwin_x86_64.tar.gz
   • SIGNING ARTIFACTS
      • pipe skipped              error=artifact signing is disabled
   • DOCKER IMAGES
      • pipe skipped              error=docker section is not configured
   • PUBLISHING
      • pipe skipped              error=publishing is disabled
   • release succeeded after 13.50s
```

6. First, export the `GITHUB_TOKEN` environment variable. The value can be retrieved from your GitHub account. For more information, see the page ["Creating a personal access token for the command line"](https://help.github.com/en/articles/creating-a-personal-access-token-for-the-command-line). Run the command `goreleaser release`.