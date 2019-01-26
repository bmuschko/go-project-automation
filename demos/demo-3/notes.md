# Demo 3

Demonstrates the purpose and usage of the tool [Golint](https://github.com/golang/lint).

## Features

* Check for proper code styling and not correctness
* Makes suggestions on improving code
* Possible to get false positives
* Typical example: `if`/`else` conditionals where `else` is unnecessary

## Usage

```
$ golint
main.go:13:9: if block ends with a return statement, so drop this else and outdent its block
```