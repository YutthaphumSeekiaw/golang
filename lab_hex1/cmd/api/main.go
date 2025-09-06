package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/exp/slog"

	orderapp "hex1/internal/app/order"
	"hex1/internal/config"
	"hex1/internal/infra/db"
)

func main() {
	cfg := config.LoadConfig()

	logger := slog.New(slog.NewTextHandler(log.Writer(), nil))
	dbConn, err := db.NewSQLServer(cfg.SQLServer)
	if err != nil {
		logger.Error("DB connection failed", "err", err)
		return
	}
	defer dbConn.Close()

	repo := db.NewOrderRepository(dbConn)
	handler := orderapp.NewHandler(repo, logger)

	ccc := handler.ListOrders
	fmt.Printf("ccc: %p\n", ccc)

	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		logger.Info("Health check called")
		return c.String(http.StatusOK, "OK")
	})

	e.GET("/orders", handler.ListOrders)
	e.POST("/orders", handler.CreateOrder)

	logger.Info("Starting server on :8080")
	e.Logger.Fatal(e.Start(":8080"))
}
