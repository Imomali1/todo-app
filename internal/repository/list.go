package repository

import (
	"database/sql"
	"github.com/Imomali1/todo-app/internal/models"
)

type TodoListRepos struct {
	db *sql.DB
}

func NewTodoListRepository(db *sql.DB) *TodoListRepos {
	return &TodoListRepos{
		db: db,
	}
}

func (r *TodoListRepos) CreateList(request models.CreateListRequest) (models.CreateListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *TodoListRepos) GetLists(request models.GetListsRequest) (models.GetListsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *TodoListRepos) GetListById(request models.GetListByIdRequest) (models.GetListByIdResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *TodoListRepos) UpdateList(request models.UpdateListRequest) (models.UpdateListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *TodoListRepos) DeleteList(request models.DeleteListRequest) (models.DeleteListResponse, error) {
	//TODO implement me
	panic("implement me")
}
