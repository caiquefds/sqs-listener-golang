package config

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

var awsCfg aws.Config

type AWSConfig struct{}

func (a AWSConfig) configure() {
	log.Println("Loading AWS config...")

	loadedCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(os.Getenv("AWS_REGION")),
	)
	if err != nil {
		log.Fatalf("Failed to load AWS config: %v", err)
	}

	awsCfg = loadedCfg
	log.Println("AWS config loaded successfully.")
}

func (a AWSConfig) GetAWSConfig() aws.Config {
	return awsCfg
}

func (a AWSConfig) GetSQSClient() *sqs.Client {
	return sqs.NewFromConfig(a.GetAWSConfig(), func(o *sqs.Options) {
		o.BaseEndpoint = aws.String(os.Getenv("AWS_BASE_ENDPOINT"))
	})
}
