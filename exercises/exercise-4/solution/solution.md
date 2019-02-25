# Solution

1. On the command line, execute `dep ensure -add github.com/stretchr/testify`.
2. Copy/paste the code from [file_testify_test.go](./utils/file_testify_test.go) and overwrite the existing code in `file_test.go`.
3. Run the command `go test ./...`.

_Note:_ The file [file_ginkgo_test.go](./utils/file_ginkgo_test.go) contains test code that uses the Ginkgo package.