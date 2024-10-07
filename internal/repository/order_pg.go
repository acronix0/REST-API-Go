package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/acronix0/REST-API-Go/internal/domain"
)

type orderRepo struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *orderRepo {
	return &orderRepo{db: db}
}

func (r orderRepo) Create(ctx context.Context, orderInput CreateOrderInput) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	err = tx.QueryRowContext(
		ctx,
		`INSERT INTO orders (user_id, total_Price,order_date,delivery_type,recipient_name, recipient_phone, recipient_email,address,comment) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		orderInput.UserID,
		orderInput.TotalPrice,
		time.Now(),
		orderInput.DeliveryType,
		orderInput.RecipientName,
		orderInput.RecipientPhone,
		orderInput.RecipientEmail,
		orderInput.Address,
		orderInput.Comment,
	).Scan(&orderInput.ID)

	if err != nil {
		tx.Rollback()
		return err
	}

	query := "INSERT INTO order_product (order_id, product_id, quantity, total_price) VALUES "
	values := []interface{}{}
	placeholders := []string{}
	for i, p := range orderInput.Products {
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d, $%d)", i*4+1, i*4+2, i*4+3, i*4+4))
		values = append(values, orderInput.ID, p.ID, p.Quantity, p.Price*float64(p.Quantity))

	}
	query += strings.Join(placeholders, ",")
	_, err = tx.ExecContext(ctx, query, values...)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (r orderRepo) GetByUserId(ctx context.Context, userID int) ([]domain.Order, error) {
	query := `
		SELECT 
			o.id AS order_id,
			o.total_price AS order_total_price,
			o.order_date AS order_date,
			o.delivery_type AS delivery_type,
			o.recipient_name AS recipient_name,
			o.recipient_phone AS recipient_phone,
			o.recipient_email AS recipient_email,
			o.address AS address,
			o.comment AS comment,
			op.quantity AS product_quantity,
			op.total_price AS product_total_price,
			p.id AS product_id,
			p.name AS product_name,
			p.article AS product_article,
			p.price AS product_price
		FROM 
			orders o
		JOIN 
			order_product op ON o.id = op.order_id
		JOIN 
			products p ON op.product_id = p.id
		WHERE 
			o.user_id = $1;
	`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ordersMap := make(map[int]*domain.Order)
	var orders []domain.Order

	for rows.Next() {
		var orderID int
		var orderProduct domain.OrderProduct
		var order domain.Order

		err := rows.Scan(
			&orderID,
			&order.TotalPrice,
			&order.OrderDate,
			&order.DeliveryType,
			&order.RecipientName,
			&order.RecipientPhone,
			&order.RecipientEmail,
			&order.Address,
			&order.Comment,
			&orderProduct.Quantity,
			&orderProduct.TotalPrice,
			&orderProduct.ProductID,
			&orderProduct.Product.Name,
			&orderProduct.Product.Article,
			&orderProduct.Product.Price,
		)
		if err != nil {
			return nil, err
		}

		if existingOrder, ok := ordersMap[orderID]; ok {
			existingOrder.Products = append(existingOrder.Products, orderProduct)
		} else {
			order.ID = orderID
			order.Products = []domain.OrderProduct{orderProduct}
			ordersMap[orderID] = &order
		}
	}

	for _, order := range ordersMap {
		orders = append(orders, *order)
	}

	return orders, nil
}
