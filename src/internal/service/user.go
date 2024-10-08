package service

import (
	"github.com/rulanugrh/eirene/src/helper"
	"github.com/rulanugrh/eirene/src/internal/entity"
	"github.com/rulanugrh/eirene/src/internal/middleware"
	"github.com/rulanugrh/eirene/src/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(req entity.UserRegister) (*helper.UserRegister, error)
	Login(req entity.UserLogin) (*helper.UserLogin, error)
	Update(username string, model entity.User) (*helper.User, error)
}

type userservice struct {
	repo     repository.UserRepository
	validate middleware.IValidate
}

func NewUserService(repo repository.UserRepository, validate middleware.IValidate) UserService {
	return &userservice{
		repo:     repo,
		validate: validate,
	}
}

func (u *userservice) Register(req entity.UserRegister) (*helper.UserRegister, error) {
	err := u.validate.Validate(req)
	if err != nil {
		return nil, u.validate.ValidationMessage(err)
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return nil, helper.BadRequest("Cannot generate password")
	}

	modelReq := entity.UserRegister{
		Username: req.Username,
		Email:    req.Email,
		Password: string(password),
	}

	data, err := u.repo.Register(modelReq)
	if err != nil {
		return nil, helper.InternalServerError("sorry cannt create user")
	}

	response := helper.UserRegister{
		Email:    data.Email,
		Username: data.Username,
	}

	return &response, nil
}
func (u *userservice) Login(req entity.UserLogin) (*helper.UserLogin, error) {
	err := u.validate.Validate(req)
	if err != nil {
		return nil, u.validate.ValidationMessage(err)
	}

	data, err := u.repo.Login(req)
	if err != nil {
		return nil, err
	}

	compare := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(req.Password))
	if compare != nil {
		return nil, helper.Unauthorize("cannot compare password")
	}

	claimToken := entity.UserLogin{
		Email:    req.Email,
		Password: req.Password,
		Username: data.Username,
	}

	token, err := middleware.GenerateToken(claimToken)
	if err != nil {
		return nil, helper.BadRequest(err.Error())
	}

	response := helper.UserLogin{
		Token: token,
	}

	return &response, nil
}

func (u *userservice) Update(username string, model entity.User) (*helper.User, error) {
	data, err := u.repo.Update(username, model)
	if err != nil {
		return nil, helper.BadRequest(err.Error())
	}

	return &helper.User{
		Username: data.Username,
		Avatar:   data.Avatar,
		Email:    data.Email,
	}, nil
}
