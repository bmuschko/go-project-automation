# Exercise 5

In this exercise, you will learn how to build binaries for different target platforms.

1. Install Gox with the command `go get github.com/mitchellh/gox` or download the latest, pre-built binary from the [list of releases](github.com/mitchellh/gox).
2. Verify that Gox was installed properly by running `gox -h`. Gox should present you with all available CLI options.
3. Navigate to the root directory of the sample project and run the `gox` command. You should see multiple binary files.
4. Provide relevant CLI options to `gox` to only produce the binary for the appropriate platform on your machine e.g. `windows/amd64`.
5. Executing the `gox` command without any additional options does not provide a version in the file name nor does it pass the version to the linker for consumption by the program. Provide the appropriate CLI options to create binaries with the version `1.0`. Running the produced binary with the `version` option should render the appropriate version in the console output. Try to think of a way to pass the version to the linker. Make sure that the binaries are created in the subdirectory `dist`.
6. (Discuss) What if you wanted to provide the current Git commit hash as part of the version. Can you come up with a way to inject the value?
7. (Discuss) What challenges could arise for a team using Gox? Can you think of any improvements to the user experience?