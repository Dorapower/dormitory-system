package model

type Auth struct {
	Aid       int `gorm:"primaryKey"`
	Type      int
	Username  string
	Password  string
	salt      string
	Uid       int
	AddedAt   int
	DeletedAt int
	Remarks   string
	Status    int
}
