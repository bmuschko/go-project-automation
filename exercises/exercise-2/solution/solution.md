# Solution

1. You can check if the environment variable is set by using the `echo` command e.g. `echo $GO111MODULES` (Linux, MacOS) or `echo %GO111MODULES%` (Windows).
2. Run the command `go get github.com/pkg/errors` to resolve the package.

```
$ go get github.com/pkg/errors
go: finding github.com/pkg/errors v0.8.1
go: downloading github.com/pkg/errors v0.8.1
go: extracting github.com/pkg/errors v0.8.1
```

The contents of `go.mod` look as follows. You can see that the `errors` package has been marked as indirect. The indirect comment indicates a dependency is not used directly by this module, only indirectly by other module dependencies.

```
module github.com/bmuschko/lets-gopher-exercise

require (
	github.com/Netflix/go-expect v0.0.0-20180928190340-9d1f4485533b // indirect
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/hinshun/vt10x v0.0.0-20180809195222-d55458df857c // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/mattn/go-colorable v0.0.9 // indirect
	github.com/mattn/go-isatty v0.0.4 // indirect
	github.com/mgutz/ansi v0.0.0-20170206155736-9520e82c474b // indirect
	github.com/mitchellh/go-homedir v1.0.0
	github.com/pkg/errors v0.8.1 // indirect                      <------
	github.com/spf13/cobra v0.0.3
	github.com/spf13/pflag v1.0.3 // indirect
	golang.org/x/crypto v0.0.0-20190103213133-ff983b9c42bc // indirect
	golang.org/x/net v0.0.0-20190110200230-915654e7eabc // indirect
	golang.org/x/sys v0.0.0-20190116161447-11f53e031339 // indirect
	gopkg.in/AlecAivazis/survey.v1 v1.7.1
	gopkg.in/src-d/go-billy.v4 v4.3.0 // indirect
	gopkg.in/src-d/go-git.v4 v4.8.1
	gopkg.in/yaml.v2 v2.2.2
)

go 1.13
```

3. On line 15 of the file `utils/error.go` you can use the instruction `errors.WithMessage(err, "error")` from the package `github.com/pkg/errors` instead of `fmt.Sprintf("error: %s", err)`. Make sure to import the package.
4. Now that we are actually using the `error` package in our code, running `go mod tidy` will fix the dependency by removing the `// indirect` comment in `go.mod`.

```
module github.com/bmuschko/lets-gopher-exercise

require (
	github.com/Netflix/go-expect v0.0.0-20180928190340-9d1f4485533b // indirect
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/hinshun/vt10x v0.0.0-20180809195222-d55458df857c // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/mattn/go-colorable v0.0.9 // indirect
	github.com/mattn/go-isatty v0.0.4 // indirect
	github.com/mgutz/ansi v0.0.0-20170206155736-9520e82c474b // indirect
	github.com/mitchellh/go-homedir v1.0.0
	github.com/pkg/errors v0.8.1                      <------
	github.com/spf13/cobra v0.0.3
	github.com/spf13/pflag v1.0.3 // indirect
	golang.org/x/crypto v0.0.0-20190103213133-ff983b9c42bc // indirect
	golang.org/x/net v0.0.0-20190110200230-915654e7eabc // indirect
	golang.org/x/sys v0.0.0-20190116161447-11f53e031339 // indirect
	gopkg.in/AlecAivazis/survey.v1 v1.7.1
	gopkg.in/src-d/go-billy.v4 v4.3.0 // indirect
	gopkg.in/src-d/go-git.v4 v4.8.1
	gopkg.in/yaml.v2 v2.2.2
)

go 1.13
```

5. Running the command `go mod graph` shows that the version `v0.8.1` was selected by our package. We can also see that an earlier version of package was requested as transitive dependency by an external package.

