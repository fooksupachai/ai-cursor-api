package mysql

import (
	"database/sql"
	"time"

	"backend-service-api/internal/models"
	"backend-service-api/internal/repositories"
)

var _ repositories.UserRepository = (*UserMySQLRepository)(nil)

type UserMySQLRepository struct {
	db *sql.DB
}

func NewUserMySQLRepository(db *sql.DB) (*UserMySQLRepository, error) {
	repo := &UserMySQLRepository{db: db}
	if err := repo.ensureSchema(); err != nil {
		return nil, err
	}
	return repo, nil
}

func (r *UserMySQLRepository) ensureSchema() error {
	_, err := r.db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id VARCHAR(64) PRIMARY KEY,
			email VARCHAR(255) NOT NULL UNIQUE,
			password_hash VARCHAR(255) NOT NULL,
			name VARCHAR(255) NOT NULL DEFAULT '',
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		)`)
	return err
}

func (r *UserMySQLRepository) GetByEmail(email string) (models.User, bool, error) {
	var u models.User
	row := r.db.QueryRow("SELECT id, email, password_hash, name, created_at, updated_at FROM users WHERE email = ?", email)
	if err := row.Scan(&u.ID, &u.Email, &u.PasswordHash, &u.Name, &u.CreatedAt, &u.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, false, nil
		}
		return models.User{}, false, err
	}
	return u, true, nil
}

func (r *UserMySQLRepository) GetByID(id string) (models.User, bool, error) {
	var u models.User
	row := r.db.QueryRow("SELECT id, email, password_hash, name, created_at, updated_at FROM users WHERE id = ?", id)
	if err := row.Scan(&u.ID, &u.Email, &u.PasswordHash, &u.Name, &u.CreatedAt, &u.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, false, nil
		}
		return models.User{}, false, err
	}
	return u, true, nil
}

func (r *UserMySQLRepository) Create(user models.User) (models.User, error) {
	_, err := r.db.Exec("INSERT INTO users (id, email, password_hash, name, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)", user.ID, user.Email, user.PasswordHash, user.Name, time.Now(), time.Now())
	return user, err
}

func (r *UserMySQLRepository) Update(user models.User) (models.User, error) {
	_, err := r.db.Exec("UPDATE users SET email = ?, password_hash = ?, name = ?, updated_at = ? WHERE id = ?", user.Email, user.PasswordHash, user.Name, time.Now(), user.ID)
	return user, err
}


