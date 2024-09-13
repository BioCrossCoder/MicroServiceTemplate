package api

import (
	"main/common"
	"main/driveradapters/api/middleware"
	"main/logics"
	"sync"

	"github.com/biocrosscoder/flex/typed/collections/set"
	"github.com/gin-gonic/gin"
)

var (
	appCtrl     *appController
	appCtrlOnce sync.Once
)

type appController struct {
	appMgntSvc logics.AppManagementService
}

func newAppController() *appController {
	appCtrlOnce.Do(func() {
		appCtrl = &appController{
			appMgntSvc: logics.NewAppManagementService(),
		}
	})
	return appCtrl
}

func (h *appController) RegisterPublic(group *gin.RouterGroup) {
	group.GET("/app")
	roler := middleware.PermissionValidator(set.Of(common.UserRole_AppAdmin, common.UserRole_SuperAdmin))
	adminGroup := group.Group("/app", roler)
	adminGroup.PUT("/:id", roler)
	adminGroup.PUT("/:id/switch", roler)
	adminGroup.GET("/config", roler)
}

func (h *appController) RegisterPrivate(group *gin.RouterGroup) {
	group.POST("/app")
	group.DELETE("/app/:id")
}
