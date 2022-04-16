package main

import (
	"log"
	"os"
	"path"
	"time"

	"github.com/rwv/mongodb-backup-s3/mongodb"
	"github.com/rwv/mongodb-backup-s3/storage"

	"github.com/robfig/cron/v3"
)

func backup(mongodbUri string) error {
	log.Print("Starting backup...")
	log.Print("Mongodb Uri: " + mongodbUri)

	now := time.Now()
	uploadFilename := "mongodb_dump_" + now.Format("20060102150405") + ".gz"
	log.Print("Backup Filename: " + uploadFilename)

	tempDir, err := os.MkdirTemp("", "mongodb-backup-s3-")

	if err != nil {
		return err
	}
	defer os.RemoveAll(tempDir)

	log.Print("Temp Dir: " + tempDir)

	tempFilename := path.Join(tempDir, "dump.gz")
	log.Print("Temp File: " + tempFilename)

	log.Print("Dumping to temp file...")
	err = mongodb.Dump(mongodbUri, tempFilename)
	if err != nil {
		return err
	}

	storage := storage.New()

	log.Print("Uploading to S3...")
	err = storage.Upload(tempFilename, uploadFilename)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	mongodbUri := os.Getenv("MONGODB_URI")
	log.Print("Mongodb Uri: " + mongodbUri)

	c := cron.New()

	c.AddFunc("@daily", func() {
		err := backup(mongodbUri)
		if err != nil {
			log.Fatal(err)
		}
	})

	c.Start()

	for {
		time.Sleep(time.Second)
	}
}
