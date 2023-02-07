package usecases

import (
	"github.com/Imomali1/todo-app/internal/models"
	"github.com/Imomali1/todo-app/internal/repository"
)

type TodoListService struct {
	repo *repository.TodoListRepository
}

func NewTodoListService(repo *repository.TodoListRepository) *TodoListService {
	return &TodoListService{
		repo: repo,
	}
}

func (s *TodoListService) CreateList(req models.CreateListRequest) (models.CreateListResponse, error) {
	// TODO: implement me
	panic("implement me")
}

func (s *TodoListService) GetLists(models.GetListsRequest) (models.GetListsResponse, error) {
	// TODO: implement me
	panic("implement me")
}

func (s *TodoListService) GetListById(models.GetListByIdRequest) (models.GetListByIdResponse, error) {
	// TODO: implement me
	panic("implement me")
}

func (s *TodoListService) UpdateList(models.UpdateListRequest) (models.UpdateListResponse, error) {
	// TODO: implement me
	panic("implement me")
}

func (s *TodoListService) DeleteList(models.DeleteListRequest) (models.DeleteListResponse, error) {
	// TODO: implement me
	panic("implement me")
}
