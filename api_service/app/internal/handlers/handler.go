package handlers

import (
	"github.com/ynuraddi/router"
)

type Handler interface {
	Register(router *router.Router)
}
