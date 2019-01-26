# Demo 8

Demonstrates how to build a pipeline with Travis CI.

The provided `.travis.yml` files contain an example Travis CI pipeline definition. They use `golangci-lint` for static code analysis and `GoReleaser` for publishing the binary artifacts to GitHub releases.

## Samples

You can find pipeline definitions for `dep` and Go modules:

* Using [dep](./dep/.travis.yml)
* Using [Go modules](./mod/.travis.yml)

## Pipeline steps

Each pipeline is comprised of multiple build steps:

1. Compile packages and dependencies.
2. Run unit tests and publish code coverage result to Codecov.
3. Run code quality analysis using `golangci-lint`.
4. Build and release the binaries to GitHub if the current commit points to Git tag.