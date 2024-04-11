package model

type User struct {
	ID    int    `json:"id" gorm:"type:uuid;primary_key;autoIncrement"`
	Name  string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"not null"`
}
