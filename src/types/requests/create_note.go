package requests

type CreateNote struct {
	Title   string `form:"title" json:"title,omitempty"`
	Content string `form:"content" json:"content,omitempty" validate:"required"`
	Author  string `form:"author" json:"author,omitempty" validate:"required"`
}
