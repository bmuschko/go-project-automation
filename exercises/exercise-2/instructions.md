# Exercise 2

In this exercise, you will declare and resolve a dependency using the `dep` package manager.

1. Install the `dep` tool. You can find installation instructions in the [user guide](https://golang.github.io/dep/docs/installation.html).
2. Manually declare a dependency on the package [github.com/pkg/errors](https://github.com/pkg/errors) in the `Gopkg.toml` file or run the relevant `dep ensure -add` command to add the dependency from the command line. What command do you need to run if you added the dependency manually?
3. Improve `utils/error.go` by replacing the code with functions from the external `errors` package.
4. Run the `dep status` command and find the `errors` package in the rendered list of dependencies. Check the constraint attribute value in the console output. Does this constraint definition have an impact on reproducibility of your project?
5. Ensure that `dep` picks a concrete version of the dependency.
6. Commit the changed files to SCM and push them to the remote repository.