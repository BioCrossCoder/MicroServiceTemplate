package proxy

import (
	"main/models"
)

type AppManagementProxy interface {
	RegisterApp(req *models.PostAppReqVO) (res *models.PostAppResVO, err error)
	CancelApp(req *models.DeleteAppByIdReqVO) (err error)
	GetAppList(req *models.GetAppReqVO) (res *models.GetAppResVO, err error)
}

var amp AppManagementProxy

func ConnectAppManagementProxy() AppManagementProxy {
	return amp
}

func InjectAppManagementService(svc AppManagementProxy) {
	amp = svc
}
