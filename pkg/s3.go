package pkg

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	_ "github.com/lib/pq"
	"os"
)

type backupS3 struct {
	s3Bucket string
	s3Prefix string
}

func NewBackupS3(s3Bucket, s3Prefix string) backupS3 {
	if s3Bucket == "" {
		panic("missing s3Bucket")
	}
	if s3Prefix == "" {
		panic("missing s3Prefix")
	}
	return backupS3{
		s3Bucket: s3Bucket,
		s3Prefix: s3Prefix,
	}
}

func (b *backupS3) BackupToS3(backupFile string) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"), // Change this to your desired AWS region
	})
	if err != nil {
		return fmt.Errorf("failed to create AWS session: %w", err)
	}

	// Open the backup file
	file, err := os.Open(backupFile)
	if err != nil {
		return fmt.Errorf("failed to open backup file: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Create an S3 client
	s3Client := s3.New(sess)

	// Upload the file to S3
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(b.s3Bucket),
		Key:    aws.String(b.s3Prefix + "/" + backupFile),
		Body:   file,
	})
	if err != nil {
		return fmt.Errorf("failed to upload backup to S3: %w", err)
	}

	return nil

}
