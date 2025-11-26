package listener

import (
	"log"

	"github.com/caiquefds/sqs-listener-golang/client"
	"github.com/caiquefds/sqs-listener-golang/service"
)

func Start() {
	log.Printf("Starting queues...")
	listener := CreateListener(service.OrderService{}, "http://localhost:4566/000000000000/order-placement", client.CreateSQSClient())
	listener.Listen()
	log.Printf("Started all queues!")
	select {}
}
