# Demo 2

Demonstrates the purpose and usage of the tool [Go vet](https://golang.org/cmd/vet/).

## Features

* Starts where the compiler ends
* Examines Go source code and reports suspicious constructs
* Mistakes could easily be missed during code reviews or unit testing
* Typical example: Missing arguments for `Printf` calls

## Usage

```
$ go vet
# _/Users/bmuschko/dev/projects/go-project-automation/demos/demo-2 [_/Users/bmuschko/dev/projects/go-project-automation/demos/demo-2.test]
./main.go:7:2: age declared and not used
```