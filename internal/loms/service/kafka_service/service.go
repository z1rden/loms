package kafka_service

import (
	"context"
	"loms/internal/loms/kafka"
	"loms/internal/loms/repository/kafka_storage"
	"sync"
)

type Service interface {
	SendMessages(ctx context.Context)
	StopSendMessages() error
}

type service struct {
	kafkaStorage    kafka_storage.Storage
	kafkaProducer   kafka.Producer
	sendMessagesWG  sync.WaitGroup
	sendMessageDone chan struct{}
}

func NewService(kafkaStorage kafka_storage.Storage, kafkaProducer kafka.Producer) Service {
	return &service{
		kafkaStorage:    kafkaStorage,
		kafkaProducer:   kafkaProducer,
		sendMessageDone: make(chan struct{}),
	}
}
