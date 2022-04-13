package storage

import (
	"bytes"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (storage *Storage) Upload(filePath string, fileName string) error {
	// Open the file from the file path
	upFile, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("could not open local filepath [%v]: %+v", filePath, err)
	}
	defer upFile.Close()

	// Get the file info
	upFileInfo, _ := upFile.Stat()
	var fileSize int64 = upFileInfo.Size()
	fileBuffer := make([]byte, fileSize)
	upFile.Read(fileBuffer)

	// Put Object to S3
	_, err = s3.New(storage.session).PutObject(&s3.PutObjectInput{
		Bucket:             storage.bucket,
		Key:                aws.String(fileName),
		Body:               bytes.NewReader(fileBuffer),
		ContentLength:      aws.Int64(fileSize),
		ContentDisposition: aws.String("attachment"),
	})

	return err
}
