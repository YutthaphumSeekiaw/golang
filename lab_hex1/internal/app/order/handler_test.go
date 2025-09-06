package orderapp

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slog"

	domain "hex1/internal/domain/order"
)

type mockRepo struct {
	orders    []domain.Order
	insertErr error
	getAllErr error
}

func (m *mockRepo) Insert(o *domain.Order) error {
	return m.insertErr
}
func (m *mockRepo) GetAll() ([]domain.Order, error) {
	if m.getAllErr != nil {
		return nil, m.getAllErr
	}
	return m.orders, nil
}

func TestListOrders(t *testing.T) {
	e := echo.New()
	repo := &mockRepo{orders: []domain.Order{{ID: 1, Product: "A", Quantity: 2, Price: 10.0}}}
	h := NewHandler(repo, slog.Default())

	req := httptest.NewRequest(http.MethodGet, "/orders", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, h.ListOrders(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "A")
	}
}

func TestCreateOrder(t *testing.T) {
	e := echo.New()
	repo := &mockRepo{}
	h := NewHandler(repo, slog.Default())

	body := `{"product":"B","quantity":3,"price":20}`
	req := httptest.NewRequest(http.MethodPost, "/orders", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, h.CreateOrder(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), "B")
	}
}

func TestCreateOrder_BadPayload(t *testing.T) {
	e := echo.New()
	repo := &mockRepo{}
	h := NewHandler(repo, slog.Default())

	req := httptest.NewRequest(http.MethodPost, "/orders", strings.NewReader("bad json"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := h.CreateOrder(c)
	assert.Error(t, err)
}

func TestListOrders_Error(t *testing.T) {
	e := echo.New()
	repo := &mockRepo{getAllErr: errors.New("db error")}
	h := NewHandler(repo, slog.Default())

	req := httptest.NewRequest(http.MethodGet, "/orders", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h.ListOrders(c)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}
