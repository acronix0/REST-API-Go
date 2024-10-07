package domain

type User struct {
	ID       int   `gorm:"primaryKey" json:"id"`
	Email    string `gorm:"unigue" json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Blocked bool `json:"blocked"`
	Role 		string `json:"role"`
}
