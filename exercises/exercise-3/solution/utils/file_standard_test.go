package utils_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	. "github.com/bmuschko/lets-gopher-exercise/utils"
)

func TestCreateDirForNonExistentDirectoryWithTestingPackage(t *testing.T) {
	dir, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Errorf("Failed to create temporary directory %s", dir)
	}
	defer os.RemoveAll(dir)

	newDir := filepath.Join(dir, "new")
	err = CreateDir(newDir)
	if err != nil {
		t.Errorf("Failed to create new directory %s", newDir)
	}

	fileInfo, err := os.Stat(newDir)
	if err != nil {
		t.Errorf("Expected directory to %s to be created", newDir)
	}
	if fileInfo.Mode().Perm() != os.FileMode(0755) {
		t.Errorf("Expected directory permissions 0755 but received %s", fileInfo.Mode().Perm())
	}
}

func TestCreateDirForExistentDirectoryWithTestingPackage(t *testing.T) {
	dir, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Errorf("Failed to create temporary directory %s", dir)
	}
	defer os.RemoveAll(dir)

	newDir := filepath.Join(dir, "new")
	err = os.Mkdir(newDir, 0755)
	if err != nil {
		t.Errorf("Failed to create new directory %s", newDir)
	}
	err = CreateDir(newDir)
	if err != nil {
		t.Errorf("Failed to create new directory %s", newDir)
	}

	fileInfo, err := os.Stat(newDir)
	if err != nil {
		t.Errorf("Expected directory to %s to be created", newDir)
	}
	if fileInfo.Mode().Perm() != os.FileMode(0755) {
		t.Errorf("Expected directory permissions 0755 but received %s", fileInfo.Mode().Perm())
	}
}
