package model

type Sys struct {
	ID       int `gorm:"primaryKey;autoIncrement"`
	KeyName  string
	KeyValue string
	IsDel    int `gorm:"default:0"`
	Remarks  string
}
