package imageon

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func GetS3Client() *s3.Client {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	s3Client := s3.NewFromConfig(sdkConfig)

	return s3Client
}

func PresignUploadUrl(bucket, key string, metadata map[string]string) (string, error) {
	s3Client := GetS3Client()

	presigner := s3.NewPresignClient(s3Client)

	presigned, err := presigner.PresignPutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:   &bucket,
		Key:      &key,
		Metadata: metadata,
	}, func(options *s3.PresignOptions) {
		options.Expires = time.Duration(24 * time.Hour)
	})
	if err != nil {
		return "", err
	}

	return presigned.URL, nil
}
