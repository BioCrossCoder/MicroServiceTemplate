package logics

import (
	"main/common"
	"main/logics/dependency"
	"main/logics/proxy"
	"main/models"
	"sync"
)

type AppManagementService interface {
	RegisterApp(req *models.PostAppReqVO) (res *models.PostAppResVO, err error)
	GetAppList(req *models.GetAppReqVO) (res *models.GetAppResVO, err error)
	UpdateApp(req *models.PutAppByIdReqVO) (err error)
	GetAppConfigs() (res *models.GetAppConfigResVO, err error)
	ToggleAppSwitch(req *models.PutAppByIdSwitchReqVO) (err error)
	CancelApp(req *models.DeleteAppByIdReqVO) (err error)
}

var (
	amSvc     AppManagementService
	amSvcOnce sync.Once
)

type appManagementService struct {
	repo      dependency.AppManagementRepo
	eventLoop common.EventLoop
}

func NewAppManagementService() AppManagementService {
	amSvcOnce.Do(func() {
		svc := &appManagementService{
			repo:      dependency.GetAppManagementRepo(),
			eventLoop: common.GetEventLoop(common.Channel),
		}
		svc.init()
		amSvc = svc
		proxy.InjectAppManagementService(amSvc)
	})
	return amSvc
}

func (s *appManagementService) init() {
	s.eventLoop.AddListener("clear_app", func(payload any) (err error) {
		for _, key := range payload.([]uint64) {
			err = s.CancelApp(&models.DeleteAppByIdReqVO{ID: int(key)})
			if err != nil {
				return
			}
		}
		return
	})
}

func (s *appManagementService) RegisterApp(req *models.PostAppReqVO) (res *models.PostAppResVO, err error) {
	//...
	return
}

func (s *appManagementService) GetAppList(req *models.GetAppReqVO) (res *models.GetAppResVO, err error) {
	//...
	return
}

func (s *appManagementService) UpdateApp(req *models.PutAppByIdReqVO) (err error) {
	//...
	return
}

func (s *appManagementService) GetAppConfigs() (res *models.GetAppConfigResVO, err error) {
	//...
	return
}

func (s *appManagementService) ToggleAppSwitch(req *models.PutAppByIdSwitchReqVO) (err error) {
	//...
	return
}

func (s *appManagementService) CancelApp(req *models.DeleteAppByIdReqVO) (err error) {
	//...
	return
}
