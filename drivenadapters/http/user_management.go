package http

import (
	"main/infra"
	"main/models"
	"sync"
)

type UserManagementClient interface {
	GetNames(req *GetNamesReq) (res *GetNamesRes, err error)
	GetAccessors(userID string) (accessors []string, err error)
}

var (
	umc     UserManagementClient
	umcOnce sync.Once
)

type userManagementClient struct {
	httpClient infra.HttpClient
}

func NewUserManagementClient() UserManagementClient {
	umcOnce.Do(func() {
		umc = &userManagementClient{
			httpClient: infra.NewHttpClient(),
		}
	})
	return umc
}

func (c *userManagementClient) GetNames(req *GetNamesReq) (res *GetNamesRes, err error) {
	// ...
	return
}

func (c *userManagementClient) GetAccessors(userID string) (accessors []string, err error) {
	// ...
	return
}

type GetNamesReq struct {
	UserIDs       []string `json:"user_ids"`
	DepartmentIDs []string `json:"department_ids"`
	GroupIDs      []string `json:"group_ids"`
}

type GetNamesRes struct {
	UserNames       []*models.ItemKey `json:"user_names"`
	DepartmentNames []*models.ItemKey `json:"department_names"`
	GroupNames      []*models.ItemKey `json:"group_names"`
}
