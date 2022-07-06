package requests

type Register struct {
	Name     string `form:"name" json:"name,omitempty" validate:"required"`
	Email    string `form:"email" json:"email,omitempty" validate:"required"`
	Login    string `form:"login" json:"login,omitempty" validate:"required"`
	Password string `form:"password" json:"password,omitempty" validate:"required"`
}
