package domain

import "time"

type Order struct {
	ID             int            `gorm:"unique" json:"id"`
	User           User           `gorm:"foreignKey:UserID" json:"User"`
	UserID         int            `json:"userId"`
	TotalPrice     float64        `json:"totalPrice"`
	OrderDate      time.Time      `json:"orderDate"`
	DeliveryType   string         `json:"deliveryType"`
	RecipientName  string         `json:"recipientName"`
	RecipientPhone string         `json:"recipientPhone"`
	RecipientEmail string         `json:"recipientEmail"`
	Address        string         `json:"address"`
	Comment        string         `json:"comment"`
	Products       []OrderProduct `json:"products"`
}
