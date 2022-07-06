package requests

type Login struct {
	Login    string `form:"login" json:"login,omitempty" validate:"required"`
	Password string `form:"password" json:"password,omitempty" validate:"required"`
}
