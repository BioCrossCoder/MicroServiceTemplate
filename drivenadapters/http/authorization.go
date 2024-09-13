package http

import (
	"main/infra"
	"sync"
)

type AuthorizationClient interface {
	ParseToken(token string) (info *TokenInfo, err error)
}

var (
	auth     AuthorizationClient
	authOnce sync.Once
)

type authorizationClient struct {
	httpClient infra.HttpClient
}

func NewAuthorizationClient() AuthorizationClient {
	authOnce.Do(func() {
		auth = &authorizationClient{
			httpClient: infra.NewHttpClient(),
		}
	})
	return auth
}

func (c *authorizationClient) ParseToken(token string) (info *TokenInfo, err error) {
	//...
	return
}

type TokenInfo struct {
	UserID    string `json:"user_id"`
	UserName  string `json:"user_name"`
	UserRoles []int  `json:"user_roles"`
}
