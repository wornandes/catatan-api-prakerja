package models

type Note struct {
	ID      uint   `gorm:"primaryKey"`
	UserID  uint   `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	User    User   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user"`
}
