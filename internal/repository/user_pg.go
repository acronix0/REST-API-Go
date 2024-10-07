package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/acronix0/REST-API-Go/internal/domain"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepo {
	return &userRepo{db: db}
}
func (r *userRepo) Login(ctx context.Context, email string, passwordHash string) (domain.User, error){
	var user domain.User
	if err := r.db.QueryRowContext(
		ctx,
		"SELECT id, name, email, password, blocked, role FROM users WHERE email = $1 AND password = $2",
		email,
		passwordHash,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Blocked,
    &user.Role,
	); err != nil {
		return user, err
	}
	return user, nil
}
func (r *userRepo) GetUsers(ctx context.Context) ([]domain.User, error) {
	rows, err := r.db.QueryContext(
		ctx,
		"SELECT id, name, email, password FROM users",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepo) GetById(ctx context.Context, id int) (domain.User, error) {
	var user domain.User
	if err := r.db.QueryRowContext(
		ctx,
		"SELECT id, name, email, password FROM users WHERE id = $1",
		id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
	); err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User
	if err := r.db.QueryRowContext(
		ctx,
		"SELECT id, name, email, password FROM users WHERE email = $1",
		email,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
	); err != nil {
		if err == sql.ErrNoRows {
			return user, domain.ErrUserNotFound
		}
		return user, err
	}
	return user, nil
}

func (r *userRepo) GetByCredentials(ctx context.Context, email string, passwordHash string) (domain.User, error){
	var user domain.User
  if err := r.db.QueryRowContext(
    ctx,
    "SELECT id, name, email, password FROM users WHERE email = $1 AND password = $2",
    email,
    passwordHash,
  ).Scan(
    &user.ID,
    &user.Name,
    &user.Email,
    &user.Password,
  ); err!= nil {
    if err == sql.ErrNoRows {
      return user, domain.ErrUserNotFound
    }
    return user, err
  }
  return user, nil
}

func (r *userRepo) Create(ctx context.Context, user *domain.User) error {
	return r.db.QueryRowContext(
		ctx,
		"INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id",
		user.Name,
		user.Email,
		user.Password,
	).Scan(&user.ID)
}

func (r *userRepo) Update(ctx context.Context, user UpdateUserInput) error {
  query := "UPDATE users SET "
  values := []interface{}{}
  placeholders := []string{}
  placeholderCount := 0

  if user.Name != nil {
    placeholderCount += 1
    placeholders = append(placeholders, fmt.Sprintf("name = $%d", placeholderCount))
    values = append(values, user.Name)
  }

  if user.Email != nil {
    placeholderCount += 1
    placeholders = append(placeholders, fmt.Sprintf("email = $%d", placeholderCount))
    values = append(values, user.Email)
  }

  if user.Phone != nil {
    placeholderCount += 1
    placeholders = append(placeholders, fmt.Sprintf("phone = $%d", placeholderCount))
    values = append(values, user.Phone)
  }

  if user.Blocked != nil {
    placeholderCount += 1
    placeholders = append(placeholders, fmt.Sprintf("blocked = $%d", placeholderCount))
    values = append(values, user.Blocked)
  }

  if len(placeholders) == 0 {
    return fmt.Errorf("no fields to update")
  }

  query += strings.Join(placeholders, ", ")
  placeholderCount += 1
  query += fmt.Sprintf(" WHERE id = $%d", placeholderCount)
  values = append(values, user.ID)

  _, err := r.db.ExecContext(ctx, query, values...)

  return err
}
func (r *userRepo) ChangePassword(ctx context.Context, userId int, newPassword string) error {
	 _, err := r.db.ExecContext(
		ctx,
	  "UPDATE users SET password = $1 WHERE id = $2",
    newPassword,
    userId,
  )
	return err
}
func (r *userRepo) GetRoleByUserID(ctx context.Context, userID int) (string, error) {
    var role string
    query := "SELECT role FROM users WHERE id = $1"
    err := r.db.QueryRowContext(ctx, query, userID).Scan(&role)
    if err != nil {
        return "", err
    }

    return role, nil
}