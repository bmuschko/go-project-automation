package utils_test

import (
	. "github.com/bmuschko/lets-gopher-exercise/utils"
	. "github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateDirForNonExistentDirectoryWithTestify(t *testing.T) {
	dir, err := ioutil.TempDir("", "test")
	Nil(t, err, "Failed to create temporary directory %s", dir)
	defer os.RemoveAll(dir)

	newDir := filepath.Join(dir, "new")
	err = CreateDir(newDir)
	Nil(t, err, "Failed to create new directory %s", newDir)

	fileInfo, err := os.Stat(newDir)
	Nil(t, err, "Expected directory to %s to be created", newDir)
	Equal(t, os.FileMode(0755), fileInfo.Mode().Perm(), "Expected directory permissions 0755 but received %s", fileInfo.Mode().Perm())
}

func TestCreateDirForExistentDirectoryWithTestify(t *testing.T) {
	dir, err := ioutil.TempDir("", "test")
	Nil(t, err, "Failed to create temporary directory %s", dir)
	defer os.RemoveAll(dir)

	newDir := filepath.Join(dir, "new")
	err = os.Mkdir(newDir, 0755)
	Nil(t, err, "Failed to create new directory %s", newDir)
	err = CreateDir(newDir)
	Nil(t, err, "Failed to create new directory %s", newDir)

	fileInfo, err := os.Stat(newDir)
	Nil(t, err, "Expected directory to %s to be created", newDir)
	Equal(t, os.FileMode(0755), fileInfo.Mode().Perm(), "Expected directory permissions 0755 but received %s", fileInfo.Mode().Perm())
}
