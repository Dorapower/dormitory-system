package model

type Buildings struct {
	ID       int `gorm:"primaryKey;autoIncrement"`
	Name     string
	OrderNum int
	IsValid  int `gorm:"default:1"`
	Remarks  string
	Describe string
	ImageUrl string
	IsDel    int `gorm:"default:0"`
	Status   int `gorm:"default:0"`
}
