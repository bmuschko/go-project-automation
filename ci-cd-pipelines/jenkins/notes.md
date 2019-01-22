# Jenkins pipeline

Jenkins can automatically bootstrap the Go runtime with the help pf the [Go plugin](https://wiki.jenkins.io/display/JENKINS/Go+Plugin). It automatically installs a pre-defined version of Go and exports the `GOROOT` and `GOPATH` environment variable.

## Samples

You can find pipeline definitions for `dep` and Go modules:

* Using [dep](./dep/Jenkinsfile): At the moment only [scripted pipelines](https://jenkins.io/doc/book/pipeline/syntax/#scripted-pipeline) are supported. For more information see [JENKINS-55630](https://issues.jenkins-ci.org/browse/JENKINS-55630).
* Using [Go modules](./mod/Jenkinsfile): This scenario does not require the checked out project to live in `$GOPATH/src`. For that reason a [declarative pipeline](https://jenkins.io/doc/book/pipeline/syntax/#declarative-pipeline) definition can be used which makes the setup much easier.

## Pipeline steps

The pipeline is comprised of multiple build steps:

1. Compile packages and dependencies.
2. Run unit tests and publish code coverage result to Codecov.
3. Run code quality analysis using `golangci-lint`.
4. Build and release the binaries to GitHub if the current commit points to Git tag.

## Managing the environment

The pipeline references environment variables, credentials and tools that need to be set up manually.

1. The Go version needs to be configured under _Manage Jenkins > Global Tool Configuration > Go_.
1. The environment variable `CODECOV_TOKEN` needs to provide the Codecov token. Environment variables can be configured under _Manage Jenkins > Configure System > Environment variables_.
2. The credential `GITHUB_TOKEN` needs to provide a valid GitHub token. Credentials can be configured under _Manage Jenkins > Configure Credentials_.