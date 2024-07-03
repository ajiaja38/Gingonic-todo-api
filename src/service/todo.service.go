package service

import (
	"learning-gin/src/model"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type TodoService struct {
	DB  *gorm.DB
	Log *logrus.Logger
}

func NewTodoService(db *gorm.DB, log *logrus.Logger) *TodoService {
	return &TodoService{
		DB:  db,
		Log: log,
	}
}

func (s *TodoService) Create(todo *model.Todo) error {
	err := s.DB.Create(todo).Error
	if err != nil {
		s.Log.Errorf("Failed to create todo: %v", err)
		return err
	}
	return nil
}

func (s *TodoService) GetAll() ([]model.Todo, error) {
	var todos []model.Todo
	err := s.DB.Find(&todos).Error
	if err != nil {
		s.Log.Errorf("Failed to fetch todos: %v", err)
		return nil, err
	}
	return todos, nil
}

func (s *TodoService) GetById(id uint) (*model.Todo, error) {
	var todo model.Todo
	err := s.DB.First(&todo, id).Error
	if err != nil {
		s.Log.Errorf("Failed to fetch todo with ID %d: %v", id, err)
		return nil, err
	}
	return &todo, nil
}

func (s *TodoService) Update(todo *model.Todo) error {
	err := s.DB.Save(todo).Error
	if err != nil {
		s.Log.Errorf("Failed to update todo with ID %d: %v", todo.ID, err)
		return err
	}
	return nil
}

func (s *TodoService) Delete(id uint) error {
	err := s.DB.Where("id = ?", id).Delete(model.Todo{}).Error
	if err != nil {
		s.Log.Errorf("Failed to delete todo with ID %d: %v", id, err)
		return err
	}
	return nil
}
