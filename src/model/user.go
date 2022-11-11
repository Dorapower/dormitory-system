package model

type User struct {
	Uid         int `gorm:"primaryKey"`
	Name        string
	Gender      int
	Email       string
	Mobile      string
	Type        int
	AddedAt     int
	DeletedAt   int
	DeniedAt    int
	LastLoginAt int
	Remarks     string
	Status      int
}

// TableName 将表名重写为user，因为gorm约定的为复数形式即users
func (User) TableName() string {
	return "user"
}
