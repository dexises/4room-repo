package auth

import (
	"net/http"

	"github.com/dexises/4room/api_service/internal/apperror"
	"github.com/dexises/4room/api_service/internal/client/user_service"
	"github.com/dexises/4room/api_service/pkg/logging"
	"github.com/ynuraddi/router"
)

const (
	authURL   = "/api/auth"
	signupURL = "/api/signup"
)

type Handler struct {
	Logger      logging.Logger
	UserService user_service.UserService
}

func (h *Handler) Register(router *router.Router) {
	router.POST(authURL, apperror.Middleware(h.Auth))
	router.PUT(authURL, apperror.Middleware(h.Auth))
	router.POST(signupURL, apperror.Middleware(h.SignUp))
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()
	w.Write([]byte("signup"))
	return nil
}

func (h *Handler) Auth(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("authenticated"))
	return nil
}