```
$ go mod graph
github.com/bmuschko/lets-gopher-exercise github.com/Netflix/go-expect@v0.0.0-20180928190340-9d1f4485533b
github.com/bmuschko/lets-gopher-exercise github.com/emirpasic/gods@v1.12.0
github.com/bmuschko/lets-gopher-exercise github.com/hinshun/vt10x@v0.0.0-20180809195222-d55458df857c
github.com/bmuschko/lets-gopher-exercise github.com/inconshreveable/mousetrap@v1.0.0
github.com/bmuschko/lets-gopher-exercise github.com/kballard/go-shellquote@v0.0.0-20180428030007-95032a82bc51
github.com/bmuschko/lets-gopher-exercise github.com/mattn/go-colorable@v0.0.9
github.com/bmuschko/lets-gopher-exercise github.com/mattn/go-isatty@v0.0.4
github.com/bmuschko/lets-gopher-exercise github.com/mgutz/ansi@v0.0.0-20170206155736-9520e82c474b
github.com/bmuschko/lets-gopher-exercise github.com/mitchellh/go-homedir@v1.0.0
github.com/bmuschko/lets-gopher-exercise github.com/pkg/errors@v0.8.1            <------
github.com/bmuschko/lets-gopher-exercise github.com/spf13/cobra@v0.0.3
github.com/bmuschko/lets-gopher-exercise github.com/spf13/pflag@v1.0.3
github.com/bmuschko/lets-gopher-exercise golang.org/x/crypto@v0.0.0-20190103213133-ff983b9c42bc
github.com/bmuschko/lets-gopher-exercise golang.org/x/net@v0.0.0-20190110200230-915654e7eabc
github.com/bmuschko/lets-gopher-exercise golang.org/x/sys@v0.0.0-20190116161447-11f53e031339
github.com/bmuschko/lets-gopher-exercise gopkg.in/AlecAivazis/survey.v1@v1.7.1
github.com/bmuschko/lets-gopher-exercise gopkg.in/src-d/go-billy.v4@v4.3.0
github.com/bmuschko/lets-gopher-exercise gopkg.in/src-d/go-git.v4@v4.8.1
github.com/bmuschko/lets-gopher-exercise gopkg.in/yaml.v2@v2.2.2
gopkg.in/src-d/go-billy.v4@v4.3.0 github.com/kr/pretty@v0.1.0
gopkg.in/src-d/go-billy.v4@v4.3.0 golang.org/x/sys@v0.0.0-20180903190138-2b024373dcd9
gopkg.in/src-d/go-billy.v4@v4.3.0 gopkg.in/check.v1@v1.0.0-20180628173108-788fd7840127
gopkg.in/yaml.v2@v2.2.2 gopkg.in/check.v1@v0.0.0-20161208181325-20d25e280405
gopkg.in/src-d/go-git.v4@v4.8.1 github.com/alcortesm/tgz@v0.0.0-20161220082320-9c5fe88206d7
gopkg.in/src-d/go-git.v4@v4.8.1 github.com/anmitsu/go-shlex@v0.0.0-20161002113705-648efa622239
gopkg.in/src-d/go-git.v4@v4.8.1 github.com/davecgh/go-spew@v1.1.1
gopkg.in/src-d/go-git.v4@v4.8.1 github.com/emirpasic/gods@v1.9.0
gopkg.in/src-d/go-git.v4@v4.8.1 github.com/flynn/go-shlex@v0.0.0-20150515145356-3f9db97f8568
gopkg.in/src-d/go-git.v4@v4.8.1 github.com/gliderlabs/ssh@v0.1.1
gopkg.in/src-d/go-git.v4@v4.8.1 github.com/google/go-cmp@v0.2.0
gopkg.in/src-d/go-git.v4@v4.8.1 github.com/jbenet/go-context@v0.0.0-20150711004518-d14ea06fba99
gopkg.in/src-d/go-git.v4@v4.8.1 github.com/jessevdk/go-flags@v1.4.0
gopkg.in/src-d/go-git.v4@v4.8.1 github.com/kevinburke/ssh_config@v0.0.0-20180830205328-81db2a75821e
gopkg.in/src-d/go-git.v4@v4.8.1 github.com/mitchellh/go-homedir@v1.0.0
gopkg.in/src-d/go-git.v4@v4.8.1 github.com/pelletier/go-buffruneio@v0.2.0
gopkg.in/src-d/go-git.v4@v4.8.1 github.com/pkg/errors@v0.8.0                      <------
gopkg.in/src-d/go-git.v4@v4.8.1 github.com/pmezard/go-difflib@v1.0.0
gopkg.in/src-d/go-git.v4@v4.8.1 github.com/sergi/go-diff@v1.0.0
gopkg.in/src-d/go-git.v4@v4.8.1 github.com/src-d/gcfg@v1.4.0
gopkg.in/src-d/go-git.v4@v4.8.1 github.com/stretchr/testify@v1.2.2
gopkg.in/src-d/go-git.v4@v4.8.1 github.com/xanzy/ssh-agent@v0.2.0
gopkg.in/src-d/go-git.v4@v4.8.1 golang.org/x/crypto@v0.0.0-20180904163835-0709b304e793
gopkg.in/src-d/go-git.v4@v4.8.1 golang.org/x/net@v0.0.0-20180906233101-161cd47e91fd
gopkg.in/src-d/go-git.v4@v4.8.1 golang.org/x/text@v0.3.0
gopkg.in/src-d/go-git.v4@v4.8.1 gopkg.in/check.v1@v1.0.0-20180628173108-788fd7840127
gopkg.in/src-d/go-git.v4@v4.8.1 gopkg.in/src-d/go-billy.v4@v4.2.1
gopkg.in/src-d/go-git.v4@v4.8.1 gopkg.in/src-d/go-git-fixtures.v3@v3.1.1
gopkg.in/src-d/go-git.v4@v4.8.1 gopkg.in/warnings.v0@v0.1.2
gopkg.in/src-d/go-billy.v4@v4.2.1 github.com/kr/pretty@v0.1.0
gopkg.in/src-d/go-billy.v4@v4.2.1 golang.org/x/sys@v0.0.0-20180903190138-2b024373dcd9
gopkg.in/src-d/go-billy.v4@v4.2.1 gopkg.in/check.v1@v1.0.0-20180628173108-788fd7840127
github.com/kr/pretty@v0.1.0 github.com/kr/text@v0.1.0
github.com/kr/text@v0.1.0 github.com/kr/pty@v1.1.1
```