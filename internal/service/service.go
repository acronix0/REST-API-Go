package service

import (
	"context"
	"mime/multipart"
	"time"

	pb "github.com/acronix0/REST-API-Go-protos/gen/go/auth"
	"github.com/acronix0/REST-API-Go/internal/domain"
	"github.com/acronix0/REST-API-Go/internal/kafka"
	"github.com/acronix0/REST-API-Go/internal/repository"
	"github.com/acronix0/REST-API-Go/pkg/auth"
	"github.com/acronix0/REST-API-Go/pkg/hash"
)

type UserRegisterInput struct {
    Name     string `json:"name" validate:"required"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=8"`
    Phone    string `json:"phone" validate:"required"`
}

type UserLoginInput struct {
	Email    string
	Password string
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}
type UpdateUserInput struct {
    ID    int     `json:"id" validate:"required"`
    Name  *string `json:"name" validate:"omitempty,min=3"`
    Email *string `json:"email" validate:"omitempty,email"`
    Phone *string `json:"phone" validate:"omitempty,min=9"`
}
type Users interface {
	SignUp(ctx context.Context, input UserRegisterInput, deviceinfo string, role string) (Tokens, error)
	SignIn(ctx context.Context, input UserLoginInput, deviceInfo string) (Tokens, error)
	GetByID(ctx context.Context, id int) (domain.User, error)
	GetUsers(ctx context.Context) ([]domain.User, error)
	GetUserRole(ctx context.Context,userID int) (string, error)
	RefreshTokens(ctx context.Context, refreshToken string, deviceInfo string) (Tokens, error)
	ChangePassword(ctx context.Context, id int, newPassword string) error
	Block(ctx context.Context, userID int) error
	Unblock(ctx context.Context, userID int) error
	DeleteAllRefreshTokens(ctx context.Context, userID int) error
	UpdateProfile(ctx context.Context, id int, input UpdateUserInput) error
}
type Categories interface{
	GetCategories(ctx context.Context) ([]domain.Category, error)
}
type Products interface{
	GetProducts(ctx context.Context) ([]domain.Product, error)
	GetByCredentials(ctx context.Context, query domain.GetProductsQuery) ([]domain.Product, error)
}
type Orders interface {
	Create(ctx context.Context, orderInput CreateOrderInput) error
	GetByUserId(ctx context.Context, userId int) ([]domain.Order, error)
}
type Imports interface {
	Parse(file1, file2 *multipart.File) error
	ImportPicture(file *multipart.File) error
}

type ProductInput struct {
	ID       int     `json:"product_id"`
	Article  string  `json:"product_article"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
	Image    string  `json:"image"`
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

 type Deps struct{
	Repos *repository.Repositories
	TokenManager auth.TokenManager
	Hasher       hash.PasswordHasher
	AccessTokenTTL   time.Duration
	RefreshTokenTTL time.Duration
	AuthClient  pb.AuthClient
	KafkaProducer *kafka.KafkaProducer
	
}

func NewServices(deps Deps) (ServiceManager, error){
	importService := NewImportsService(deps.Repos.Category, deps.Repos.Product)
	userService := NewUsersService( deps.AuthClient,deps.Hasher, deps.Repos.User, deps.Repos.Auth, deps.AccessTokenTTL, deps.RefreshTokenTTL)
	productService := NewProductsService(deps.Repos.Product)
	orderService := NewOrdersService(deps.Repos.Order)
	categoriesService := NewCategoriesService(deps.Repos.Category, deps.KafkaProducer)
	return &services{
		users: userService, 
		categories: categoriesService,
		products: productService,
    orders: orderService,
		imports: importService,
	}, nil
}


type ServiceManager interface {
	Categories() Categories
	Products() Products
	Orders() Orders
	Users() Users
	Imports() Imports
}
type services struct{
	categories Categories
	products Products
	orders Orders
	users Users
	imports Imports
}

func (s *services) Categories() Categories {
	return s.categories
}

func (s *services) Products() Products {
	return s.products
}

func (s *services) Orders() Orders {
	return s.orders
}

func (s *services) Users() Users {

	return s.users
}

func (s *services) Imports() Imports {
	return s.imports
}