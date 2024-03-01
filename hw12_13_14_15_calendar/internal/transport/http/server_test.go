package http_server

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/logger"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/models"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/storage"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	ctx := context.Background()

	err := logger.InitLogger("INFO")
	require.NoError(t, err)

	iStorage, err := storage.NewStorage(ctx, storage.MemoryBaseStorageType, "")
	require.NoError(t, err)

	server := httptest.NewServer(NewServer("", iStorage, "test").router)

	t.Run("create event", func(t *testing.T) {
		now := time.Now()
		duration := 1 * time.Minute

		newEvent := new(models.Event)
		newEvent.Title = "Title"
		newEvent.DateTimeStart = now
		newEvent.DateTimeEnd = now.Add(duration)
		newEvent.Description = "Description"
		newEvent.UserId = 1
		newEvent.NotificationDuration = &duration

		data, err := json.Marshal(newEvent)
		require.NoError(t, err)

		res, err := http.Post(server.URL+"/event", "application/json", bytes.NewReader(data))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.StatusCode)

		b, err := io.ReadAll(res.Body)
		require.NoError(t, err)

		createdEvent := new(models.Event)
		err = json.Unmarshal(b, createdEvent)
		require.NoError(t, err)

		checkEqual(t, newEvent, createdEvent)
	})

	//Далее по списку

	t.Run("shutdown", func(t *testing.T) {
		server.Close()
		err = iStorage.Close()
		require.NoError(t, err)
	})
}

func checkEqual(t *testing.T, newEvent *models.Event, createdEvent *models.Event) {
	require.Equal(t, newEvent.Title, createdEvent.Title)
	require.Equal(t, newEvent.DateTimeStart.Format("2006-01-02 15:04:05"), createdEvent.DateTimeStart.Format("2006-01-02 15:04:05"))
	require.Equal(t, newEvent.DateTimeEnd.Format("2006-01-02 15:04:05"), createdEvent.DateTimeEnd.Format("2006-01-02 15:04:05"))
	require.Equal(t, newEvent.Description, createdEvent.Description)
	require.Equal(t, newEvent.UserId, createdEvent.UserId)
	require.Equal(t, newEvent.NotificationDuration, createdEvent.NotificationDuration)
}
