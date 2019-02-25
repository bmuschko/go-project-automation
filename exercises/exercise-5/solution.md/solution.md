# Solution

3. It takes quite a while to produce all binary files. The output should look similar to this:

```
$ gox
Number of parallel builds: 7

-->    linux/mipsle: github.com/bmuschko/lets-gopher-exercise
-->     linux/amd64: github.com/bmuschko/lets-gopher-exercise
-->   windows/amd64: github.com/bmuschko/lets-gopher-exercise
-->      netbsd/arm: github.com/bmuschko/lets-gopher-exercise
-->      darwin/386: github.com/bmuschko/lets-gopher-exercise
-->     freebsd/386: github.com/bmuschko/lets-gopher-exercise
-->    darwin/amd64: github.com/bmuschko/lets-gopher-exercise
-->    linux/mips64: github.com/bmuschko/lets-gopher-exercise
-->     linux/s390x: github.com/bmuschko/lets-gopher-exercise
-->       linux/arm: github.com/bmuschko/lets-gopher-exercise
-->       linux/386: github.com/bmuschko/lets-gopher-exercise
-->      netbsd/386: github.com/bmuschko/lets-gopher-exercise
-->     freebsd/arm: github.com/bmuschko/lets-gopher-exercise
-->  linux/mips64le: github.com/bmuschko/lets-gopher-exercise
-->    netbsd/amd64: github.com/bmuschko/lets-gopher-exercise
-->      linux/mips: github.com/bmuschko/lets-gopher-exercise
-->     openbsd/386: github.com/bmuschko/lets-gopher-exercise
-->   freebsd/amd64: github.com/bmuschko/lets-gopher-exercise
-->   openbsd/amd64: github.com/bmuschko/lets-gopher-exercise
-->     windows/386: github.com/bmuschko/lets-gopher-exercise
```

4. Gox provides two different options to generate a particular binary for a target platform. You can use invidual CLI options `-os` and `-arch`: `gox -os="windows" -arch="amd64"`. Alternatively, you can provide a combination of OS and architecture with the `-osarch` CLI option: `gox -osarch="windows/amd64"`. For a list of available OS and architecure values run the command `gox -osarch-list`. The output should look similar to this:

```
$ gox -osarch="windows/amd64"
Number of parallel builds: 7

-->   windows/amd64: github.com/bmuschko/lets-gopher-exercise
```

5. By default, Gox uses the pattern `{{.Dir}}_{{.OS}}_{{.Arch}}` for building the file name of a binary. You can set a custom output file pattern with the CLI option `-output`. To set a value for the `version` variable in `main.go`, you will have to provide a linker option: `-ldflags="-X main.version=1.0"`. The full command could look as such: `gox -osarch="darwin/amd64" -output "dist/{{.Dir}}-1.0-{{.OS}}-{{.Arch}}" -ldflags="-X main.version=1.0"`. The output should look similar to this:

```
$ gox -osarch="darwin/amd64" -output "dist/{{.Dir}}-1.0-{{.OS}}-{{.Arch}}" -ldflags="-X main.version=1.0"
Number of parallel builds: 7

-->    darwin/amd64: github.com/bmuschko/lets-gopher-exercise
```

6. You can inject the output of a Git command line execution as expression: `-ldflags="-X main.version=$(git rev-parse HEAD)"`. The output should look similar to this:

```
gox -osarch="darwin/amd64" -output "dist/{{.Dir}}-1.0-{{.OS}}-{{.Arch}}" -ldflags="-X main.version=$(git rev-parse HEAD)"
Number of parallel builds: 7

-->    darwin/amd64: github.com/bmuschko/lets-gopher-exercise
```

7. Different team members need to know and understand available CLI options. Often times you will want to use the same set of options but with different values. It would be great of Gox would support setting values in a configuration file instead of having to provide them on the command line.