# Solution

1. On the command line, execute `go get github.com/stretchr/testify`.

```
$ go get github.com/stretchr/testify
go: finding github.com/stretchr/testify v1.4.0
go: downloading github.com/stretchr/testify v1.4.0
go: extracting github.com/stretchr/testify v1.4.0
go: downloading github.com/stretchr/objx v0.1.0
go: extracting github.com/stretchr/objx v0.1.0
go: finding github.com/stretchr/objx v0.1.0
```

2. Copy/paste the code from [file_testify_test.go](./utils/file_testify_test.go) and overwrite the existing code in `file_test.go`.
3. Run the command `go test ./...`.

```
$ go test ./... -v
=== RUN   TestCreateDirForNonExistentDirectoryWithTestingPackage
--- PASS: TestCreateDirForNonExistentDirectoryWithTestingPackage (0.00s)
=== RUN   TestCreateDirForExistentDirectoryWithTestingPackage
--- PASS: TestCreateDirForExistentDirectoryWithTestingPackage (0.00s)
PASS
ok  	github.com/bmuschko/lets-gopher-exercise/utils	0.007s
```

_Note:_ The file [file_ginkgo_test.go](./utils/file_ginkgo_test.go) contains test code that uses the Ginkgo package.