package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	accessKeyId := os.Getenv("R2_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("R2_SECRET_ACCESS_KEY")

	cfg := &aws.Config{
		Endpoint:    aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", os.Getenv("R2_ACCOUNT_ID"))),
		Region:      aws.String("auto"),
		Credentials: credentials.NewStaticCredentials(accessKeyId, secretAccessKey, ""),
	}
	sess := session.Must(session.NewSession(cfg))

	r2 := s3.New(sess)

	head, err := r2.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		panic(err)
	}

	for _, bucket := range head.Buckets {
		fmt.Printf("%s\n", *bucket.Name)
	}
}
