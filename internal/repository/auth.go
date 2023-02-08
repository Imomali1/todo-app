package repository

import (
	"database/sql"
	"github.com/Imomali1/todo-app/internal/models"
)

type AuthRepos struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepos {
	return &AuthRepos{
		db: db,
	}
}

func (r *AuthRepos) Register(req models.RegisterRequest) (models.RegisterResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *AuthRepos) Login(req models.LoginRequest) (models.LoginResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *AuthRepos) Logout(req models.LogoutRequest) (models.LogoutResponse, error) {
	//TODO implement me
	panic("implement me")
}
