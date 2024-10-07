package domain

type OrderProduct struct {
	OrderID    int     `gorm:"primarykey" json:"orderId"`
	Order      Order   `gorm:"foreignKey:OrderID" json:"order"`
	ProductID  int     `gorm:"primarykey" json:"productId"`
	Product    Product `gorm:"foreignKey:ProductID" json:"product"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"totalPrice"`
}
