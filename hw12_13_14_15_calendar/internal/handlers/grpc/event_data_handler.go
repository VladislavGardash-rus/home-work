package grpc_handlers

import (
	"context"
	"errors"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/api/calendar_grpc"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/models"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/services"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/storage"
	"golang.org/x/exp/slices"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type EventDataHandler struct {
	calendar_grpc.UnimplementedCalendarServer
	eventDataService *services.EventDataService
}

func NewEventDataHandler(storage storage.IStorage) *EventDataHandler {
	eventDataHandler := new(EventDataHandler)
	eventDataHandler.eventDataService = services.NewEventDataService(storage)
	return eventDataHandler
}

func (h *EventDataHandler) CreateEvent(ctx context.Context, event *calendar_grpc.Event) (*calendar_grpc.Event, error) {
	notificationDuration := event.GetNotificationDuration().AsDuration()

	newEvent := models.Event{}
	newEvent.Title = event.GetTitle()
	newEvent.DateTimeStart = event.GetDateTimeStart().AsTime()
	newEvent.DateTimeEnd = event.GetDateTimeEnd().AsTime()
	newEvent.Description = event.GetDescription()
	newEvent.UserId = int(event.GetUserId())
	newEvent.NotificationDuration = &notificationDuration

	id, err := h.eventDataService.CreateEvent(ctx, newEvent)
	if err != nil {
		return nil, err
	}
	event.Id = int32(id)

	return event, nil
}

func (h *EventDataHandler) UpdateEvent(ctx context.Context, event *calendar_grpc.Event) (*calendar_grpc.Event, error) {
	notificationDuration := event.GetNotificationDuration().AsDuration()

	newEvent := models.Event{}
	newEvent.ID = int(event.GetId())
	newEvent.Title = event.GetTitle()
	newEvent.DateTimeStart = event.GetDateTimeStart().AsTime()
	newEvent.DateTimeEnd = event.GetDateTimeEnd().AsTime()
	newEvent.Description = event.GetDescription()
	newEvent.UserId = int(event.GetUserId())
	newEvent.NotificationDuration = &notificationDuration

	err := h.eventDataService.UpdateEvent(ctx, newEvent.ID, newEvent)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (h *EventDataHandler) DeleteEvent(ctx context.Context, params *calendar_grpc.DeleteRequest) (*emptypb.Empty, error) {
	err := h.eventDataService.DeleteEvent(ctx, int(params.GetId()))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (h *EventDataHandler) DeleteOldEvents(ctx context.Context, params *emptypb.Empty) (*emptypb.Empty, error) {
	err := h.eventDataService.DeleteOldEvents(ctx)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (h *EventDataHandler) EventList(ctx context.Context, params *calendar_grpc.ListRequest) (*calendar_grpc.ListResult, error) {
	if params.GetPeriodType() != "" && !slices.Contains(services.Periods, params.GetPeriodType()) {
		return nil, errors.New("if you have param periodType, use values lastDay, lastWeek, lastMonth")
	}

	events, err := h.eventDataService.GetEvents(ctx, params.GetPeriodType())
	if err != nil {
		return nil, err
	}

	result := new(calendar_grpc.ListResult)
	for _, event := range events {
		tsDateTimeStart := timestamppb.New(event.DateTimeStart)
		tsDateTimeEnd := timestamppb.New(event.DateTimeEnd)
		notificationDuration := durationpb.New(*event.NotificationDuration)

		resultEvent := new(calendar_grpc.Event)
		resultEvent.Id = int32(event.ID)
		resultEvent.Title = event.Title
		resultEvent.DateTimeStart = tsDateTimeStart
		resultEvent.DateTimeEnd = tsDateTimeEnd
		resultEvent.Description = event.Description
		resultEvent.UserId = int32(event.UserId)
		resultEvent.NotificationDuration = notificationDuration

		result.Events = append(result.Events, resultEvent)
	}

	return result, nil
}
