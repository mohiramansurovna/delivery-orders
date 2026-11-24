package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mohiramansurovna/orders/internal/application/services"
	"github.com/mohiramansurovna/orders/internal/interface/httpapi/dtos"
	"github.com/mohiramansurovna/orders/internal/interface/httpapi/errorhandler"
	"github.com/mohiramansurovna/orders/internal/interface/httpapi/httpapiutil"
	"github.com/mohiramansurovna/orders/internal/interface/httpapi/presenters"
)

type OrdersController struct {
	ordersService *services.OrdersService
}

func NewOrdersController(ordersService *services.OrdersService) *OrdersController {
	return &OrdersController{ordersService: ordersService}
}

func (this *OrdersController) Routes(mux *http.ServeMux) {
	mux.HandleFunc("POST /orders", errorhandler.Wrap(this.create))
}

func (this *OrdersController) create(w http.ResponseWriter, r *http.Request) error {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20) //TODO: understand this

	var req dtos.CreateOrderRequest
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&req); err != nil {
		return errorhandler.BadRequest("invalid JSON body")
	}

	created, err := this.ordersService.CreateOrder(r.Context(), services.CreateOrderArgs{
		CustomerID:      req.CustomerID,
		Currency:        req.Currency,
		DeliveryAddress: req.DeliveryAddress,
		Items:           req.ToRequestedItems(),
	})
	if err != nil {
		return err
	}

	httpapiutil.WriteJSON(w, http.StatusCreated, presenters.PresentOrder(created))
	return nil
}
