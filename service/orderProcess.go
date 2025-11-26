package service

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type OrderService struct{}

func (o OrderService) Processs(dto types.Message) error {
	log.Println(*dto.Body)
	return nil
}
