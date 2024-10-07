package domain

type Product struct {
	ID         int     `gorm:"primarykey" json:"id"`
	Article    string  `gorm:"unique"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Image      string  `json:"image"`
	Quantity   int     `json:"quantity"`
	CategoryID int     `json:"categoryId"`
}
