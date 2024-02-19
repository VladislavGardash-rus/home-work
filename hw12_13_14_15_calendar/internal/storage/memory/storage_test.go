package memorystorage

import (
	"context"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/models"
	"github.com/stretchr/testify/require"
	"sync"
	"testing"
	"time"
)

func TestMemoryStorage(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		storage := New()

		ctx := context.Background()
		emptyEventsList := make([]models.Event, 0)
		events, err := storage.GetEvents(ctx)
		require.NoError(t, err)
		require.Equal(t, emptyEventsList, events)
	})

	t.Run("use simple script", func(t *testing.T) {
		storage := New()

		ctx := context.Background()
		now := time.Now()
		duration := 24 * time.Hour

		event := models.Event{
			Title:                "TestEventTitle",
			DateTimeStart:        now.AddDate(0, 0, -1),
			DateTimeEnd:          now,
			Description:          "TestEventDescription",
			UserId:               1,
			NotificationDuration: &duration,
		}

		id, err := storage.CreateEvent(ctx, event)
		require.NoError(t, err)

		events, err := storage.GetEvents(ctx)
		require.NoError(t, err)
		require.Equal(t, len(events), 1)
		require.Equal(t, events[0].Title, event.Title)

		event.Title = "NewTestEventTitle"
		err = storage.UpdateEvent(ctx, id, event)
		require.NoError(t, err)

		events, err = storage.GetEvents(ctx)
		require.NoError(t, err)
		require.Equal(t, len(events), 1)
		require.Equal(t, events[0].Title, event.Title)

		err = storage.DeleteEvent(ctx, id)
		require.NoError(t, err)

		err = storage.UpdateEvent(ctx, id, event)
		require.Equal(t, EventNotFoundError, err)
	})

	t.Run("multithreading", func(t *testing.T) {
		storage := New()
		wg := &sync.WaitGroup{}
		wg.Add(2)

		ctx := context.Background()
		event := models.Event{
			Title:       "TestEventTitle",
			Description: "TestEventDescription",
		}

		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				storage.CreateEvent(ctx, event)
			}
		}()

		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				storage.CreateEvent(ctx, event)
			}
		}()

		wg.Wait()
	})
}
