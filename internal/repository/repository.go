package repository

import (
	"database/sql"
	"github.com/Imomali1/todo-app/internal/models"
)

type Authorization interface {
	Register(models.RegisterRequest) (models.RegisterResponse, error)
	Login(models.LoginRequest) (models.LoginResponse, error)
	Logout(models.LogoutRequest) (models.LogoutResponse, error)
}

type TodoListRepository interface {
	CreateList(models.CreateListRequest) (models.CreateListResponse, error)
	GetLists(models.GetListsRequest) (models.GetListsResponse, error)
	GetListById(models.GetListByIdRequest) (models.GetListByIdResponse, error)
	UpdateList(models.UpdateListRequest) (models.UpdateListResponse, error)
	DeleteList(models.DeleteListRequest) (models.DeleteListResponse, error)
}

type TodoItemRepository interface {
	CreateItem(models.CreateItemRequest) (models.CreateItemResponse, error)
	GetItems(models.GetItemsRequest) (models.GetItemsResponse, error)
	GetItemById(models.GetItemByIdRequest) (models.GetItemByIdResponse, error)
	UpdateItem(models.UpdateItemRequest) (models.UpdateItemResponse, error)
	DeleteItem(models.DeleteItemRequest) (models.DeleteItemResponse, error)
}

type Repository struct {
	Authorization
	TodoListRepository
	TodoItemRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization:      NewAuthRepository(db),
		TodoListRepository: NewTodoListRepository(db),
		TodoItemRepository: NewTodoItemRepository(db),
	}
}
