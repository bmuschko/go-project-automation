# Demo 7

Demonstrates building and publishing binaries to Artifactory with GoReleaser.

* Download open source version of Artifactory at https://jfrog.com/open-source/.
* Start Artifactory with via `<installation-directory>/artifactory-oss-6.6.1/bin/artifactory.sh`.
* Set environment variable `ARTIFACTORY_PRODUCTION_USERNAME`: `export ARTIFACTORY_PRODUCTION_USERNAME=admin`
* Set environment variable `ARTIFACTORY_PRODUCTION_SECRET`: `export ARTIFACTORY_PRODUCTION_SECRET=admin`
* Configure `GoReleaser` to publish to Artifactory.
* Publish a version with `GoReleaser` from project directory: `goreleaser`
* Open URL `http://localhost:8081/artifactory/example-repo-local` in the browser to check released binaries.