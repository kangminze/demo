package model


type User struct {
	ID uint `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
}

const (
	TaskStatusProcess = 0
	TaskStatusSucc    = 1
	TaskStatusFail    = 2
)

type Task struct {
	ID uint `json:"id" gorm:"primary_key"`
	Content string `json:"content"`
	status int `json:"status"`
}