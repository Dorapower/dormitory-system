package auth

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type RefreshRequest struct {
	RefreshToken string `form:"refresh_token" json:"refresh_token" binding:"required"`
}
