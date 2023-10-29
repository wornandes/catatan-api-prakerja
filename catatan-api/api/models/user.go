package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"-"`
	Notes    []Note `json:"notes"`
}
