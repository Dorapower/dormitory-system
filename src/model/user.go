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
