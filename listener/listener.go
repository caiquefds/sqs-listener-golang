package listener

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/caiquefds/sqs-listener-golang/client"
	"github.com/caiquefds/sqs-listener-golang/service"
)

type Listener struct {
	sqsClient         client.SqsClient
	queueUrl          string
	processor         service.Processor
	processedMessages chan types.Message
	receivedMessages  chan types.Message
}

func (l Listener) Listen() {

	l.receivedMessages = make(chan types.Message)
	l.processedMessages = make(chan types.Message)

	go l.getMessages()
	go l.process()
	go l.acknowledgeMessages()

}

func (l Listener) getMessages() {
	for {
		msg, err := l.sqsClient.GetMessages(context.TODO(), l.queueUrl, 10, 10)

		if err != nil {
			log.Printf("Error getting messages\n: %v", err)
		}

		for _, m := range msg {
			log.Printf("Received Message id %v and address %v\n", *m.MessageId, m)
			l.receivedMessages <- m
		}
	}
}

func (l Listener) process() {
	for msg := range l.receivedMessages {
		log.Printf("Processing: id %v and address %v\n", *msg.MessageId, msg)
		l.processor.Processs(msg)
		l.processedMessages <- msg
	}
}

func (l Listener) acknowledgeMessages() {

	for m := range l.processedMessages {
		log.Printf("Deleting message id %v and address %v", *m.MessageId, m)
		err := l.sqsClient.DeleteMessages(context.TODO(), l.queueUrl, []types.Message{m})
		if err != nil {
			log.Printf("Couldn't delete messages from queue %v. Here's why: %v\n", l.queueUrl, err)
		}
	}

}

func CreateListener(processor service.Processor, url string, sqs client.SqsClient) Listener {
	return Listener{
		processor:         processor,
		queueUrl:          url,
		sqsClient:         sqs,
		receivedMessages:  make(chan types.Message),
		processedMessages: make(chan types.Message),
	}
}
