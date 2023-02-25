package dto

type UserLoginDTO struct {
	Name     string `form:"name" json:"name" binding:"required" message:"用户名错误" required_err:"用户名不能为空"`
	Password string `form:"password" json:"password" binding:"required" message:"密码不能为空"`
}
