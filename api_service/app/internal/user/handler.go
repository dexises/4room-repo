package user

import (
	"net/http"

	"github.com/dexises/4room/api_service/internal/handlers"
)

const (
	usersURL = "/users"
	userURL  = "/users/:id"
)

type handler struct{}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *http.ServeMux) {
	router.HandleFunc(usersURL, h.GetUsersList)
	router.HandleFunc(usersURL, h.CreateUser)
	router.HandleFunc(userURL, h.GetUserByID)
	router.HandleFunc(userURL, h.UpdateUser)
	router.HandleFunc(userURL, h.DeleteUser)
}

func (h *handler) GetUsersList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("list of users"))
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create user"))
}

func (h *handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user by id"))
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update user"))
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete user"))
}
