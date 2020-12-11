package user

import "grpc-server-1/model"

type UserRepo interface {
	AddUser(user *model.User) (*model.User, error)
	FindUserById(id model.UserId) (*model.User, error)
	FindUsers() (*[]model.User, error)
	UpdateUser(user *model.UserUpdate) (*model.User, error)
	DeleteUser(id *model.UserId) error 
}