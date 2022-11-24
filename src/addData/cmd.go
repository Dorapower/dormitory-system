package addData

type Auth struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Uid      int    `form:"uid" json:"uid" binding:"required"`
	Remarks  string `form:"remarks" json:"remarks" binding:"required"`
	IsDel    int    `form:"is_del" json:"is_del"`
	Status   int    `form:"status" json:"status"`
}

type User struct {
	Name          string `form:"name" json:"name" binding:"required"`
	Gender        int    `form:"gender" json:"gender" binding:"required"`
	Email         string `form:"email" json:"email" binding:"required"`
	Tel           string `form:"tel" json:"tel" binding:"required"`
	Type          int    `form:"type" json:"type"`
	IsDeny        int    `form:"is_deny" json:"is_deny"`
	IsDel         int    `form:"is_del" json:"is_del"`
	LastLoginTime int    `form:"last_login_time" json:"last_login_time"`
	Remarks       string `form:"remarks" json:"remarks" binding:"required"`
	Status        int    `form:"status" json:"status"`
}
