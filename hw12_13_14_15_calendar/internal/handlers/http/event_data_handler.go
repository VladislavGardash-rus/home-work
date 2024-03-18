//nolint:all
package http_handlers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/models"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/services"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/storage"
	"github.com/gorilla/mux"
	"golang.org/x/exp/slices"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type EventDataHandler struct {
	eventDataService *services.EventDataService
}

func NewEventDataHandler(storage storage.IStorage) *EventDataHandler {
	eventDataHandler := new(EventDataHandler)
	eventDataHandler.eventDataService = services.NewEventDataService(storage)
	return eventDataHandler
}

func (h *EventDataHandler) GetEvents(ctx context.Context, r *http.Request) (interface{}, error) {
	values, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		return nil, err
	}

	if values.Has("periodType") && !slices.Contains(services.Periods, values.Get("periodType")) {
		return nil, errors.New("if you have param periodType, use values lastDay, lastWeek, lastMonth")
	}

	return h.eventDataService.GetEvents(ctx, values.Get("periodType"))
}

func (h *EventDataHandler) PostCreateEvent(ctx context.Context, r *http.Request) (interface{}, error) {
	event := new(models.Event)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, event)
	if err != nil {
		return nil, err
	}

	id, err := h.eventDataService.CreateEvent(ctx, *event)
	if err != nil {
		return nil, err
	}
	event.ID = id

	return event, nil
}

func (h *EventDataHandler) PostUpdateEvent(ctx context.Context, r *http.Request) (interface{}, error) {
	event := new(models.Event)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, event)
	if err != nil {
		return nil, err
	}

	err = h.eventDataService.UpdateEvent(ctx, event.ID, *event)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (h *EventDataHandler) PostDeleteEvent(ctx context.Context, r *http.Request) (interface{}, error) {
	variables := mux.Vars(r)
	idParam, ok := variables["id"]
	if !ok {
		return nil, errors.New("id is missing in parameters")
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return nil, err
	}

	err = h.eventDataService.DeleteEvent(ctx, id)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (h *EventDataHandler) PostDeleteOldEvents(ctx context.Context, r *http.Request) (interface{}, error) {
	err := h.eventDataService.DeleteOldEvents(ctx)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
