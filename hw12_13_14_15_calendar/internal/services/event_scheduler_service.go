package services

import (
	"context"
	"time"
)

type EventSchedulerService struct {
}

func NewEventSchedulerService() *EventSchedulerService {
	return &EventSchedulerService{}
}

func (s *EventSchedulerService) Start(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Minute)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:

		}
	}
}
