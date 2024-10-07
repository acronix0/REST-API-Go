package domain

type SearchQuery struct {
	Search string `form:"search"`
}
type PaginationQuery struct {
	Skip  int64 `form:"skip"`
	Limit int64 `form:"limit"`
}

type SortedType int

const (
	Name = iota
	PriceUp
	PriceDown
	Count
)

type ProductFiltersQuery struct {
	SortedType SortedType `json:"sorted_type"`
	MinPrice   int        `json:"min_price"`
	MaxPrice   int        `json:"max_price"`
	InStock    bool       `json:"in_stock"`
}
type GetProductsQuery struct {
	SearchQuery
	PaginationQuery
	ProductFiltersQuery
}
