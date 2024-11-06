package usecase

import (
	"crm-admin/internal/entity"
	"crm-admin/internal/usecase/hashing"
	"log/slog"
)

type UserUseCase struct {
	repo UsersRepo
	log  *slog.Logger
}

func NewUserUseCase(repo UsersRepo, log *slog.Logger) *UserUseCase {
	return &UserUseCase{
		repo: repo,
		log:  log,
	}
}

func (u *UserUseCase) RegisterAdmin(in entity.AdminPass) (entity.Message, error) {
	hash, err := hashing.HashPassword(in.Password)
	if err != nil {
		u.log.Error("Error in hashing password", "error", err)
		return entity.Message{}, err
	}

	in.Password = hash

	res, err := u.repo.AddAdmin(in)
	if err != nil {
		u.log.Error("Error in adding admin", "error", err)
		return entity.Message{}, err
	}

	return res, nil
}

func (u *UserUseCase) AddUser(in entity.User) (entity.UserRequest, error) {
	hash, err := hashing.HashPassword(in.Password)
	if err != nil {
		u.log.Error("Error in hashing password", "error", err)
		return entity.UserRequest{}, err
	}

	in.Password = hash

	res, err := u.repo.CreateUser(in)
	if err != nil {
		u.log.Error("Error in adding user", "error", err)
		return entity.UserRequest{}, err
	}

	return res, nil
}

func (u *UserUseCase) GetUser(in entity.UserID) (entity.UserRequest, error) {
	res, err := u.repo.GetUser(in)
	if err != nil {
		u.log.Error("Error in getting user", "error", err)
		return entity.UserRequest{}, err
	}

	return res, nil
}

func (u *UserUseCase) UpdateUser(in entity.UserRequest) (entity.UserRequest, error) {
	res, err := u.repo.UpdateUser(in)
	if err != nil {
		u.log.Error("Error in updating user", "error", err)
		return entity.UserRequest{}, err
	}

	return res, nil
}

func (u *UserUseCase) GetUserList(in entity.FilterUser) (entity.UserList, error) {
	res, err := u.repo.GetListUser(in)
	if err != nil {
		u.log.Error("Error in getting user list", "error", err)
		return entity.UserList{}, err
	}

	return res, nil
}

func (u *UserUseCase) DeleteUser(in entity.UserID) (entity.Message, error) {
	res, err := u.repo.DeleteUser(in)
	if err != nil {
		u.log.Error("Error in deleting user", "error", err)
		return entity.Message{}, err
	}

	return res, nil
}
