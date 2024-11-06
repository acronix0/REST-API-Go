// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/service/service.go

// Package mocks is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	domain "github.com/acronix0/REST-API-Go/internal/domain"
	service "github.com/acronix0/REST-API-Go/internal/service"
	gomock "github.com/golang/mock/gomock"
)

// MockUsers is a mock of Users interface.
type MockUsers struct {
	ctrl     *gomock.Controller
	recorder *MockUsersMockRecorder
}

// MockUsersMockRecorder is the mock recorder for MockUsers.
type MockUsersMockRecorder struct {
	mock *MockUsers
}

// NewMockUsers creates a new mock instance.
func NewMockUsers(ctrl *gomock.Controller) *MockUsers {
	mock := &MockUsers{ctrl: ctrl}
	mock.recorder = &MockUsersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsers) EXPECT() *MockUsersMockRecorder {
	return m.recorder
}

// Block mocks base method.
func (m *MockUsers) Block(ctx context.Context, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Block", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Block indicates an expected call of Block.
func (mr *MockUsersMockRecorder) Block(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Block", reflect.TypeOf((*MockUsers)(nil).Block), ctx, userID)
}

// ChangePassword mocks base method.
func (m *MockUsers) ChangePassword(ctx context.Context, id int, newPassword string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", ctx, id, newPassword)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangePassword indicates an expected call of ChangePassword.
func (mr *MockUsersMockRecorder) ChangePassword(ctx, id, newPassword interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockUsers)(nil).ChangePassword), ctx, id, newPassword)
}

// DeleteAllRefreshTokens mocks base method.
func (m *MockUsers) DeleteAllRefreshTokens(ctx context.Context, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllRefreshTokens", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllRefreshTokens indicates an expected call of DeleteAllRefreshTokens.
func (mr *MockUsersMockRecorder) DeleteAllRefreshTokens(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllRefreshTokens", reflect.TypeOf((*MockUsers)(nil).DeleteAllRefreshTokens), ctx, userID)
}

// GetByID mocks base method.
func (m *MockUsers) GetByID(ctx context.Context, id int) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockUsersMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockUsers)(nil).GetByID), ctx, id)
}

// GetUserRole mocks base method.
func (m *MockUsers) GetUserRole(ctx context.Context, userID int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserRole", ctx, userID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserRole indicates an expected call of GetUserRole.
func (mr *MockUsersMockRecorder) GetUserRole(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserRole", reflect.TypeOf((*MockUsers)(nil).GetUserRole), ctx, userID)
}

// GetUsers mocks base method.
func (m *MockUsers) GetUsers(ctx context.Context) ([]domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", ctx)
	ret0, _ := ret[0].([]domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockUsersMockRecorder) GetUsers(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUsers)(nil).GetUsers), ctx)
}

// RefreshTokens mocks base method.
func (m *MockUsers) RefreshTokens(ctx context.Context, refreshToken, deviceInfo string) (service.Tokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshTokens", ctx, refreshToken, deviceInfo)
	ret0, _ := ret[0].(service.Tokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshTokens indicates an expected call of RefreshTokens.
func (mr *MockUsersMockRecorder) RefreshTokens(ctx, refreshToken, deviceInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshTokens", reflect.TypeOf((*MockUsers)(nil).RefreshTokens), ctx, refreshToken, deviceInfo)
}

// SignIn mocks base method.
func (m *MockUsers) SignIn(ctx context.Context, input service.UserLoginInput, deviceInfo string) (service.Tokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignIn", ctx, input, deviceInfo)
	ret0, _ := ret[0].(service.Tokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignIn indicates an expected call of SignIn.
func (mr *MockUsersMockRecorder) SignIn(ctx, input, deviceInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockUsers)(nil).SignIn), ctx, input, deviceInfo)
}

// SignUp mocks base method.
func (m *MockUsers) SignUp(ctx context.Context, input service.UserRegisterInput, deviceinfo, role string) (service.Tokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", ctx, input, deviceinfo, role)
	ret0, _ := ret[0].(service.Tokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignUp indicates an expected call of SignUp.
func (mr *MockUsersMockRecorder) SignUp(ctx, input, deviceinfo, role interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockUsers)(nil).SignUp), ctx, input, deviceinfo, role)
}

