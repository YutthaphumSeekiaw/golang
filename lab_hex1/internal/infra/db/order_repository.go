package db

import (
	"database/sql"
	orderapp "hex1/internal/app/order"
	"hex1/internal/domain/order"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) orderapp.IOrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Insert(o *order.Order) error {
	_, err := r.db.Exec("INSERT INTO [Order] (Product, Quantity, Price) VALUES (?, ?, ?)", o.Product, o.Quantity, o.Price)
	return err
}

func (r *OrderRepository) GetAll() ([]order.Order, error) {
	rows, err := r.db.Query("SELECT ID, Product, Quantity, Price FROM [Order]")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []order.Order
	for rows.Next() {
		var o order.Order
		if err := rows.Scan(&o.ID, &o.Product, &o.Quantity, &o.Price); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, nil
}

func (r *OrderRepository) Insert2(o *order.Order) error {
	_, err := r.db.Exec("INSERT INTO [Order] (Product, Quantity, Price) VALUES (?, ?, ?)", o.Product, o.Quantity, o.Price)
	return err
}

// Ensure interface compliance
// var _ interface {
// 	Insert(*order.Order) error
// 	GetAll() ([]order.Order, error)
// } = (*OrderRepository)(nil)
