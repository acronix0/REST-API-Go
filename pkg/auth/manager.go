package auth

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenManager interface {
	NewJWT(userId int, ttl time.Duration) (string, error)
	Parse(accessToken string) (int, error)
	NewRefreshToken() (string, error)
}

type Manager struct {
	jwtSecret string
}

func NewManager(jwtSecret string) (*Manager, error) {
	if jwtSecret == "" {
		return nil, fmt.Errorf("empty signing key")
	}

	return &Manager{jwtSecret: jwtSecret}, nil
}

func (m *Manager) Parse(accessToken string) (int, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(m.jwtSecret), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("error get user claims from token")
	}
  id, err := strconv.Atoi(claims["sub"].(string))
    if err != nil {
      return 0, err
    }
	return id, nil
}

func (m *Manager) NewJWT(userId int, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(ttl).Unix(),
		Subject:   strconv.Itoa(userId),
	})
	return token.SignedString([]byte(m.jwtSecret))
}

func (m *Manager) NewRefreshToken() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	if _, err := r.Read(b); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", b), nil
}
