package usecases

import (
	"github.com/Imomali1/todo-app/internal/models"
	"github.com/Imomali1/todo-app/internal/repository"
)

type Authorization interface {
	Register(models.RegisterRequest) (models.RegisterResponse, error)
	Login(models.LoginRequest) (models.LoginResponse, error)
	Logout(models.LogoutRequest) (models.LogoutResponse, error)
}

type TodoList interface {
	CreateList(models.CreateListRequest) (models.CreateListResponse, error)
	GetLists(models.GetListsRequest) (models.GetListsResponse, error)
	GetListById(models.GetListByIdRequest) (models.GetListByIdResponse, error)
	UpdateList(models.UpdateListRequest) (models.UpdateListResponse, error)
	DeleteList(models.DeleteListRequest) (models.DeleteListResponse, error)
}

type TodoItem interface {
	CreateItem(models.CreateItemRequest) (models.CreateItemResponse, error)
	GetItems(models.GetItemsRequest) (models.GetItemsResponse, error)
	GetItemById(models.GetItemByIdRequest) (models.GetItemByIdResponse, error)
	UpdateItem(models.UpdateItemRequest) (models.UpdateItemResponse, error)
	DeleteItem(models.DeleteItemRequest) (models.DeleteItemResponse, error)
}

type Services struct {
	Authorization
	TodoList
	TodoItem
}

func NewServices(repos *repository.Repository) *Services {
	return &Services{
		Authorization: NewAuthService(&repos.Authorization),
		TodoList:      NewTodoListService(&repos.TodoListRepository),
		TodoItem:      NewTodoItemService(&repos.TodoItemRepository),
	}
}
