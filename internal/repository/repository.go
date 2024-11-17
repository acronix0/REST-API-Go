package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/acronix0/REST-API-Go/internal/cache"
	"github.com/acronix0/REST-API-Go/internal/domain"
)
type UpdateUserInput struct {
	ID       int     `json:"id"`
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Phone    *string `json:"phone"`
	Blocked *bool `json:"is_blocked"`
}
type User interface {
	Login(ctx context.Context, email string, passwordHash string) (domain.User, error)
	GetById(ctx context.Context, id int) (domain.User, error)
	GetByEmail(ctx context.Context, email string) (domain.User, error)
	GetByCredentials(ctx context.Context, email string, passwordHash string) (domain.User, error)
	GetUsers(ctx context.Context) ([]domain.User, error)
	GetRoleByUserID(ctx context.Context, userID int) (string, error)
	Create(ctx context.Context, user *domain.User) error
	Update(ctx context.Context, user UpdateUserInput) error
	ChangePassword(ctx context.Context, userId int, newPassword string) error
}

type Auth interface {
	SaveRefreshToken(ctx context.Context, userID int, refreshToken string, expiresAt time.Time, deviceInfo string) error
	ValidateRefreshToken(ctx context.Context, userID int, refreshToken string) (bool, error)
	DeleteRefreshToken(ctx context.Context, userID int, deviceInfo string) error
	DeleteAllRefreshTokens(ctx context.Context, userID int) error
}


type ProductInput struct {
	ID       int     `json:"product_id"`
	Article  string  `json:"product_article"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
	Image    string  `json:"image"`
}
type Product interface {
	GetProducts(ctx context.Context) ([]domain.Product, error)
	GetByCredentials(ctx context.Context, query domain.GetProductsQuery) ([]domain.Product, error)
	GetByArticles(ctx context.Context, articles []string) (map[string]domain.Product, error)
	CreateOrUpdateBatch(ctx context.Context, products []domain.Product) error
}

type Category interface {
	GetCategories(ctx context.Context) ([]domain.Category, error)
	GetByArticles(ctx context.Context, articles []string) (map[string]domain.Category, error)
	CreateOrUpdateBatch(ctx context.Context, category []domain.Category) error
	CreateBatch(ctx context.Context, category []domain.Category) error
	UpdateBatch(ctx context.Context, category []domain.Category) error
}

type CreateOrderInput struct {
	ID             int
	UserID         int            `json:"user_id"`
	Products       []ProductInput `json:"products"`
	TotalPrice     float64        `json:"total_price"`
	DeliveryType   string         `json:"delivery_type"`
	RecipientName  string         `json:"recipient_name"`
	RecipientPhone string         `json:"recipient_phone"`
	RecipientEmail string         `json:"recipient_email"`
	Address        string         `json:"address"`
	Comment        string         `json:"comment"`
}
type Order interface {
	Create(ctx context.Context, order CreateOrderInput) error
	GetByUserId(ctx context.Context, userId int) ([]domain.Order, error)
}
type UserRole interface{
	Get(ctx context.Context, userID int) (string, error)
	Set(ctx context.Context, id int, userRole string)  error
	Delete(ctx context.Context, userID int) error
}
type Repositories struct{
	User User
	Auth Auth
	Product Product
	Category Category
	Order Order
	UserCache UserRole
}
func NewRepositories(db *sql.DB, cache cache.CacheProvider) *Repositories{
	return &Repositories{
    User: NewUserRepository(db),
    Auth: NewAuthRepository(db),
    Product: NewProductRepository(db),
    Category: NewCategoryRepository(db),
    Order: NewOrderRepository(db),
		UserCache: NewUserRoleRepo(cache),
  }
}