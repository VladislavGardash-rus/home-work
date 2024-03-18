package grpc_server

import (
	"context"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/api/calendar_grpc"
	grpc_handlers "github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/handlers/grpc"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/storage"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	iStorage, err := storage.NewStorage(ctx, storage.MemoryBaseStorageType, "")
	require.NoError(t, err)

	server := grpc.NewServer()
	calendar_grpc.RegisterCalendarServer(server, grpc_handlers.NewEventDataHandler(iStorage))

	listener := bufconn.Listen(1024 * 1024)
	go func() {
		select {
		case <-ctx.Done():
			return
		default:
			err = server.Serve(listener)
			require.NoError(t, err)
		}
	}()

	connFunc := func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(connFunc))
	require.NoError(t, err)

	client := calendar_grpc.NewCalendarClient(conn)

	now := time.Now()
	minute := 1 * time.Minute
	tsDateTimeStart := timestamppb.New(now)
	tsDateTimeEnd := timestamppb.New(now.Add(minute))
	notificationDuration := durationpb.New(minute)

	newEvent := new(calendar_grpc.Event)
	newEvent.Title = "Title"
	newEvent.DateTimeStart = tsDateTimeStart
	newEvent.DateTimeEnd = tsDateTimeEnd
	newEvent.Description = "Description"
	newEvent.UserId = int32(1)
	newEvent.NotificationDuration = notificationDuration

	createdEvent := new(calendar_grpc.Event)

	t.Run("create event", func(t *testing.T) {
		createdEvent, err = client.CreateEvent(ctx, newEvent)
		require.NoError(t, err)

		checkEqual(t, newEvent, createdEvent)
	})

	//Далее по списку

	t.Run("shutdown", func(t *testing.T) {
		cancel()

		err = conn.Close()
		require.NoError(t, err)

		server.GracefulStop()
	})
}

func checkEqual(t *testing.T, newEvent *calendar_grpc.Event, createdEvent *calendar_grpc.Event) {
	require.Equal(t, newEvent.GetTitle(), createdEvent.GetTitle())
	require.Equal(t, newEvent.GetDateTimeStart().AsTime(), createdEvent.GetDateTimeStart().AsTime())
	require.Equal(t, newEvent.GetDateTimeEnd().AsTime(), createdEvent.GetDateTimeEnd().AsTime())
	require.Equal(t, newEvent.GetDescription(), createdEvent.GetDescription())
	require.Equal(t, newEvent.GetUserId(), createdEvent.GetUserId())
	require.Equal(t, newEvent.GetNotificationDuration().AsDuration(), createdEvent.GetNotificationDuration().AsDuration())
}
