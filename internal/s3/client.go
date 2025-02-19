package s3

import (
	"codegen-service/internal/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

type S3Client struct {
	client *s3.S3
	bucket string
}

func NewS3Client(cfg *config.Config) *S3Client {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-central-1"),
		Credentials: credentials.NewStaticCredentials(cfg.S3AccessKey, cfg.S3SecretKey, ""),
		Endpoint:    aws.String(cfg.S3Path),
	})
	if err != nil {
		return nil
	}

	return &S3Client{
		client: s3.New(sess),
		bucket: cfg.S3Bucket,
	}
}

func (r *S3Client) UploadFile(file *os.File, key string) error {
	_, err := r.client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(r.bucket),
		Key:    aws.String(key + ".zip"),
		Body:   file,
		ACL:    aws.String(s3.ObjectCannedACLPrivate),
	})
	return err
}
