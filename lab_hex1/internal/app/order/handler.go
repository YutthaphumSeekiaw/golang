package orderapp

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/exp/slog"

	domain "hex1/internal/domain/order"
)

type IOrderRepository interface {
	Insert(o *domain.Order) error
	GetAll() ([]domain.Order, error)
}

type handler struct {
	Repo   IOrderRepository
	Logger *slog.Logger
}

func NewHandler(repo IOrderRepository, logger *slog.Logger) *handler {
	return &handler{Repo: repo, Logger: logger}
}

func (h *handler) ListOrders(c echo.Context) error {
	orders, err := h.Repo.GetAll()
	if err != nil {
		h.Logger.Error("Failed to get orders", "err", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to get orders"})
	}
	return c.JSON(http.StatusOK, orders)
}

func (h *handler) CreateOrder(c echo.Context) error {
	var o domain.Order
	if err := c.Bind(&o); err != nil {
		h.Logger.Error("Invalid order payload", "err", err)
		c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid payload"})
		return err
	}
	if err := h.Repo.Insert(&o); err != nil {
		h.Logger.Error("Failed to insert order", "err", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create order"})
	}
	return c.JSON(http.StatusCreated, o)
}
