package domain

type Product struct {
	ID         int     `gorm:"primarykey" json:"id"`
	Article    string  `gorm:"unique" xml:"Артикул"`
	Name       string  `json:"name" xml:"Наименование"`
	Price      float64 `json:"price" xml:"ЦенаЗаЕдиницу"`
	Image      string  `json:"image" xml:"Картинка"`
	Quantity   int     `json:"quantity" xml:"Количество"`
	CategoryID int     `json:"categoryId"`
}
