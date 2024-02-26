package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/nanozuki/akko/examples/todolist/service"
)

type Server struct {
	service *service.Service
	mux     *http.ServeMux
}

func NewServer(service *service.Service) *Server {
	s := &Server{
		service: service,
	}

	s.mux.HandleFunc("GET /v1/todos", s.GetUserTodos)
	s.mux.HandleFunc("GET /v1/todos/{id}", s.GetTodoByID)
	s.mux.HandleFunc("POST /v1/todos", s.AddTodo)
	s.mux.HandleFunc("DELETE /v1/todos/{id}", s.DeleteTodo)
	s.mux.HandleFunc("PATCH /v1/todos/{id}", s.PatchTodo)
	s.mux.HandleFunc("GET /v1/users", s.GetUser)
	s.mux.HandleFunc("PATCH /v1/users/{id}", s.PatchUser)

	return s
}

func (s *Server) ListenAndServe(address string) error {
	return http.ListenAndServe(address, s.mux)
}

func (s *Server) GetUserTodos(w http.ResponseWriter, r *http.Request) {
	user, err := service.LoadUserByToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // TODO: error code
		return
	}
	withCompleted := r.URL.Query().Has("with_completed")
	ret, err := s.service.GetUserTodos(r.Context(), user, withCompleted)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(ret); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) GetTodoByID(w http.ResponseWriter, r *http.Request) {
	user, err := service.LoadUserByToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // TODO: error code
		return
	}
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid params 'id': %v", id), http.StatusBadRequest)
		return
	}
	ret, err := s.service.GetTodoByID(r.Context(), user, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(ret); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) AddTodo(w http.ResponseWriter, r *http.Request) {
	user, err := service.LoadUserByToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // TODO: error code
		return
	}
	var todo service.TodoInput
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ret, err := s.service.AddTodo(r.Context(), user, todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(ret); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	user, err := service.LoadUserByToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // TODO: error code
		return
	}
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid params 'id': %v", id), http.StatusBadRequest)
		return
	}
	if err := s.service.DeleteTodo(r.Context(), user, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) PatchTodo(w http.ResponseWriter, r *http.Request) {
	user, err := service.LoadUserByToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // TODO: error code
		return
	}
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid params 'id': %v", id), http.StatusBadRequest)
		return
	}
	var patch service.TodoPatch
	if err := json.NewDecoder(r.Body).Decode(&patch); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	todo, err := s.service.PatchTodo(r.Context(), user, id, patch)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	user, err := service.LoadUserByToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // TODO: error code
		return
	}
	user, err = s.service.GetUser(r.Context(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) PatchUser(w http.ResponseWriter, r *http.Request) {
	user, err := service.LoadUserByToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // TODO: error code
		return
	}
	var patch service.UserPatch
	if err := json.NewDecoder(r.Body).Decode(&patch); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user, err = s.service.PatchUser(r.Context(), user, patch)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
