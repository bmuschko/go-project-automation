# Solution

1. On the command line, execute `dep ensure -add github.com/stretchr/testify`.

```
$ dep ensure -add github.com/stretchr/testify
Fetching sources...

"github.com/stretchr/testify" is not imported by your project, and has been temporarily added to Gopkg.lock and vendor/.
If you run "dep ensure" again before actually importing it, it will disappear from Gopkg.lock and vendor/.
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