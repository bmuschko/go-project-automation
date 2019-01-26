# Demo 6

Demonstrates the purpose and usage of [Go report card](https://goreportcard.com/).

## Usage of web application

* Open Go report card
* Enter `github.com/bmuschko/lets-gopher-exercise` into input field
* Click "Generate Report" button

## Usage of CLI application

```
$ go get -u github.com/gojp/goreportcard/cmd/goreportcard-cli
$ cd $GOPATH/src/github.com/bmuschko/lets-gopher-exercise
$ goreportcard-cli
2019/01/26 13:35:28 ERROR: (ineffassign) exit status 2
2019/01/26 13:35:28 ERROR: (misspell) exit status 2
Grade: A+ (92.4%)
Files: 681
Issues: 81
gofmt: 99%
go_vet: 99%
golint: 98%
gocyclo: 89%
ineffassign: 0%
license: 100%
misspell: 0%
```