package typesutils

type Student struct {
	Id    int `json:"id"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
	Age   int    `json:"age" validate:"required"`
}

type PageData struct {
	Title    string
	Message  string
	Error    string
	Students []Student
}