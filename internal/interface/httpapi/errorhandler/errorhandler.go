package errorhandler

import (
	"errors"
	"log"
	"net/http"

	"github.com/mohiramansurovna/orders/internal/core/order"
	"github.com/mohiramansurovna/orders/internal/interface/httpapi/httpapiutil"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

type HTTPError struct {
	Status  int
	Message string
}

func (e *HTTPError) Error() string { return e.Message }

func BadRequest(msg string) error {
	return &HTTPError{Status: http.StatusBadRequest, Message: msg}
}

func Wrap(h HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			handle(w, err)
		}
	}
}

func handle(w http.ResponseWriter, err error) {
	var httpErr *HTTPError

	switch {
	case errors.As(err, &httpErr):
		httpapiutil.WriteError(w, httpErr.Status, httpErr.Message)
	case errors.Is(err, order.ErrInvalidOrder):
		httpapiutil.WriteError(w, http.StatusBadRequest, err.Error())
	case errors.Is(err, order.ErrItemNotFound):
		httpapiutil.WriteError(w, http.StatusUnprocessableEntity, err.Error())
	case errors.Is(err, order.ErrInsufficientStock):
		httpapiutil.WriteError(w, http.StatusConflict, err.Error())
	case errors.Is(err, order.ErrNotFound):
		httpapiutil.WriteError(w, http.StatusNotFound, err.Error())
	default:
		log.Printf("unexpected error: %v", err) // details only in logs
		httpapiutil.WriteError(w, http.StatusInternalServerError, "internal error")
	}
}
