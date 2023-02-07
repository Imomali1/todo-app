package usecases

import (
	"github.com/Imomali1/todo-app/internal/models"
	"github.com/Imomali1/todo-app/internal/repository"
)

type TodoItemService struct {
	repos *repository.TodoItemRepository
}

func NewTodoItemService(repos *repository.TodoItemRepository) *TodoItemService {
	return &TodoItemService{
		repos: repos,
	}
}

func (s *TodoItemService) CreateItem(models.CreateItemRequest) (models.CreateItemResponse, error) {
	// TODO: implement me
	panic("implement me")
}

func (s *TodoItemService) GetItems(models.GetItemsRequest) (models.GetItemsResponse, error) {
	// TODO: implement me
	panic("implement me")
}

func (s *TodoItemService) GetItemById(models.GetItemByIdRequest) (models.GetItemByIdResponse, error) {
	// TODO: implement me
	panic("implement me")
}

func (s *TodoItemService) UpdateItem(models.UpdateItemRequest) (models.UpdateItemResponse, error) {
	// TODO: implement me
	panic("implement me")
}

func (s *TodoItemService) DeleteItem(models.DeleteItemRequest) (models.DeleteItemResponse, error) {
	// TODO: implement me
	panic("implement me")
}
