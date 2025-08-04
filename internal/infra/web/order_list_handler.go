package web

import (
	"encoding/json"
	"net/http"

	"github.com/devfullcycle/20-CleanArch/internal/entity"
	"github.com/devfullcycle/20-CleanArch/internal/usecase"
)

type WebOrderListHandler struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewWebOrderListHandler(OrderRepository entity.OrderRepositoryInterface) *WebOrderListHandler {
	return &WebOrderListHandler{
		OrderRepository: OrderRepository,
	}
}

func (h *WebOrderListHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	listOrders := usecase.NewListOrdersUseCase(h.OrderRepository)
	output, err := listOrders.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
