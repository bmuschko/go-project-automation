# Solution

## Writing and executing tests

1. Create a new file `file_test.go` in the `utils` package.
2. Copy/paste the code from [file_standard_test.go](./utils/file_standard_test.go) into the file.
3. Run the command `go test ./...`.

## Generating code coverage metrics

1. Run the command `go test ./... -coverprofile=coverage.txt -covermode=count`.
2. Run the command `go tool cover -html=coverage.txt`. A browser window will open automatically and render the HTML report.
