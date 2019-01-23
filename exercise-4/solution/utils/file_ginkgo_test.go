package utils_test

import (
	. "github.com/bmuschko/lets-gopher-exercise/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestFileUtils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "File Utils Suite")
}

var _ = Describe("File Utilities", func() {
	var newDir string

	BeforeEach(func() {
		dir, err := ioutil.TempDir("", "test")
		Expect(err).To(BeNil(), "Failed to create temporary directory %s", dir)
		defer os.RemoveAll(dir)

		newDir = filepath.Join(dir, "new")
	})

	Describe("Create Directory", func() {
		Context("that doesn't exist", func() {
			It("should be created", func() {
				err := CreateDir(newDir)
				Expect(err).To(BeNil(), "Failed to create new directory %s", newDir)

				fileInfo, err := os.Stat(newDir)
				Expect(err).To(BeNil(), "Expected directory to %s to be created", newDir)
				Expect(fileInfo.Mode().Perm()).To(Equal(os.FileMode(0755)), "Expected directory permissions 0755 but received %s", fileInfo.Mode().Perm())
			})
		})

		Context("that does already exist", func() {
			It("should not be recreated", func() {
				err := os.MkdirAll(newDir, 0755)
				Expect(err).To(BeNil(), "Failed to create new directory %s", newDir)
				err = CreateDir(newDir)
				Expect(err).To(BeNil(), "Failed to create new directory %s", newDir)

				fileInfo, err := os.Stat(newDir)
				Expect(err).To(BeNil(), "Expected directory to %s to be created", newDir)
				Expect(fileInfo.Mode().Perm()).To(Equal(os.FileMode(0755)), "Expected directory permissions 0755 but received %s", fileInfo.Mode().Perm())
			})
		})
	})
})
