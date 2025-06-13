package service_provider

import (
	"context"
	"github.com/IBM/sarama"
	"loms/internal/loms/kafka"
	"loms/internal/loms/logger"
	"time"
)

func (s *ServiceProvider) GetProducer(ctx context.Context) kafka.Producer {
	if s.producer == nil {
		var err error
		s.producer, err = kafka.NewProducer(
			[]string{s.cfg.KafkaAddr},
			kafka.WithIdempotent(),
			kafka.WithRequiredAcks(sarama.WaitForAll),
			kafka.WithMaxOpenRequests(1),
			kafka.WithMaxRetries(5),
			kafka.WithRetryBackoff(10*time.Millisecond),
			// kafka.WithProducerPartitioner(sarama.NewRoundRobinPartitioner),
		)
		if err != nil {
			logger.Fatalf(ctx, "failed to create kafka producer: %v", err)
		}

		s.GetCloser(ctx).Add(s.producer.Close)
	}

	return s.producer
}
