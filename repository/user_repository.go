package repository

import "rea_games/entity"

type UserRepository struct {
	*BaseRepository
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		BaseRepository: NewBaseRepository(),
	}
}

func (r *UserRepository) CreateUser(user *entity.User) error {
	query := `INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING user_id, created_at, updated_at`
	return r.db.QueryRow(query, user.Email, user.PasswordHash).Scan(&user.User_Id, &user.CreatedAt, &user.UpdatedAt)
}

func (r *UserRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	query := `SELECT user_id, email, password_hash, created_at, updated_at FROM users WHERE email = $1 AND deleted_at IS NULL`
	err := r.db.QueryRow(query, email).Scan(&user.User_Id, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetAllUsers() ([]entity.User, error) {
	query := `SELECT user_id, email, created_at FROM users WHERE deleted_at IS NULL`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.UserId, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) UpdateUser(user *entity.User) error {
	query := `UPDATE users SET email = $1, updated_at = NOW() WHERE user_id = $2`
	_, err := r.db.Exec(query, user.Email, user.UserId)
	return err
}

func (r *UserRepository) DeleteUser(id int) error {
	query := `UPDATE users SET deleted_at = NOW() WHERE user_id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *UserRepository) GetUserByID(id int) (*entity.User, error) {
	var user entity.User
	query := `SELECT user_id, email, password_hash FROM users WHERE user_id = $1 AND deleted_at IS NULL`
	err := r.db.QueryRow(query, id).Scan(&user.UserId, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
