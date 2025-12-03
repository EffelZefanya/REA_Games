package repository

import (
	"database/sql"
	"rea_games/entity"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func setupMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock, *OrderRepository) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	repo := &OrderRepository{
		BaseRepository: &BaseRepository{db: db},
	}

	return db, mock, repo
}

var fixedTime = time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC)

func TestCreateOrder(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	order := &entity.Order{
		UserID:       1,
		GameID:       2,
		OrderDate:    fixedTime,
		GameQuantity: 3,
		TotalPrice:   150.50,
	}

	mock.ExpectQuery(regexp.QuoteMeta(`
		INSERT INTO orders (
			user_id, 
			game_id, 
			order_date, 
			game_quantity, 
			total_price
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING order_id, created_at
	`)).
		WithArgs(order.UserID, order.GameID, order.OrderDate, order.GameQuantity, order.TotalPrice).
		WillReturnRows(
			sqlmock.NewRows([]string{"order_id", "created_at"}).
				AddRow(10, fixedTime),
		)

	err := repo.CreateOrder(order)

	assert.NoError(t, err)
	assert.Equal(t, 10, order.OrderID)
	assert.Equal(t, fixedTime, order.CreatedAt)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetOrderByID(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT
			order_id,
			user_id,
			game_id,
			game_quantity,
			total_price,
			created_at
		FROM orders
		WHERE order_id = $1
		AND deleted_at IS NULL
	`)).
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"order_id",
				"user_id",
				"game_id",
				"game_quantity",
				"total_price",
				"created_at",
			}).AddRow(1, 2, 3, 4, 199.99, fixedTime),
		)

	order, err := repo.GetOrderByID(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, order.OrderID)
	assert.Equal(t, 2, order.UserID)
	assert.Equal(t, 3, order.GameID)
	assert.Equal(t, 4, order.GameQuantity)
	assert.Equal(t, 199.99, order.TotalPrice)
	assert.Equal(t, fixedTime, order.CreatedAt)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllOrders(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT 
			o.order_id,
			o.user_id,
			o.game_id,
			o.order_date,
			o.game_quantity,
			o.total_price,
			o.created_at,
			u.email,
			g.title
		FROM orders o
		JOIN users u ON o.user_id = u.user_id
		JOIN games g ON o.game_id = g.game_id
		WHERE o.deleted_at IS NULL
		ORDER BY o.created_at DESC
	`)).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"order_id",
				"user_id",
				"game_id",
				"order_date",
				"game_quantity",
				"total_price",
				"created_at",
				"email",
				"title",
			}).
				AddRow(1, 2, 3, fixedTime, 2, 120.00, fixedTime, "user@mail.com", "God of War").
				AddRow(2, 4, 5, fixedTime, 1, 60.00, fixedTime, "another@mail.com", "Elden Ring"),
		)

	orders, err := repo.GetAllOrders()

	assert.NoError(t, err)
	assert.Len(t, orders, 2)
	assert.Equal(t, "God of War", orders[0].GameTitle)
	assert.Equal(t, "Elden Ring", orders[1].GameTitle)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetOrdersByUserID(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT 
			o.order_id,
			o.game_id,
			o.game_quantity,
			o.total_price,
			o.order_date,
			o.created_at,
			g.title
		FROM orders o
		JOIN games g ON o.game_id = g.game_id
		WHERE o.user_id = $1
		AND o.deleted_at IS NULL
		ORDER BY o.created_at DESC
	`)).
		WithArgs(7).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"order_id",
				"game_id",
				"game_quantity",
				"total_price",
				"order_date",
				"created_at",
				"title",
			}).
				AddRow(5, 3, 1, 69.9, fixedTime, fixedTime, "Cyberpunk"),
		)

	orders, err := repo.GetOrdersByUserID(7)

	assert.NoError(t, err)
	assert.Len(t, orders, 1)
	assert.Equal(t, 5, orders[0].OrderID)
	assert.Equal(t, "Cyberpunk", orders[0].GameTitle)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateOrder(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta(`
		UPDATE orders
		SET game_quantity = $1,
		    total_price = $2,
		    updated_at = NOW()
		WHERE order_id = $3
		  AND deleted_at IS NULL
	`)).
		WithArgs(5, 89.99, 1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.UpdateOrder(1, 5, 89.99)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteOrder(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta(`
		UPDATE orders
		SET deleted_at = NOW()
		WHERE order_id = $1
		AND deleted_at IS NULL
	`)).
		WithArgs(3).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.DeleteOrder(3)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetOrderByID_NotFound(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT
			order_id,
			user_id,
			game_id,
			game_quantity,
			total_price,
			created_at
		FROM orders
		WHERE order_id = $1
		AND deleted_at IS NULL
	`)).
		WithArgs(999).
		WillReturnError(sql.ErrNoRows)

	order, err := repo.GetOrderByID(999)

	assert.Nil(t, order)
	assert.Error(t, err)
	assert.Equal(t, sql.ErrNoRows, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

