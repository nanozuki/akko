package service

import (
	"context"
	"errors"
)

type Todo struct {
	ID        int    `json:"id,omitempty"`
	UserID    int    `json:"user_id,omitempty"`
	Title     string `json:"title,omitempty"`
	Content   string `json:"content,omitempty"`
	Completed bool   `json:"completed,omitempty"`
}

var todos = []*Todo{
	{ID: 1, UserID: 1, Title: "Create example project", Content: "Create an example project for the Akko library.", Completed: false},
	{ID: 2, UserID: 1, Title: "Write documentation", Content: "Write documentation for the Akko library.", Completed: false},
	{ID: 3, UserID: 2, Title: "Create example project", Content: "Create an example project for the Akko library.", Completed: false},
	{ID: 4, UserID: 2, Title: "Write documentation", Content: "Write documentation for the Akko library.", Completed: false},
}

// GetUserTodos returns all of the todos for a user.
// @get /todos?withCompleted+user
func (s *Service) GetUserTodos(ctx context.Context, user *User, withCompleted bool) ([]Todo, error) {
	var userTodos []Todo
	for _, todo := range todos {
		if todo.UserID == user.ID {
			if !withCompleted && todo.Completed {
				continue
			}
			userTodos = append(userTodos, *todo)
		}
	}
	return userTodos, nil
}

// GetTodoByID returns a todo item by its ID.
// @get /todos/id+user
func (s *Service) GetTodoByID(ctx context.Context, user *User, id int) (*Todo, error) {
	for _, todo := range todos {
		if todo.ID == id {
			if todo.UserID != user.ID {
				return nil, errors.New("todo does not belong to user")
			}
			return todo, nil
		}
	}
	return nil, errors.New("todo not found")
}

type TodoInput struct {
	Title     string `json:"title,omitempty"`
	Content   string `json:"content,omitempty"`
	Completed bool   `json:"completed,omitempty"`
}

// Todo represents a todo item.
// @post /todos+user@todo
func (s *Service) AddTodo(ctx context.Context, user *User, todo TodoInput) (Todo, error) {
	newTodo := Todo{
		ID:        len(todos) + 1,
		UserID:    user.ID,
		Title:     todo.Title,
		Content:   todo.Content,
		Completed: todo.Completed,
	}
	todos = append(todos, &newTodo)
	return newTodo, nil
}

// DeleteTodo deletes a todo item by its ID.
// @delete /todos/{id}+user
func (s *Service) DeleteTodo(ctx context.Context, user *User, id int) error {
	for i, todo := range todos {
		if todo.ID == id {
			if todo.UserID != user.ID {
				return errors.New("todo does not belong to user")
			}
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return errors.New("todo not found")
}

type TodoPatch struct {
	Title     *string `json:"title,omitempty"`
	Content   *string `json:"content,omitempty"`
	Completed *bool   `json:"completed,omitempty"`
}

// PatchTodo updates a todo item by its ID.
// @patch /todos/{id}+user@patch
func (s *Service) PatchTodo(ctx context.Context, user *User, id int, patch TodoPatch) (*Todo, error) {
	var todo Todo
	for _, t := range todos {
		if t.ID == id {
			todo = *t
		}
	}
	if todo.UserID != user.ID {
		return nil, errors.New("todo does not belong to user")
	}
	if patch.Title != nil {
		todo.Title = *patch.Title
	}
	if patch.Content != nil {
		todo.Content = *patch.Content
	}
	if patch.Completed != nil {
		todo.Completed = *patch.Completed
	}
	return &todo, nil
}
