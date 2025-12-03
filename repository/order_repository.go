package repository

import (
	"rea_games/entity"
)

type OrderRepository struct {
	*BaseRepository
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		BaseRepository: NewBaseRepository(),
	}
}

func (r *OrderRepository) CreateOrder(order *entity.Order) error {
	query := `
		INSERT INTO orders (
			user_id, 
			game_id, 
			order_date, 
			quantity, 
			price
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING order_id, created_at
	`

	return r.db.QueryRow(
		query,
		order.UserID,
		order.GameID,
		order.OrderDate,
		order.Quantity,
		order.Price,
	).Scan(
		&order.OrderID,
		&order.CreatedAt,
	)
}

func (r *OrderRepository) GetAllOrders() ([]entity.Order, error) {
	query := `
		SELECT 
			o.order_id,
			o.user_id,
			o.game_id,
			o.order_date,
			o.quantity,
			o.price,
			o.created_at,
			u.email,
			g.title
		FROM orders o
		JOIN users u ON o.user_id = u.user_id
		JOIN games g ON o.game_id = g.game_id
		WHERE o.deleted_at IS NULL
		ORDER BY o.created_at DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []entity.Order

	for rows.Next() {
		var order entity.Order

		err := rows.Scan(
			&order.OrderID,
			&order.UserID,
			&order.GameID,
			&order.OrderDate,
			&order.Quantity,
			&order.Price,
			&order.CreatedAt,
			&order.UserEmail,
			&order.GameTitle,
		)

		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (r *OrderRepository) GetOrdersByUserID(userID int) ([]entity.Order, error) {
	query := `
		SELECT 
			o.order_id,
			o.game_id,
			o.quantity,
			o.price,
			o.order_date,
			o.created_at,
			g.title
		FROM orders o
		JOIN games g ON o.game_id = g.game_id
		WHERE o.user_id = $1
		AND o.deleted_at IS NULL
		ORDER BY o.created_at DESC
	`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []entity.Order

	for rows.Next() {
		var order entity.Order

		err := rows.Scan(
			&order.OrderID,
			&order.GameID,
			&order.Quantity,
			&order.Price,
			&order.OrderDate,
			&order.CreatedAt,
			&order.GameTitle,
		)

		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (r *OrderRepository) UpdateOrder(order *entity.Order) error {
	query := `
		UPDATE orders
		SET
			game_id    = $1,
			quantity   = $2,
			price      = $3,
			order_date = $4,
			updated_at = NOW()
		WHERE order_id = $5
		AND deleted_at IS NULL
	`

	_, err := r.db.Exec(
		query,
		order.GameID,
		order.Quantity,
		order.Price,
		order.OrderDate,
		order.OrderID,
	)

	return err
}

func (r *OrderRepository) DeleteOrder(id int) error {
	query := `
		UPDATE orders
		SET deleted_at = NOW()
		WHERE order_id = $1
		AND deleted_at IS NULL
	`

	_, err := r.db.Exec(query, id)
	return err
}
