package model

type ClassRoom struct {
	ID      int `gorm:"primaryKey;autoIncrement"`
	ClassId int
	RoomId  int
	Remarks string
	IsDel   int `gorm:"default:0"`
	Status  int `gorm:"default:0"`
}
