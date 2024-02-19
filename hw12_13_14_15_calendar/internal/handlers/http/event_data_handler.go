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

func (h *EventDataHandler) GetEvents(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *EventDataHandler) PostCreateEvent(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *EventDataHandler) PostUpdateEvent(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *EventDataHandler) PostDeleteEvent(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
