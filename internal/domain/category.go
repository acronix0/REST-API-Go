package domain

type Category struct {
	ID      int    `gorm:"primarykey" json:"id"`
	Article string `gorm:"unique" json:"article"`
	Name    string `json:"name"`
	Image   string `json:"image"`
}
