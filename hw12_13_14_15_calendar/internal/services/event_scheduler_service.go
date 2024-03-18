package services

import (
	"context"
	"encoding/json"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/brokers/rabbit_mq"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/logger"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/storage"
	"time"
)

const eventQueueName = "event"

type EventSchedulerService struct {
	rabbitMqManager *rabbit_mq.Manager
	storage         storage.IStorage
}

func NewEventSchedulerService(rabbitMqManager *rabbit_mq.Manager, storage storage.IStorage) *EventSchedulerService {
	return &EventSchedulerService{rabbitMqManager: rabbitMqManager, storage: storage}
}

func (s *EventSchedulerService) Start(ctx context.Context) {
	//для проверки можно сделать маленький интервал, но будут дублироваться события при отправке в очередь
	//можно сделать признак isSent в объекте Event и проверять по нему, но мне лень
	ticker := time.NewTicker(24 * time.Hour)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := s.sentEvents(ctx)
			if err != nil {
				logger.UseLogger().Error(err)
			}

			err = s.deleteOldEvents(ctx)
			if err != nil {
				logger.UseLogger().Error(err)
			}
		}
	}
}

func (s *EventSchedulerService) sentEvents(ctx context.Context) error {
	now := time.Now()
	events, err := s.storage.GetEventsByLastDay(ctx, now.AddDate(0, 0, -1))
	if err != nil {
		return err
	}

	for _, event := range events {
		b, err := json.Marshal(&event)
		if err != nil {
			return err
		}

		err = s.rabbitMqManager.Send(b, eventQueueName, "application/json")
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *EventSchedulerService) deleteOldEvents(ctx context.Context) error {
	err := s.storage.DeleteEventsOldThenLastYear(ctx)
	if err != nil {
		return err
	}

	return nil
}
