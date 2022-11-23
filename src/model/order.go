package model

type Orders struct {
	ID            int `gorm:"primaryKey;autoIncrement"`
	Uid           int
	GroupId       int `gorm:"default:0"`
	BuildingId    int
	SubmitTime    int
	CreatTime     int
	FinishTime    int `gorm:"default:0"`
	RoomId        int
	ResultContent string
	Remarks       string
	IsDel         int `gorm:"default:0"`
	Status        int `gorm:"default:0"`
}
