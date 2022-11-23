package model

type Groups struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	Name       string
	InviteCode string
	Describe   string
	IsDel      int `gorm:"default:0"`
	Status     int `gorm:"default:0"`
}
