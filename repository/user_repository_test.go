package repository

import (
	"database/sql"
	"rea_games/entity"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func setupMockDBUSER(t *testing.T) (*sql.DB, sqlmock.Sqlmock, *UserRepository) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	repo := &UserRepository{
		BaseRepository: &BaseRepository{db: db},
	}

	return db, mock, repo
}

func TestCreateUser(t *testing.T) {
	db, mock, repo := setupMockDBUSER(t)
	defer db.Close()

	user := &entity.User{
		Email:        "blah@email.com",
		PasswordHash: "hashed",
	}

	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING user_id, created_at, updated_at`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"user_id", "created_at", "updated_at"}).
				AddRow(7, fixedTime, fixedTime),
		)

	err := repo.CreateUser(user)

	assert.NoError(t, err)
	assert.Equal(t, 7, user.UserId)
	assert.Equal(t, fixedTime, user.CreatedAt)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUserByEmail(t *testing.T) {
	db, mock, repo := setupMockDBUSER(t)
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT user_id, email, password_hash, created_at, updated_at FROM users WHERE email = $1 AND deleted_at IS NULL`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"user_id", "email", "password_hash", "created_at", "updated_at"}).
				AddRow(1, "john.doe@example.com", "hashedpassword1", fixedTime, fixedTime),
		)

	user, err := repo.GetUserByEmail("john.doe@example.com")

	assert.NoError(t, err)
	assert.Equal(t, 1, user.UserId)
	assert.Equal(t, "john.doe@example.com", user.Email)
	assert.Equal(t, "hashedpassword1", user.PasswordHash)
	assert.Equal(t, fixedTime, user.CreatedAt)
	assert.Equal(t, fixedTime, user.UpdatedAt)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllUsers(t *testing.T) {
	db, mock, repo := setupMockDBUSER(t)
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT user_id, email, created_at FROM users WHERE deleted_at IS NULL`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"user_id", "email", "created_at"}).
				AddRow(1, "john.doe@example.com", fixedTime).
				AddRow(2, "jane.smith@example.com", fixedTime),
		)

	user, err := repo.GetAllUsers()

	assert.NoError(t, err)
	assert.Len(t, user, 2)
	assert.Equal(t, "john.doe@example.com", user[0].Email)
	assert.Equal(t, "jane.smith@example.com", user[1].Email)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateUsers(t *testing.T) {
	db, mock, repo := setupMockDBUSER(t)
	defer db.Close()

	user := &entity.User{
		Email:  "blah@email.com",
		UserId: 1,
	}

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE users SET email = $1, updated_at = NOW() WHERE user_id = $2`)).
		WithArgs(user.Email, user.UserId).
		WillReturnResult(sqlmock.NewResult(0, 1))
	err := repo.UpdateUser(user)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteUser(t *testing.T) {
	db, mock, repo := setupMockDBUSER(t)
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE users SET deleted_at = NOW() WHERE user_id = $1`)).
		WithArgs(3).
		WillReturnResult(sqlmock.NewResult(0, 1))
	err := repo.DeleteUser(3)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUserByID(t *testing.T) {
	db, mock, repo := setupMockDBUSER(t)
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT user_id, email, password_hash FROM users WHERE user_id = $1 AND deleted_at IS NULL`)).
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows([]string{"user_id", "email", "password_hash"}).
				AddRow(1, "john.doe@example.com", "hashedpassword1"),
		)

	user, err := repo.GetUserByID(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, user.UserId)
	assert.Equal(t, "john.doe@example.com", user.Email)
	assert.Equal(t, "hashedpassword1", user.PasswordHash)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUserByID_NOTFOUND(t *testing.T) {
	db, mock, repo := setupMockDBUSER(t)
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT user_id, email, password_hash FROM users WHERE user_id = $1 AND deleted_at IS NULL`)).
		WithArgs(999).
		WillReturnRows(
			sqlmock.NewRows([]string{"user_id", "email", "password_hash"}).
				AddRow(1, "john.doe@example.com", "hashedpassword1"),
		)

	user, err := repo.GetUserByID(999)

	assert.NoError(t, err)
	assert.Equal(t, 1, user.UserId)
	assert.Equal(t, "john.doe@example.com", user.Email)
	assert.Equal(t, "hashedpassword1", user.PasswordHash)
	assert.NoError(t, mock.ExpectationsWereMet())
}
