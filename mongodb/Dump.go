package mongodb

import (
	"os"
	"time"

	"github.com/mongodb/mongo-tools/mongodump"

	"log"
	"path"
)

const (
	progressBarLength   = 24
	progressBarWaitTime = time.Second * 3
)

const VersionStr = "100.5.2"
const GitCommit = "e2842eb54930"

func Dump(mongodbUri string, target string) error {
	tempDir, err := os.MkdirTemp("", "mongodb-backup-s3-")

	if err != nil {
		return err
	}
	defer os.RemoveAll(tempDir)

	tempFile := path.Join(tempDir, "dump")

	args := []string{"--uri=" + mongodbUri, "--gzip", "--archive=" + tempFile}

	opts, err := mongodump.ParseOptions(args, VersionStr, GitCommit)
	if err != nil {
		return err
	}

	dump := mongodump.MongoDump{
		ToolOptions:   opts.ToolOptions,
		OutputOptions: opts.OutputOptions,
		InputOptions:  opts.InputOptions,
	}

	if err = dump.Init(); err != nil {
		return err
	}

	if err = dump.Dump(); err != nil {
		return err
	}

	log.Print("Moving " + tempFile + " to " + target)
	err = os.Rename(tempFile, target)
	if err != nil {
		return err
	}

	return nil
}
