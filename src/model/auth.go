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

// TableName 将表名重写为auth，因为gorm约定的为复数形式即auths
func (Auth) TableName() string {
	return "auth"
}
