package mongodb

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path"
)

func Dump(mongodbUri string, target string) error {
	tempDir, err := ioutil.TempDir("dir", "prefix")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tempDir)

	tempFile := path.Join(tempDir, "dump")

	// run mongodump
	cmd := exec.Command("mongodump", "--uri="+mongodbUri, "--gzip", "--archive="+tempFile)
	err = cmd.Run()

	if err != nil {
		return err
	}

	err = os.Rename(tempFile, target)
	if err != nil {
		return err
	}
	return nil
}
