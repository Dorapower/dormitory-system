package model

import "time"

type Logs struct {
	ID        int `gorm:"primaryKey;autoIncrement"`
	uid       int
	Operation string
	ClientIp  string
	CreatTime time.Time
	Content   string
	Status    int
	IsDel     int `gorm:"default:0"`
}
