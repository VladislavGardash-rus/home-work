package memorystorage

import (
	"context"
	"errors"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/models"
	"sort"
	"sync"
	"time"
)

var EventNotFoundError = errors.New("event not found")

type Storage struct {
	data   map[int]models.Event
	lastId int
	mu     sync.RWMutex //nolint:unused
}

func New() *Storage {
	return &Storage{data: make(map[int]models.Event)}
}

func (s *Storage) Connect(_ context.Context, params string) error {
	return nil
}

func (s *Storage) Close() error {
	return nil
}

func (s *Storage) CreateEvent(_ context.Context, event models.Event) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.lastId++

	event.ID = s.lastId
	s.data[s.lastId] = event

	return s.lastId, nil
}

func (s *Storage) UpdateEvent(_ context.Context, id int, event models.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.data[id]
	if !ok {
		return EventNotFoundError
	}

	e := event
	s.data[id] = e

	return nil
}

func (s *Storage) DeleteEvent(_ context.Context, id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.data[id]
	if !ok {
		return EventNotFoundError
	}

	delete(s.data, id)

	return nil
}

func (s *Storage) DeleteEventsOldThenLastYear(_ context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()
	for _, event := range s.data {
		if event.DateTimeEnd.Before(now.AddDate(-1, 0, 0)) {
			delete(s.data, event.ID)
		}
	}

	return nil
}

func (s *Storage) GetEvents(_ context.Context) ([]models.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]models.Event, 0)

	for _, event := range s.data {
		result = append(result, event)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].DateTimeStart.Before(result[j].DateTimeStart)
	})

	return result, nil
}

func (s *Storage) GetEventsByLastDay(_ context.Context, date time.Time) ([]models.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]models.Event, 0)

	year, month, day := date.Date()
	for _, event := range s.data {
		eventYear, eventMonth, eventDay := event.DateTimeStart.Date()
		if eventYear == year && eventMonth == month && eventDay == day {
			result = append(result, event)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].DateTimeStart.Before(result[j].DateTimeStart)
	})

	return result, nil
}

func (s *Storage) GetEventsByLastWeek(_ context.Context, date time.Time) ([]models.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]models.Event, 0)

	year, week := date.ISOWeek()
	for _, event := range s.data {
		eventYear, eventWeek := event.DateTimeStart.ISOWeek()
		if eventYear == year && eventWeek == week {
			result = append(result, event)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].DateTimeStart.Before(result[j].DateTimeStart)
	})

	return result, nil
}

func (s *Storage) GetEventsByLastMonth(_ context.Context, date time.Time) ([]models.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]models.Event, 0)

	year, month, _ := date.Date()
	for _, event := range s.data {
		eventYear, eventMonth, _ := event.DateTimeStart.Date()
		if eventYear == year && eventMonth == month {
			result = append(result, event)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].DateTimeStart.Before(result[j].DateTimeStart)
	})

	return result, nil
}
