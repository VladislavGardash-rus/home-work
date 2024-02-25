package http_handlers

import (
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/services"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/storage"
	"net/http"
)

type EventDataHandler struct {
	eventDataService *services.EventDataService
}

func NewEventDataHandler(storage storage.IStorage) *EventDataHandler {
	eventDataHandler := new(EventDataHandler)
	eventDataHandler.eventDataService = services.NewEventDataService(storage)
	return eventDataHandler
}

func (h *EventDataHandler) GetEvents(r *http.Request) (interface{}, error) {
	return nil, nil
}

func (h *EventDataHandler) PostCreateEvent(r *http.Request) (interface{}, error) {
	return nil, nil
}

func (h *EventDataHandler) PostUpdateEvent(r *http.Request) (interface{}, error) {
	return nil, nil
}

func (h *EventDataHandler) PostDeleteEvent(r *http.Request) (interface{}, error) {
	return nil, nil
}
