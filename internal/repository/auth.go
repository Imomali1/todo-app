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
	return r.Register(req)
}

func (r *AuthRepos) Login(req models.LoginRequest) (models.LoginResponse, error) {
	return r.Login(req)
}

func (r *AuthRepos) Logout(req models.LogoutRequest) (models.LogoutResponse, error) {
	return r.Logout(req)
}
