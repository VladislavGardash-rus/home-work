package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/models"
	memorystorage "github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/storage/sql"
	"time"
)

const DataBaseStorageType = "database"
const MemoryBaseStorageType = "memory"

func NewStorage(ctx context.Context, storageType string, connection string) (IStorage, error) {
	switch storageType {
	case DataBaseStorageType:
		storage := sqlstorage.New()
		err := storage.Connect(ctx, connection)
		if err != nil {
			return nil, err
		}

		return storage, err
	case MemoryBaseStorageType:
		return memorystorage.New(), nil
	default:
		return nil, errors.New(fmt.Sprintf("storage type %s not supported", storageType))
	}
}

type IStorage interface {
	Connect(ctx context.Context, params string) error
	Close() error

	CreateEvent(ctx context.Context, event models.Event) (int, error)
	UpdateEvent(ctx context.Context, id int, event models.Event) error
	DeleteEvent(ctx context.Context, id int) error
	DeleteEventsOldThenLastYear(ctx context.Context) error
	GetEvents(ctx context.Context) ([]models.Event, error)
	GetEventsByLastDay(ctx context.Context, date time.Time) ([]models.Event, error)
	GetEventsByLastWeek(ctx context.Context, date time.Time) ([]models.Event, error)
	GetEventsByLastMonth(ctx context.Context, date time.Time) ([]models.Event, error)
}
