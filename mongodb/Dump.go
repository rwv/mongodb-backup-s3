package mongodb

import (
	"log"
	"os"
	"os/exec"
	"path"
)

func Dump(mongodbUri string, target string) error {
	tempDir, err := os.MkdirTemp("", "mongodb-backup-s3-")

	if err != nil {
		return err
	}
	defer os.RemoveAll(tempDir)

	tempFile := path.Join(tempDir, "dump")

	// run mongodump
	cmd := exec.Command("/usr/bin/mongodump", "--uri="+mongodbUri, "--gzip", "--archive="+tempFile)
	log.Print(cmd)
	err = cmd.Run()

	if err != nil {
		return err
	}

	log.Print("Moving " + tempFile + " to " + target)
	err = os.Rename(tempFile, target)
	if err != nil {
		return err
	}
	return nil
}
