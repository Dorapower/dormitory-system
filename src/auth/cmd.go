package auth

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Type     string `form:"type" json:"type" binding:"required"`
}
