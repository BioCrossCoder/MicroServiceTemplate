package api

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	hCtrl     *healthController
	hCtrlOnce sync.Once
)

type healthController struct{}

func newHealthController() *healthController {
	hCtrlOnce.Do(func() {
		hCtrl = &healthController{}
	})
	return hCtrl
}

func (c *healthController) RegisterPrivate(group *gin.RouterGroup) {
	group.GET("/health/readt", c.healthTest)
	group.GET("/health/alive", c.healthTest)
}

func (c *healthController) healthTest(ctx *gin.Context) {}
