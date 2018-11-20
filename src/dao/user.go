package dao

import (
	"demo/src/model"
	"github.com/sevenNt/wzap"
)

func AddUser(user *model.User) error {
	if err := db.Create(user).Error; err != nil {
		wzap.Error("create user error", "error", err.Error(), "user", user)
		return err
	}
	return nil
}

func DeleteUser(id int) error {

	if err := db.Where("ID=?", id).Delete(model.User{}).Error; err != nil {
		wzap.Error("delete user error", "error", err.Error(), "id", id)
		return err
	}
	return nil

}

func ListByPage(page int, pageSize int, content string) ([]model.User, int, error) {
	Db := db
	users := make([]model.User, 0)
	total := 0
	if len(content) > 0 {
		Db = Db.Where("username LIKE ?", "%"+content+"%")
	}
	Db.Model(&model.User{}).Count(&total)
	Db = Db.Limit(pageSize).Offset((page - 1) * pageSize)
	if err := Db.Find(&users).Error; err != nil {
		wzap.Error("delete user error", "error", err.Error(), "page", page, "pageSize", pageSize)
		return nil, 0, err
	}
	return users, total, nil
}
