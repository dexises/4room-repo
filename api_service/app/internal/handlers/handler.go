package handlers

import (
	"github.com/dexises/4room/api_service/pkg/router"
)

type Handler interface {
	Register(router *router.Router)
}
