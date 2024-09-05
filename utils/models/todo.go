package models

type Todo struct {
	ID    int    `gorm:"primaryKey"`
	Title string `gorm:"not null"`
	Done  bool   `gorm:"default:false"`
}
