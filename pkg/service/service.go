package service

import (
	"github.com/stirk1337/awesomeProject/pkg/repository"
	"github.com/stirk1337/awesomeProject/pkg/service/auth"
	"github.com/stirk1337/awesomeProject/pkg/todo"
	"github.com/stirk1337/awesomeProject/pkg/user"
)

type Authorization interface {
	CreateUser(user user.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId int, listId int) (todo.TodoList, error)
	Delete(userId int, listId int) error
	Update(userId int, listId int, input todo.UpdateListInput) error
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: auth.NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList)}
}
