package repo

import (
	"grpc-server-1/model"
	"grpc-server-1/user"

	"github.com/jinzhu/gorm"
)

type UserRepoImpl struct {
	db *gorm.DB
}

func CreateUserRepoImpl(db *gorm.DB) user.UserRepo {
	return &UserRepoImpl{db}
}

func (e *UserRepoImpl) AddUser(user *model.User) (*model.User, error) {
	var userDB = model.UserDB{
		Name:     user.Name,
		Email:    user.Email,
		Alamat:   user.Alamat,
		Password: user.Password,
	}
	err := e.db.Table("user").Save(&userDB).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (e *UserRepoImpl) FindUserById(id model.UserId) (*model.User, error) {
	var user model.User
	err := e.db.Table("user").Where("id = ?", id.Id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (e *UserRepoImpl) FindUsers() (*[]model.User, error) {
	var users []model.User
	err := e.db.Table("user").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (e *UserRepoImpl) UpdateUser(user *model.UserUpdate) (*model.User, error) {
	var us model.User
	err := e.db.Table("user").Where("id = ?", user.Id).First(&us).Update(&user.User).Error
	if err != nil {
		return nil, err
	}
	return &us, nil
}

func (e *UserRepoImpl) DeleteUser(id *model.UserId) error {
	var user model.User
	err := e.db.Table("user").Where("id = ?", id.Id).First(&user).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}
