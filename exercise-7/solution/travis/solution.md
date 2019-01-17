# Solution

The provided [`.travis.yml`](./.travis.yml) contains an example Travis CI pipeline definition. It uses `dep` for package management and `golangci-lint` for static code analysis.

The pipeline is comprised of multiple build steps:

1. Compile packages and dependencies.
2. Run unit tests and publish code coverage result to Codecov.
3. Run code quality analysis using `golangci-lint`.
4. Build and release the binaries to GitHub if the current commit points to Git tag.