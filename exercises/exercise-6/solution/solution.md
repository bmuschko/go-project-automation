# Solution

1. A properly installed binary should render help options similar to the ones below:

```
$ goreleaser -h

usage: goreleaser [<flags>] <command> [<args> ...]

Deliver Go binaries as fast and easily as possible

Flags:
  -h, --help     Show context-sensitive help (also try --help-long and --help-man).
  -v, --version  Show application version.

Commands:
  help [<command>...]
    Show help.

  init
    Generates a .goreleaser.yml file

  release* [<flags>]
    Releases the current project
```

2. After running the command you should see the file `.goreleaser.yml` in the directory. The output of the command looks as such:

```
$ goreleaser init

   • Generating .goreleaser.yml file
   • config created; please edit accordingly to your needs file=.goreleaser.yml
```

Remove the following lines from the `.goreleaser.yml` file:

```yaml
before:
  hooks:
    # you may remove this if you don't use vgo
    - go mod download
    # you may remove this if you don't need go generate
    - go generate ./...
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

   • releasing using goreleaser 0.95.0...
   • loading config file       file=.goreleaser.yml
   • RUNNING BEFORE HOOKS
   • GETTING AND VALIDATING GIT STATE
      • releasing v0.4, commit 5a4f318aecd575ea184903c6386c7d47bd4c6c4e
      • skipped                   reason=validation is disabled
   • SETTING DEFAULTS
   • SNAPSHOTING
      • skipped                   reason=not a snapshot
   • CHECKING ./DIST
   • WRITING EFFECTIVE CONFIG FILE
      • writing                   config=dist/config.yaml
   • GENERATING CHANGELOG
      • writing                   changelog=dist/CHANGELOG.md
   • LOADING ENVIRONMENT VARIABLES
      • skipped                   reason=publishing is disabled
   • BUILDING BINARIES
      • building                  binary=dist/darwin_386/lets-gopher-exercise
      • building                  binary=dist/linux_386/lets-gopher-exercise
      • building                  binary=dist/linux_amd64/lets-gopher-exercise
      • building                  binary=dist/darwin_amd64/lets-gopher-exercise
   • ARCHIVES
      • creating                  archive=dist/lets-gopher-exercise_0.4_Darwin_x86_64.tar.gz
      • creating                  archive=dist/lets-gopher-exercise_0.4_Darwin_i386.tar.gz
      • creating                  archive=dist/lets-gopher-exercise_0.4_Linux_i386.tar.gz
      • creating                  archive=dist/lets-gopher-exercise_0.4_Linux_x86_64.tar.gz
   • LINUX PACKAGES WITH NFPM
      • skipped                   reason=no output formats configured
   • SNAPCRAFT PACKAGES
      • skipped                   reason=no summary nor description were provided
   • CALCULATING CHECKSUMS
      • checksumming              file=lets-gopher-exercise_0.4_Darwin_x86_64.tar.gz
      • checksumming              file=lets-gopher-exercise_0.4_Linux_i386.tar.gz
      • checksumming              file=lets-gopher-exercise_0.4_Darwin_i386.tar.gz
      • checksumming              file=lets-gopher-exercise_0.4_Linux_x86_64.tar.gz
   • SIGNING ARTIFACTS
      • skipped                   reason=artifact signing is disabled
   • DOCKER IMAGES
      • skipped                   reason=docker section is not configured
   • PUBLISHING
      • skipped                   reason=publishing is disabled
   • release succeeded after 2.62s
```

6. First, export the `GITHUB_TOKEN` environment variable. The value can be retrieved from your GitHub account. For more information, see the page ["Creating a personal access token for the command line"](https://help.github.com/en/articles/creating-a-personal-access-token-for-the-command-line). Run the command `goreleaser release`.