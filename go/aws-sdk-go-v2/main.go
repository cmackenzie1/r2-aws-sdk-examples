package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	r2EndpointResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{URL: os.Getenv("R2_ENDPOINT"), SigningRegion: "auto"}, nil
	})

	accessKeyId := os.Getenv("R2_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("R2_SECRET_ACCESS_KEY")

	cfg, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithRegion("auto"),
		config.WithEndpointResolverWithOptions(r2EndpointResolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, secretAccessKey, "")),
		config.WithClientLogMode(aws.LogRetries|aws.LogRequest|aws.LogSigning))
	if err != nil {
		panic(err)
	}

	r2 := s3.NewFromConfig(cfg)

	resp, err := r2.ListBuckets(context.Background(), &s3.ListBucketsInput{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("response: %#v", resp)

}
