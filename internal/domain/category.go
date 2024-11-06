package domain

type Category struct {
	ID      int    `gorm:"primarykey" json:"id"`
	Article string `gorm:"unique" json:"article" xml:"Ид"`
	Name    string `json:"name" xml:"Наименование"`
	Image   string `json:"image"`
}
