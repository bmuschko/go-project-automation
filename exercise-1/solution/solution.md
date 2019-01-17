# Solution

## Using an individual code analysis tool

To check all packages beneath the current directory, run `errcheck ./...`. You should see the following error warnings:

```
$ errcheck ./...
cmd/install.go:24:29:	installCmd.MarkFlagRequired("url")
cmd/install.go:39:13:	errors.New("Currently templates can only be installed from a Git repository")
templ/gen.go:22:15:	utils.CopyDir(templatePath, targetPath)
utils/file.go:30:19:	defer srcfd.Close()
utils/file.go:35:19:	defer dstfd.Close()
```

## Using a linter

To check all packages beneath the current directory, run `golangci-lint run`. You should see the following error warnings:

```
$ golangci-lint run
cmd/install.go:38:11: `remote.GitRepo` composite literal uses unkeyed fields (govet)
		repo = &remote.GitRepo{repoUrl, templ.TemplateDir}
```

You should get a similar result if you just run with the `go vet` tool.

```
$ go vet ./...
# github.com/bmuschko/lets-gopher-exercise/cmd
cmd/install.go:38: github.com/bmuschko/lets-gopher-exercise/remote.GitRepo composite literal uses unkeyed fields
```