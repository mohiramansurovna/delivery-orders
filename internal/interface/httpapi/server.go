package httpapi

import (
	"net/http"

	"github.com/mohiramansurovna/orders/internal/interface/httpapi/controllers"
)

func NewRouter(db Pinger, orders *controllers.OrdersController) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("GET /health", healthHandler{db: db})
	orders.Routes(mux)
	return mux
}
