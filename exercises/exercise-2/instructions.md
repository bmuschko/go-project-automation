# Exercise 2

In this exercise, you will declare and resolve a dependency using Go Modules.

1. Make sure the environment variable `GO111MODULES=on` has been set. You can find additional information in Go Modules in the [Wiki](https://github.com/golang/go/wiki/Modules).
2. Add a dependency on the package [github.com/pkg/errors](https://github.com/pkg/errors) using the `go get` command. Inspect the `go.mod` file and find the dependency.
3. Improve `utils/error.go` by replacing the code with functions from the external `errors` package.
4. Run the `go mod tidy` command and inspect the `go.mod` file. Do you notice a difference?
5. Run the `go mod graph` command and find the `errors` package in the rendered list of dependencies. Check the resolved version value in the console output.
6. Commit the changed files to SCM and push them to the remote repository.