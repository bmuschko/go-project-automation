# Solution

1. One of the options to install `dep` is via `go get`: `go get -d -u github.com/golang/dep`. You will get the latest version of `dep`.
2. Run the command `dep ensure -add github.com/pkg/errors` to resolve the package.

```
$ dep ensure -add github.com/pkg/errors
Fetching sources...

"github.com/pkg/errors" is not imported by your project, and has been temporarily added to Gopkg.lock and vendor/.
If you run "dep ensure" again before actually importing it, it will disappear from Gopkg.lock and vendor/.
```

3. On line 15 of the file `utils/error.go` you can use the instruction `errors.WithMessage(err, "error")` from the package `github.com/pkg/errors` instead of `fmt.Sprintf("error: %s", err)`. Make sure to import the package.
4. Running the command `dep status` should show that the version `^0.8.1` was requested but the version `0.8.1` was selected. If a newer version would be released then `dep` would pick a newer version.

```
$ dep status
PROJECT                               CONSTRAINT     VERSION        REVISION  LATEST   PKGS USED
github.com/emirpasic/gods             v1.12.0        v1.12.0        1615341   v1.12.0  6
github.com/inconshreveable/mousetrap  v1.0           v1.0           76626ae   v1.0     1
github.com/jbenet/go-context          branch master  branch master  d14ea06   d14ea06  1
github.com/kballard/go-shellquote     branch master  branch master  95032a8   95032a8  1
github.com/kevinburke/ssh_config      0.5            0.5            81db2a7   0.5      1
github.com/mattn/go-colorable         v0.0.9         v0.0.9         167de6b   v0.0.9   1
github.com/mattn/go-isatty            v0.0.4         v0.0.4         6ca4dbf   v0.0.4   1
github.com/mgutz/ansi                 branch master  branch master  9520e82   9520e82  1
github.com/mitchellh/go-homedir       1.0.0          v1.0.0         ae18d6b   v1.0.0   1
github.com/pelletier/go-buffruneio    v0.2.0         v0.2.0         c37440a   v0.2.0   1
github.com/pkg/errors                 ^0.8.1         v0.8.1         ba968bf   v0.8.1   1. <------
github.com/sergi/go-diff              v1.0.0         v1.0.0         1744e29   v1.0.0   1
github.com/spf13/cobra                0.0.3          v0.0.3         ef82de7   v0.0.3   1
github.com/spf13/pflag                v1.0.3         v1.0.3         298182f   v1.0.3   1
github.com/src-d/gcfg                 v1.4.0         v1.4.0         1ac3a1a   v1.4.0   4
github.com/xanzy/ssh-agent            v0.2.0         v0.2.0         640f0ab   v0.2.0   1
golang.org/x/crypto                   branch master  branch master  ff983b9   7f87c0f  16
golang.org/x/net                      branch master  branch master  915654e   afe646c  1
golang.org/x/sys                      branch master  branch master  11f53e0   775f819  2
golang.org/x/text                     v0.3.0         v0.3.0         f21a4df   v0.3.0   6
gopkg.in/AlecAivazis/survey.v1        1.7.1          v1.7.1         e205523   v1.7.1   3
gopkg.in/src-d/go-billy.v4            v4.3.0         v4.3.0         9826264   v4.3.0   5
gopkg.in/src-d/go-git.v4              3dbfb89                       3dbfb89            40
gopkg.in/warnings.v0                  v0.1.2         v0.1.2         ec4a0fe   v0.1.2   1
gopkg.in/yaml.v2                      2.2.2          v2.2.2         51d6538   v2.2.2   1
```

5. Use the equal sign to pick a concrete version. Re-run the `dep ensure` command. The output of the `dep status` should change.

```yaml
[[constraint]]
  name = "github.com/pkg/errors"
  version = "=0.8.1"
```

```
$ dep status
PROJECT                               CONSTRAINT     VERSION        REVISION  LATEST   PKGS USED
github.com/emirpasic/gods             v1.12.0        v1.12.0        1615341   v1.12.0  6
github.com/inconshreveable/mousetrap  v1.0           v1.0           76626ae   v1.0     1
github.com/jbenet/go-context          branch master  branch master  d14ea06   d14ea06  1
github.com/kballard/go-shellquote     branch master  branch master  95032a8   95032a8  1
github.com/kevinburke/ssh_config      0.5            0.5            81db2a7   0.5      1
github.com/mattn/go-colorable         v0.0.9         v0.0.9         167de6b   v0.0.9   1
github.com/mattn/go-isatty            v0.0.4         v0.0.4         6ca4dbf   v0.0.4   1
github.com/mgutz/ansi                 branch master  branch master  9520e82   9520e82  1
github.com/mitchellh/go-homedir       1.0.0          v1.0.0         ae18d6b   v1.0.0   1
github.com/pelletier/go-buffruneio    v0.2.0         v0.2.0         c37440a   v0.2.0   1
github.com/pkg/errors                 0.8.1          v0.8.1         ba968bf   v0.8.1   1 <------
github.com/sergi/go-diff              v1.0.0         v1.0.0         1744e29   v1.0.0   1
github.com/spf13/cobra                0.0.3          v0.0.3         ef82de7   v0.0.3   1
github.com/spf13/pflag                v1.0.3         v1.0.3         298182f   v1.0.3   1
github.com/src-d/gcfg                 v1.4.0         v1.4.0         1ac3a1a   v1.4.0   4
github.com/xanzy/ssh-agent            v0.2.0         v0.2.0         640f0ab   v0.2.0   1
golang.org/x/crypto                   branch master  branch master  ff983b9   7f87c0f  16
golang.org/x/net                      branch master  branch master  915654e   afe646c  1
golang.org/x/sys                      branch master  branch master  11f53e0   775f819  2
golang.org/x/text                     v0.3.0         v0.3.0         f21a4df   v0.3.0   6
gopkg.in/AlecAivazis/survey.v1        1.7.1          v1.7.1         e205523   v1.7.1   3
gopkg.in/src-d/go-billy.v4            v4.3.0         v4.3.0         9826264   v4.3.0   5
gopkg.in/src-d/go-git.v4              3dbfb89                       3dbfb89            40
gopkg.in/warnings.v0                  v0.1.2         v0.1.2         ec4a0fe   v0.1.2   1
gopkg.in/yaml.v2                      2.2.2          v2.2.2         51d6538   v2.2.2   1
```