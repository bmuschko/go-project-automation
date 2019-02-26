# Demo 7

Demonstrates building and publishing binaries to Artifactory with GoReleaser.

1. Download open source version of Artifactory at https://jfrog.com/open-source/.
2. Start Artifactory with via `<installation-directory>/artifactory-oss-6.6.1/bin/artifactory.sh`.
3. (Optional) Set environment variable `ARTIFACTORY_PRODUCTION_USERNAME`: `export ARTIFACTORY_PRODUCTION_USERNAME=admin`
4. Set environment variable `ARTIFACTORY_PRODUCTION_SECRET`: `export ARTIFACTORY_PRODUCTION_SECRET=admin`
5. Add the [GoReleaser configuration file](../.goreleaser.yml) to the root directory of the example project.
6. Tag the current release (e.g. with version 0.4) by running the command `git tag -a v0.4 -m "Version 0.4"`.
7. Publish the binary files with the command `goreleaser release`. You may have to delete existing binaries if the `dist` directory already exists. Use the command line option `--skip-validate` if your local repository checkout still has uncommitted changes.

```
$ goreleaser release --skip-validate

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
      • skipped                   reason=release pipe is disabled
   • BUILDING BINARIES
      • building                  binary=dist/linux_386/lets-gopher-exercise
      • building                  binary=dist/linux_amd64/lets-gopher-exercise
      • building                  binary=dist/windows_amd64/lets-gopher-exercise.exe
      • building                  binary=dist/freebsd_arm_6/lets-gopher-exercise
      • building                  binary=dist/netbsd_386/lets-gopher-exercise
      • building                  binary=dist/openbsd_386/lets-gopher-exercise
      • building                  binary=dist/netbsd_amd64/lets-gopher-exercise
      • building                  binary=dist/netbsd_arm_6/lets-gopher-exercise
      • building                  binary=dist/openbsd_amd64/lets-gopher-exercise
      • building                  binary=dist/freebsd_386/lets-gopher-exercise
      • building                  binary=dist/darwin_386/lets-gopher-exercise
      • building                  binary=dist/linux_arm_6/lets-gopher-exercise
      • building                  binary=dist/linux_arm64/lets-gopher-exercise
      • building                  binary=dist/freebsd_amd64/lets-gopher-exercise
      • building                  binary=dist/darwin_amd64/lets-gopher-exercise
      • building                  binary=dist/windows_386/lets-gopher-exercise.exe
      • building                  binary=dist/openbsd_arm_6/lets-gopher-exercise
   • ARCHIVES
      • creating                  archive=dist/lets-gopher-exercise-0.4-openbsd-386.tar.gz
      • creating                  archive=dist/lets-gopher-exercise-0.4-linux-arm64.tar.gz
      • creating                  archive=dist/lets-gopher-exercise-0.4-darwin-amd64.tar.gz
      • creating                  archive=dist/lets-gopher-exercise-0.4-netbsd-armv6.tar.gz
      • creating                  archive=dist/lets-gopher-exercise-0.4-freebsd-amd64.tar.gz
      • creating                  archive=dist/lets-gopher-exercise-0.4-freebsd-armv6.tar.gz
      • creating                  archive=dist/lets-gopher-exercise-0.4-windows-amd64.zip
      • creating                  archive=dist/lets-gopher-exercise-0.4-netbsd-386.tar.gz
      • creating                  archive=dist/lets-gopher-exercise-0.4-windows-386.zip
      • creating                  archive=dist/lets-gopher-exercise-0.4-netbsd-amd64.tar.gz
      • creating                  archive=dist/lets-gopher-exercise-0.4-openbsd-amd64.tar.gz
      • creating                  archive=dist/lets-gopher-exercise-0.4-linux-armv6.tar.gz
      • creating                  archive=dist/lets-gopher-exercise-0.4-linux-386.tar.gz
      • creating                  archive=dist/lets-gopher-exercise-0.4-freebsd-386.tar.gz
      • creating                  archive=dist/lets-gopher-exercise-0.4-openbsd-armv6.tar.gz
      • creating                  archive=dist/lets-gopher-exercise-0.4-linux-amd64.tar.gz
      • creating                  archive=dist/lets-gopher-exercise-0.4-darwin-386.tar.gz
   • LINUX PACKAGES WITH NFPM
      • skipped                   reason=no output formats configured
   • SNAPCRAFT PACKAGES
      • skipped                   reason=no summary nor description were provided
   • CALCULATING CHECKSUMS
      • checksumming              file=lets-gopher-exercise-0.4-windows-amd64.zip
      • checksumming              file=lets-gopher-exercise-0.4-windows-386.zip
      • checksumming              file=lets-gopher-exercise-0.4-linux-386.tar.gz
      • checksumming              file=lets-gopher-exercise-0.4-darwin-amd64.tar.gz
      • checksumming              file=lets-gopher-exercise-0.4-netbsd-armv6.tar.gz
      • checksumming              file=lets-gopher-exercise-0.4-openbsd-armv6.tar.gz
      • checksumming              file=lets-gopher-exercise-0.4-netbsd-amd64.tar.gz
      • checksumming              file=lets-gopher-exercise-0.4-openbsd-386.tar.gz
      • checksumming              file=lets-gopher-exercise-0.4-freebsd-armv6.tar.gz
      • checksumming              file=lets-gopher-exercise-0.4-openbsd-amd64.tar.gz
      • checksumming              file=lets-gopher-exercise-0.4-netbsd-386.tar.gz
      • checksumming              file=lets-gopher-exercise-0.4-linux-amd64.tar.gz
      • checksumming              file=lets-gopher-exercise-0.4-linux-armv6.tar.gz
      • checksumming              file=lets-gopher-exercise-0.4-linux-arm64.tar.gz
      • checksumming              file=lets-gopher-exercise-0.4-freebsd-386.tar.gz
      • checksumming              file=lets-gopher-exercise-0.4-freebsd-amd64.tar.gz
      • checksumming              file=lets-gopher-exercise-0.4-darwin-386.tar.gz
   • SIGNING ARTIFACTS
      • skipped                   reason=artifact signing is disabled
   • DOCKER IMAGES
      • skipped                   reason=docker section is not configured
   • PUBLISHING
      • S3
      • skipped                   reason=s3 section is not configured
      • releasing with HTTP PUT
      • skipped                   reason=put section is not configured
      • Artifactory
      • skipped                   reason=artifactory section 'production' is not configured properly (missing ARTIFACTORY_PRODUCTION_SECRET environment variable)
      • Docker images
      • Snapcraft Packages
      • GitHub Releases
      • skipped                   reason=release pipe is disabled
      • homebrew tap formula
      • skipped                   reason=brew section is not configured
      • scoop manifest
      • skipped                   reason=scoop section is not configured
   • release succeeded after 11.20s
```

8. Verify the release on Artifactory by opening a browser with the URL `http://localhost:8081/artifactory/webapp/#/artifacts/browse/tree/General/example-repo-local/lets-gopher-exercise/0.4`. Open the folder to see all binary files.

![Artifactory repository browser](./images/artifactory-repo-browser.png)