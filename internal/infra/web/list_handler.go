package web

import (
	"encoding/json"
	"net/http"

	"github.com/devfullcycle/20-CleanArch/internal/entity"
	"github.com/devfullcycle/20-CleanArch/internal/usecase"
	"github.com/devfullcycle/20-CleanArch/pkg/events"
)

type WebListOrderHandler struct {
	EventDispatcher       events.EventDispatcherInterface
	OrderRepository       entity.OrderRepositoryInterface
	ListOrderCreatedEvent events.EventInterface
}

func NewWebListOrderHandler(
	EventDispatcher events.EventDispatcherInterface,
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreatedEvent events.EventInterface,
) *WebListOrderHandler {
	return &WebListOrderHandler{
		EventDispatcher: EventDispatcher,
		OrderRepository: OrderRepository,
	}
}

func (h *WebListOrderHandler) List(w http.ResponseWriter, r *http.Request) {

	listOrder := usecase.NewListOrderUseCase(h.OrderRepository, h.ListOrderCreatedEvent)
	output, err := listOrder.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
