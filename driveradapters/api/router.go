package api

import (
	"main/drivenadapters/http"
	"main/driveradapters/api/middleware"
	"sync"

	"github.com/gin-gonic/gin"
)

type PublicController interface {
	RegisterPublic(group *gin.RouterGroup)
}

type PrivateController interface {
	RegisterPrivate(group *gin.RouterGroup)
}

type Router interface {
	RegisterPrivateAPI(group *gin.RouterGroup)
	RegisterPublicAPI(group *gin.RouterGroup)
}

var (
	r     Router
	rOnce sync.Once
)

type router struct {
	publicControllers  []PublicController
	privateControllers []PrivateController
}

func NewRouter() Router {
	rOnce.Do(func() {
		r = &router{
			publicControllers: []PublicController{},
			privateControllers: []PrivateController{
				newHealthController(),
			},
		}
	})
	return r
}

func (r *router) RegisterPrivateAPI(group *gin.RouterGroup) {
	group.Use()
	for _, ctrl := range r.privateControllers {
		ctrl.RegisterPrivate(group)
	}
}

func (r *router) RegisterPublicAPI(group *gin.RouterGroup) {
	authValidator := middleware.TokenParser(http.NewAuthorizationClient())
	group.Use(authValidator)
	for _, ctrl := range r.publicControllers {
		ctrl.RegisterPublic(group)
	}
}
