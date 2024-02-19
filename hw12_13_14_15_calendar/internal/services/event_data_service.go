package services

import (
	"context"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/models"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/storage"
	"time"
)

const (
	periodTypeLastDay   = "lastDay"
	periodTypeLastWeek  = "lastWeek"
	periodTypeLastMonth = "lastMonth"
)

type EventDataService struct {
	storage storage.IStorage
}

func NewEventDataService(storage storage.IStorage) *EventDataService {
	return &EventDataService{storage: storage}
}

func (s *EventDataService) CreateEvent(ctx context.Context, event models.Event) (int, error) {
	return s.storage.CreateEvent(ctx, event)
}

func (s *EventDataService) UpdateEvent(ctx context.Context, id int, event models.Event) error {
	return s.storage.UpdateEvent(ctx, id, event)
}

func (s *EventDataService) DeleteEvent(ctx context.Context, id int, deleteOld bool) error {
	if deleteOld {
		return s.storage.DeleteEventsOldThenLastYear(ctx)
	}

	return s.storage.DeleteEvent(ctx, id)
}

func (s *EventDataService) DeleteEventsByLastYear(ctx context.Context) error {
	return s.storage.DeleteEventsOldThenLastYear(ctx)
}

func (s *EventDataService) GetEvents(ctx context.Context, periodType string) ([]models.Event, error) {
	switch periodType {
	case periodTypeLastDay:
		period := time.Now().AddDate(0, 0, -1)
		return s.storage.GetEventsByLastDay(ctx, period)
	case periodTypeLastWeek:
		period := time.Now().AddDate(0, 0, -7)
		return s.storage.GetEventsByLastWeek(ctx, period)
	case periodTypeLastMonth:
		period := time.Now().AddDate(0, -1, 0)
		return s.storage.GetEventsByLastMonth(ctx, period)
	default:
		return s.storage.GetEvents(ctx)
	}
}
