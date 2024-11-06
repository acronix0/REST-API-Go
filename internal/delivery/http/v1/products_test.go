package v1

/* import (
	"testing"

	"github.com/acronix0/REST-API-Go/internal/domain"
	mock_service "github.com/acronix0/REST-API-Go/internal/service/mocks"
)
 */
/* func TestHandler_searchProducts(t *testing.T){
	type mockBehavior func (s *mock_service.MockProducts)
	var test domain.GetProductsQuery = domain.ProductFiltersQuery{Search: "test"}
	testTable := []struct{
		name string
		inputBody string
		inputFilters domain.GetProductsQuery
    mockBehavior mockBehavior
		expectedStatusCode int
		expectedRequestBody string
	}{
		{
			name: "Search products by text", 
			inputBody: `{"search": "test"}`, 
			inputFilters: domain.GetProductsQuery{
				SearchQuery: domain.SearchQuery{Search: "test"},
			}, 
			mockBehavior: func(s *mock_service.MockProducts) {
        s.EXPECT().SearchProducts(domain.SearchQuery{Search: "test"}).Return([]domain.Product{{ID: 1, Article: "test", Name: "Test Product", Price: 100.0, Image: "test.jpg", Quantity: 10, CategoryID: 1}}, nil)
      }, 
			expectedStatusCode: 200, 
			expectedRequestBody: `[{"id":1,"article":"test","name":"Test Product","price":100.0,"image":"test.jpg","quantity":10,"categoryId":1}]`},
    
	}
} */