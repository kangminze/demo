package service

import (
	"demo/src/dao"
	"demo/src/model"
	"errors"
)

type user struct{}

var User user

func (user) AddUser(user *model.User) error {
	return dao.AddUser(user)
}

func (user) Delete(id int) error {
	return dao.DeleteUser(id)
}

func (user) ListByPage(page int, pageSize int, content string) ([]model.User, int, error) {
	return dao.ListByPage(page, pageSize, content)
}

func (user) Validate(username string, password string) (model.User, error) {
	user, _ := dao.FindByUsername(username)
	if len(user.Username) <= 0 {
		return user, errors.New("该用户不存在或密码不匹配")
	}
	if password != user.Password {
		return user, errors.New("该用户不存在或密码不匹配")
	}
	return user, nil
}
