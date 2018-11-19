package service

import (
	"demo/src/dao"
	"demo/src/model"
)

type user struct{}

var User user

func (user) AddUser(user *model.User) error {
	return dao.AddUser(user)
}

func (user) Delete(id int) error {
	return dao.DeleteUser(id)
}

func (user) ListByPage(page int, pageSize int) ([]model.User, int, error) {
	return dao.ListByPage(page, pageSize)
}
