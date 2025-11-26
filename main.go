package main

import (
	"github.com/caiquefds/sqs-listener-golang/config"
	"github.com/caiquefds/sqs-listener-golang/listener"
)

func main() {
	config.Configure()
	listener.Start()
}
