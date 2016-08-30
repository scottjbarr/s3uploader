// The s3uploader package provides a little wrapper for uploading
// files to Amazon S3.
package s3uploader

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Uploader facilitates uploading files to Amazon S3.
type Uploader struct {
	Bucket string
	Key    string
	Body   []byte
}

// NewUploader returns a new Uploader
func NewUploader(bucket string, key string, body []byte) *Uploader {
	return &Uploader{
		Bucket: bucket,
		Key:    key,
		Body:   body,
	}
}

// Upload a file to S3
func (u *Uploader) Upload() error {
	reader := bytes.NewReader(u.Body)

	poi := s3.PutObjectInput{
		Bucket: &u.Bucket,
		Key:    &u.Key,
		Body:   reader,
	}

	service := s3.New(session.New())
	_, err := service.PutObject(&poi)

	return err
}
