package client

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/caiquefds/sqs-listener-golang/config"
)

type SqsClient struct {
	AwsSqsClient *sqs.Client
}

func (s SqsClient) GetMessages(ctx context.Context, queueUrl string, maxMessages int32, waitTime int32) ([]types.Message, error) {
	var messages []types.Message
	result, err := s.AwsSqsClient.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueUrl),
		MaxNumberOfMessages: maxMessages,
		WaitTimeSeconds:     waitTime,
	})
	if err != nil {
		log.Printf("Couldn't get messages from queue %v. Here's why: %v\n", queueUrl, err)
	} else {
		messages = result.Messages
	}
	return messages, err
}

func (s SqsClient) DeleteMessages(ctx context.Context, queueUrl string, messages []types.Message) error {
	entries := make([]types.DeleteMessageBatchRequestEntry, len(messages))
	for msgIndex := range messages {
		entries[msgIndex].Id = aws.String(fmt.Sprintf("%v", msgIndex))
		entries[msgIndex].ReceiptHandle = messages[msgIndex].ReceiptHandle
	}

	log.Println(entries)

	_, err := s.AwsSqsClient.DeleteMessageBatch(ctx, &sqs.DeleteMessageBatchInput{
		Entries:  entries,
		QueueUrl: aws.String(queueUrl),
	})

	if err != nil {
		log.Printf("Couldn't delete messages from queue %v. Here's why: %v\n", queueUrl, err)
	}

	return err
}

func CreateSQSClient() SqsClient {
	sqsClient := SqsClient{}
	sqsClient.AwsSqsClient = config.AWSConfig{}.GetSQSClient()
	return sqsClient
}
