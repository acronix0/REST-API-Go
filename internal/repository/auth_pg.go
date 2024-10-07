package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/acronix0/REST-API-Go/internal/domain"
)

type authRepo struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *authRepo {
	return &authRepo{db}
}

func (r *authRepo) ValidateRefreshToken(ctx context.Context, userID int, refreshToken string) (bool, error) {
	var token string

	err := r.db.QueryRowContext(
		ctx,
		"SELECT token FROM refresh_tokens WHERE user_id = $1 AND token = $2 AND expires_at > NOW()",
		userID,
		refreshToken,
	).Scan(&token)

	if err == sql.ErrNoRows {
		return false, domain.ErrRTokenNotFound
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
func (r *authRepo) SaveRefreshToken(ctx context.Context, userID int, refreshToken string, expiresAt time.Time, deviceInfo string) error{
	_, err := r.db.ExecContext(
    ctx,
    "INSERT INTO refresh_tokens (token, user_id, expires_at, device_info) VALUES ($1, $2, $3, $4)",
    refreshToken,
    userID,
    expiresAt,
    deviceInfo,
  )
  return err
}
func (r *authRepo) DeleteRefreshToken(ctx context.Context, userID int, deviceInfo string) error {
	_, err := r.db.ExecContext(
		ctx,
		"DELETE FROM refresh_tokens WHERE user_id = $1 AND device_info = $2",
		userID,
		deviceInfo,
	)
	return err
}

func (r *authRepo) DeleteAllRefreshTokens(ctx context.Context, userID int) error {
	_, err := r.db.ExecContext(
		ctx,
		"DELETE FROM refresh_tokens WHERE user_id = $1",
		userID,
	)
	return err
}
