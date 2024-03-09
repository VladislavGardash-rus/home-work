package services

import (
	"context"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/brokers/rabbit_mq"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/logger"
)

type EventSenderService struct {
	rabbitMqManager *rabbit_mq.Manager
}

func NewEventSenderService(rabbitMqManager *rabbit_mq.Manager) *EventSenderService {
	return &EventSenderService{rabbitMqManager: rabbitMqManager}
}

func (s *EventSenderService) Start(ctx context.Context) {
	err := s.rabbitMqManager.Consume(ctx, eventQueueName, s.sendNotification)
	if err != nil {
		logger.UseLogger().Error(err)
	}
}

func (s *EventSenderService) sendNotification(message []byte) error {
	logger.UseLogger().Info(string(message))
	return nil
}
