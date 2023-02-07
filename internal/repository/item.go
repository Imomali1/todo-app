package repository

import (
	"database/sql"
	"github.com/Imomali1/todo-app/internal/models"
)

type TodoItemRepos struct {
	db *sql.DB
}

func NewTodoItemRepository(db *sql.DB) *TodoItemRepos {
	return &TodoItemRepos{
		db: db,
	}
}

func (r *TodoItemRepos) CreateItem(request models.CreateItemRequest) (models.CreateItemResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *TodoItemRepos) GetItems(request models.GetItemsRequest) (models.GetItemsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *TodoItemRepos) GetItemById(request models.GetItemByIdRequest) (models.GetItemByIdResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *TodoItemRepos) UpdateItem(request models.UpdateItemRequest) (models.UpdateItemResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *TodoItemRepos) DeleteItem(request models.DeleteItemRequest) (models.DeleteItemResponse, error) {
	//TODO implement me
	panic("implement me")
}
