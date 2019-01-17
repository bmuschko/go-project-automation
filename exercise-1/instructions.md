# Exercise 1

In this exercise, you will learn how to use code analysis tools to improve the code and find potential errors.

## Using an individual code analysis tool

1. Install the tool [`errcheck`](github.com/kisielk/errcheck
) using the command `go get -u github.com/kisielk/errcheck`.
2. Run the tool on the project. The tool should report multiple errors. Fix the code based on the suggestions. Ensure that the tool stops reporting the issues by re-running it.

## Using a linter

1. Install the tool `golangci-lint`. You can find the installation instructions on the [web page](https://github.com/golangci/golangci-lint#install). Ensure that the tool is installed properly by running `golangci-lint help`.
2. Run the tool on the project with the command `golangci-lint run`. The tool should report multiple errors. Fix them the code based on the suggestions. Ensure that the tool stops reporting the issues after re-running it.
3. (Optional) Try out integrating and running the tool from VSCode or GoLand if you one of them installed already.



