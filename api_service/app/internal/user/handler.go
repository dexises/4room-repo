package user

import (
	"net/http"

	"github.com/dexises/4room/api_service/internal/handlers"
	"github.com/dexises/4room/api_service/pkg/router"
)

const (
	usersURL = "/api/users"
	userURL  = "/api/user"
)

type handler struct{}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *router.Router) {
	router.GET(usersURL, h.GetUsersList)
	router.POST(usersURL, h.CreateUser)
	router.GET(userURL, h.GetUserByID)
	router.PATCH(userURL, h.UpdateUser)
	router.DELETE(userURL, h.DeleteUser)
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
