# Exercise 1

In this exercise, you will learn how to use code analysis tools to improve the code and find potential errors.

## Using an individual code analysis tool

1. Install the tool [`errcheck`](github.com/kisielk/errcheck
) using the command `go get -u github.com/kisielk/errcheck`.
2. Run the tool on the project. The tool should report multiple errors. Fix the code based on the indicated file and line number.
3. Ensure that the `errcheck` stops reporting the issues by re-running the command.

## Using a linter

1. Install the tool `golangci-lint`. You can find the installation instructions on the [web page](https://github.com/golangci/golangci-lint#install). Ensure that the tool is installed properly by running `golangci-lint help`.
2. Run the tool on the project with the command `golangci-lint run`. The tool should report multiple errors. Fix the code based on the suggestions
3. Ensure that the tool stops reporting the issues after re-running it.
4. (Optional) Try out integrating and running the tool from VSCode or GoLand if you have one of them installed on your machine.

## Using Go Report Card

1. Open [Go Report Card](https://goreportcard.com/) in a browser of your choice.
2. Enter the `go get` path `github.com/<your-github-account>/lets-gopher-exercise`.
3. Click the button _Generate Report_. Explore the report metrics.
4. (Optional) Implement some of the suggestions and push the changes to your repository. Do the metrics and/or suggestions change after reloading the report?



