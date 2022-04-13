package storage

import (
	"fmt"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var lock = &sync.Mutex{}
var storageInstance *Storage

type Storage struct {
	client  *s3.S3
	bucket  *string
	session *session.Session
}

func New() Storage {
	accessKeyID := os.Getenv("S3_ACCESS_KEY_ID")
	secretKeyID := os.Getenv("S3_SECRET_ACCESS_KEY")
	endpoint := os.Getenv("S3_ENDPOINT")
	region := os.Getenv("S3_REGION")
	bucket := os.Getenv("S3_BUCKET")

	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(accessKeyID, secretKeyID, ""),
		Endpoint:         aws.String(endpoint),
		Region:           aws.String(region),
		S3ForcePathStyle: aws.Bool(false),
	}

	newSession := session.New(s3Config)
	s3Client := s3.New(newSession)

	storage := Storage{}

	storage.session = newSession
	storage.client = s3Client
	storage.bucket = aws.String(bucket)
	return storage
}

func GetInstance() *Storage {
	if storageInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if storageInstance == nil {
			fmt.Println("Creating Storage single instance.")
			storage := New()
			storageInstance = &storage
		}
	}

	return storageInstance
}
