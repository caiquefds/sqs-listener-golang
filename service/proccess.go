package service

import "github.com/aws/aws-sdk-go-v2/service/sqs/types"

type Processor interface {
	Processs(dto types.Message) error
}
