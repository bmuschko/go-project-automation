package templ

import (
	"fmt"
	"go/build"
	"os"
	"path/filepath"

	"github.com/bmuschko/lets-gopher-exercise/utils"
)

func GenerateProject(templatePath string, goHomeBasePath string) {
	srcGoPath := filepath.Join(build.Default.GOPATH, "src")
	targetPath := filepath.Join(srcGoPath, goHomeBasePath)

	if _, err := os.Stat(targetPath); !os.IsNotExist(err) {
		fmt.Printf("The target directory %q already exists\n", targetPath)
		os.Exit(1)
	}

	fmt.Printf("Generating project in %q\n", targetPath)
	err := utils.CopyDir(templatePath, targetPath)
	utils.CheckIfError(err)
}
