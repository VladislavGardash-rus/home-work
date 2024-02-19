package services

import (
	"context"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/logger"
	"time"
)

type EventSenderService struct{}

func NewEventSenderService() *EventSenderService {
	return &EventSenderService{}
}

func (s *EventSenderService) Start(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Minute)

	for {
		select {
		case <-ctx.Done():
			logger.UseLogger().Info("event sender stopped")
			return
		case <-ticker.C:

		}
	}
}
