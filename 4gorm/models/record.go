package models

type Record struct {
	ID    uint64 `gorm:"primary_key"`
	Name  string
	Money float64
}
