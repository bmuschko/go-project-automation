package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateDirForNonExistentDirectory(t *testing.T) {
	dir, err := ioutil.TempDir("", "test")
	assert.Nil(t, err, "Failed to create temporary directory %s", dir)
	defer os.RemoveAll(dir)

	newDir := filepath.Join(dir, "new")
	err = CreateDir(filepath.Join(dir, "new"))
	assert.Nil(t, err, "Failed to create new directory %s", newDir)

	fileInfo, err := os.Stat(newDir)
	assert.Nil(t, err, "Expected directory to %s to be created", newDir)
	assert.Equal(t, os.FileMode(0755), fileInfo.Mode().Perm(), "Expected directory permissions 0755 but received %s", fileInfo.Mode().Perm())
}

func TestCreateDirForExistentDirectory(t *testing.T) {
	dir, err := ioutil.TempDir("", "test")
	assert.Nil(t, err, "Failed to create temporary directory %s", dir)
	defer os.RemoveAll(dir)

	newDir := filepath.Join(dir, "new")
	err = os.Mkdir(newDir, 0755)
	assert.Nil(t, err, "Failed to create new directory %s", newDir)
	err = CreateDir(filepath.Join(dir, "new"))
	assert.Nil(t, err, "Failed to create new directory %s", newDir)

	fileInfo, err := os.Stat(newDir)
	assert.Nil(t, err, "Expected directory to %s to be created", newDir)
	assert.Equal(t, os.FileMode(0755), fileInfo.Mode().Perm(), "Expected directory permissions 0755 but received %s", fileInfo.Mode().Perm())
}
