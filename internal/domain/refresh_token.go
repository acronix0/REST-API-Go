package domain

import "time"

type RefreshToken struct {
	ID         int `gorm:"primarykey"`
	UserID     int
	Token      string `gorm:"unigue"`
	DeviceInfo string
	ExpiresAt  time.Time
	CreatedAt  time.Time
}
