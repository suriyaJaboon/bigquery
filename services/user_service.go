package services

import (
	"github.com/suriyajaboon/bigquery/structs"
)

type StructUserService struct {
	*structs.Users
}

type IUserService interface {
	GetUserService() (*[]StructUserService, error)
	GetIdUserService(id int) (*StructUserService, error)
	CreateUserService() (*StructUserService, error)
	UpdateUserService() (*StructUserService, error)
	DeleteUserService(id int) (bool, error)

}

func NewUserService(sus *StructUserService) IUserService {
	return &StructUserService{&structs.Users{sus.UUID, sus.Username, sus.Password}}
}

func (sus *StructUserService) GetUserService() (*[]StructUserService, error) {
	getAll := []StructUserService{
		StructUserService{&structs.Users{sus.UUID, sus.Username, sus.Password}},
		StructUserService{&structs.Users{"2", "username2", "2"}},
	}
	return &getAll, nil
}

func (sus *StructUserService) GetIdUserService(id int) (*StructUserService, error) {
	return nil ,nil
}

func (sus *StructUserService) CreateUserService() (*StructUserService, error) {
	return nil, nil
}

func (sus *StructUserService) UpdateUserService() (*StructUserService, error) {
	return nil, nil
}

func (sus *StructUserService) DeleteUserService(id int) (bool, error) {
	return true, nil
}
