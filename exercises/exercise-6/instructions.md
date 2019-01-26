# Exercise 6

In the exercise, you will practice how to use [GoReleaser](https://goreleaser.com/) for generating binaries and publishing them to GitHub releases.

1. Install GoReleaser on your machine. You can find the installation instructions on the [web page](https://goreleaser.com/). Alternatively, download and install the latest binary from the [release page](https://github.com/goreleaser/goreleaser/releases) if you do not use any of the proposed package managers. Run `goreleaser -h` to verify the proper installation.
2. [Auto-generate](https://goreleaser.com/quick-start/) a GoReleaser configuration file for your sample project with the command `goreleaser init`.
3. Modify the `.goreleaser.yml` file so that the produced archive files include the `LICENSE` file of the project. Archives targeting the Windows OS, the archive file format should be `.zip` instead of `.tar.gz`.
4. Modify the `.goreleaser.yml` file so that the Git tag is passed to the version variable in `main.go`.
5. Generate the binaries _without_ publishing them to GitHub. Please consult the user guide for more information.
6. Create a local Git tag and publish the produced archives to GitHub. You should see the release on the [GitHub releases tab](github.com/<your-github-account>/lets-gopher/releases).