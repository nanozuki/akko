package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

type AkkoServer struct {
	Deps struct {
		DB *gorm.DB
	}
	router *httprouter.Router
}

func NewAkkoServer() (*AkkoServer, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	s := AkkoServer{Deps: struct{ DB *gorm.DB }{DB: db}}
	s.router = httprouter.New()
	s.router.GET("/user/:id", s.GetUserByID)
	s.router.GET("/user", s.ListUsers)
	s.router.POST("/user", s.CreateUser)
	return &s, nil
}

func (s *AkkoServer) Run(addr string) error {
	return http.ListenAndServe(addr, s.router)
}

// GetUserByID get user by id
// #[get(/user/<id>)]
func (s *AkkoServer) GetUserByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("id"))
	user, _ := GetUserByID(r.Context(), s.Deps.DB, id)
	res, _ := json.Marshal(user)
	_, _ = w.Write(res)
}

// #[post(/user), data=<user>:json]
func (s *AkkoServer) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	_ = CreateUser(r.Context(), s.Deps.DB, &user)
}

// #[get(/user?<name>)]
func (s *AkkoServer) ListUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	name := r.URL.Query()["name"]
	users, _ := ListUsers(r.Context(), s.Deps.DB, name)
	res, _ := json.Marshal(users)
	_, _ = w.Write(res)
}
