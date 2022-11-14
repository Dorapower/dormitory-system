package addData

type Auth struct {
	//Aid       int `gorm:"primaryKey"`
	Type     int    `form:"type" json:"type"`
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	//Salt      string
	Uid int `form:"uid" json:"uid" binding:"required"`
	//AddedAt   int
	DeletedAt int    `form:"deletedAt" json:"deletedAt"`
	Remarks   string `form:"remarks" json:"remarks"`
	Status    int    `form:"status" json:"status" `
}

type User struct {
	//Uid         int `gorm:"primaryKey"`
	Name   string `form:"name" json:"name" binding:"required"`
	Gender int    `form:"gender" json:"gender" binding:"required"`
	Email  string `form:"email" json:"email"`
	Mobile string `form:"mobile" json:"mobile"`
	Type   int    `form:"type" json:"type"`
	//AddedAt     int
	DeletedAt   int    `form:"deletedAt" json:"deletedAt"`
	DeniedAt    int    `form:"deniedAt" json:"deniedAt"`
	LastLoginAt int    `form:"lastLoginAt" json:"LastLoginAt"`
	Remarks     string `form:"remarks" json:"remarks"`
	Status      int    `form:"status" json:"status"`
}
