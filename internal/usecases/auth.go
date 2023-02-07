package usecases

import (
	"github.com/Imomali1/todo-app/internal/models"
	"github.com/Imomali1/todo-app/internal/repository"
)

type AuthService struct {
	repo *repository.Authorization
}

func NewAuthService(repo *repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(req models.RegisterRequest) (models.RegisterResponse, error) {
	// TODO: implement me
	panic("implement me")
}

func (s *AuthService) Login(req models.LoginRequest) (models.LoginResponse, error) {
	// TODO: implement me
	panic("implement me")
}

func (s *AuthService) Logout(req models.LogoutRequest) (models.LogoutResponse, error) {
	// TODO: implement me
	panic("implement me")
}