// Unblock mocks base method.
func (m *MockUsers) Unblock(ctx context.Context, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unblock", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unblock indicates an expected call of Unblock.
func (mr *MockUsersMockRecorder) Unblock(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unblock", reflect.TypeOf((*MockUsers)(nil).Unblock), ctx, userID)
}

// UpdateProfile mocks base method.
func (m *MockUsers) UpdateProfile(ctx context.Context, id int, input service.UpdateUserInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfile", ctx, id, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProfile indicates an expected call of UpdateProfile.
func (mr *MockUsersMockRecorder) UpdateProfile(ctx, id, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfile", reflect.TypeOf((*MockUsers)(nil).UpdateProfile), ctx, id, input)
}

// MockCategories is a mock of Categories interface.
type MockCategories struct {
	ctrl     *gomock.Controller
	recorder *MockCategoriesMockRecorder
}

// MockCategoriesMockRecorder is the mock recorder for MockCategories.
type MockCategoriesMockRecorder struct {
	mock *MockCategories
}

// NewMockCategories creates a new mock instance.
func NewMockCategories(ctrl *gomock.Controller) *MockCategories {
	mock := &MockCategories{ctrl: ctrl}
	mock.recorder = &MockCategoriesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCategories) EXPECT() *MockCategoriesMockRecorder {
	return m.recorder
}

// GetCategories mocks base method.
func (m *MockCategories) GetCategories(ctx context.Context) ([]domain.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategories", ctx)
	ret0, _ := ret[0].([]domain.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategories indicates an expected call of GetCategories.
func (mr *MockCategoriesMockRecorder) GetCategories(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategories", reflect.TypeOf((*MockCategories)(nil).GetCategories), ctx)
}

// MockProducts is a mock of Products interface.
type MockProducts struct {
	ctrl     *gomock.Controller
	recorder *MockProductsMockRecorder
}

// MockProductsMockRecorder is the mock recorder for MockProducts.
type MockProductsMockRecorder struct {
	mock *MockProducts
}

// NewMockProducts creates a new mock instance.
func NewMockProducts(ctrl *gomock.Controller) *MockProducts {
	mock := &MockProducts{ctrl: ctrl}
	mock.recorder = &MockProductsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProducts) EXPECT() *MockProductsMockRecorder {
	return m.recorder
}

// GetByCredentials mocks base method.
func (m *MockProducts) GetByCredentials(ctx context.Context, query domain.GetProductsQuery) ([]domain.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByCredentials", ctx, query)
	ret0, _ := ret[0].([]domain.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByCredentials indicates an expected call of GetByCredentials.
func (mr *MockProductsMockRecorder) GetByCredentials(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByCredentials", reflect.TypeOf((*MockProducts)(nil).GetByCredentials), ctx, query)
}

// GetProducts mocks base method.
func (m *MockProducts) GetProducts(ctx context.Context) ([]domain.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProducts", ctx)
	ret0, _ := ret[0].([]domain.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProducts indicates an expected call of GetProducts.
func (mr *MockProductsMockRecorder) GetProducts(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProducts", reflect.TypeOf((*MockProducts)(nil).GetProducts), ctx)
}

// MockOrders is a mock of Orders interface.
type MockOrders struct {
	ctrl     *gomock.Controller
	recorder *MockOrdersMockRecorder
}

// MockOrdersMockRecorder is the mock recorder for MockOrders.
type MockOrdersMockRecorder struct {
	mock *MockOrders
}

// NewMockOrders creates a new mock instance.
func NewMockOrders(ctrl *gomock.Controller) *MockOrders {
	mock := &MockOrders{ctrl: ctrl}
	mock.recorder = &MockOrdersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrders) EXPECT() *MockOrdersMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockOrders) Create(ctx context.Context, orderInput service.CreateOrderInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, orderInput)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockOrdersMockRecorder) Create(ctx, orderInput interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOrders)(nil).Create), ctx, orderInput)
}

// GetByUserId mocks base method.
func (m *MockOrders) GetByUserId(ctx context.Context, userId int) ([]domain.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUserId", ctx, userId)
	ret0, _ := ret[0].([]domain.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUserId indicates an expected call of GetByUserId.
func (mr *MockOrdersMockRecorder) GetByUserId(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUserId", reflect.TypeOf((*MockOrders)(nil).GetByUserId), ctx, userId)
}

// MockImports is a mock of Imports interface.
type MockImports struct {
	ctrl     *gomock.Controller
	recorder *MockImportsMockRecorder
}

// MockImportsMockRecorder is the mock recorder for MockImports.
type MockImportsMockRecorder struct {
	mock *MockImports
}

// NewMockImports creates a new mock instance.
func NewMockImports(ctrl *gomock.Controller) *MockImports {
	mock := &MockImports{ctrl: ctrl}
	mock.recorder = &MockImportsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImports) EXPECT() *MockImportsMockRecorder {
	return m.recorder
}

// ImportCategories mocks base method.
func (m *MockImports) ImportCategories(ctx context.Context, categories []domain.Category) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImportCategories", ctx, categories)
	ret0, _ := ret[0].(error)
	return ret0
}

// ImportCategories indicates an expected call of ImportCategories.
func (mr *MockImportsMockRecorder) ImportCategories(ctx, categories interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportCategories", reflect.TypeOf((*MockImports)(nil).ImportCategories), ctx, categories)
}

// ImportProducts mocks base method.
func (m *MockImports) ImportProducts(ctx context.Context, products []domain.Product) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImportProducts", ctx, products)
	ret0, _ := ret[0].(error)
	return ret0
}

// ImportProducts indicates an expected call of ImportProducts.
func (mr *MockImportsMockRecorder) ImportProducts(ctx, products interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportProducts", reflect.TypeOf((*MockImports)(nil).ImportProducts), ctx, products)
}

// MockServiceManager is a mock of ServiceManager interface.
type MockServiceManager struct {
	ctrl     *gomock.Controller
	recorder *MockServiceManagerMockRecorder
}

// MockServiceManagerMockRecorder is the mock recorder for MockServiceManager.
type MockServiceManagerMockRecorder struct {
	mock *MockServiceManager
}

// NewMockServiceManager creates a new mock instance.
func NewMockServiceManager(ctrl *gomock.Controller) *MockServiceManager {
	mock := &MockServiceManager{ctrl: ctrl}
	mock.recorder = &MockServiceManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceManager) EXPECT() *MockServiceManagerMockRecorder {
	return m.recorder
}

// Categories mocks base method.
func (m *MockServiceManager) Categories() service.Categories {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Categories")
	ret0, _ := ret[0].(service.Categories)
	return ret0
}

// Categories indicates an expected call of Categories.
func (mr *MockServiceManagerMockRecorder) Categories() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Categories", reflect.TypeOf((*MockServiceManager)(nil).Categories))
}

// Imports mocks base method.
func (m *MockServiceManager) Imports() service.Imports {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Imports")
	ret0, _ := ret[0].(service.Imports)
	return ret0
}

// Imports indicates an expected call of Imports.
func (mr *MockServiceManagerMockRecorder) Imports() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Imports", reflect.TypeOf((*MockServiceManager)(nil).Imports))
}

// Orders mocks base method.
func (m *MockServiceManager) Orders() service.Orders {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Orders")
	ret0, _ := ret[0].(service.Orders)
	return ret0
}

// Orders indicates an expected call of Orders.
func (mr *MockServiceManagerMockRecorder) Orders() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Orders", reflect.TypeOf((*MockServiceManager)(nil).Orders))
}

// Products mocks base method.
func (m *MockServiceManager) Products() service.Products {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Products")
	ret0, _ := ret[0].(service.Products)
	return ret0
}

// Products indicates an expected call of Products.
func (mr *MockServiceManagerMockRecorder) Products() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Products", reflect.TypeOf((*MockServiceManager)(nil).Products))
}

// Users mocks base method.
func (m *MockServiceManager) Users() service.Users {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Users")
	ret0, _ := ret[0].(service.Users)
	return ret0
}

// Users indicates an expected call of Users.
func (mr *MockServiceManagerMockRecorder) Users() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Users", reflect.TypeOf((*MockServiceManager)(nil).Users))
}
