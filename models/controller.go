package models

type PostAppReqVO struct {
	Name  string `json:"name"`
	Entry string `json:"entry"`
}

type PostAppResVO struct {
	ID    string `json:"id"`
	Order int    `json:"order"`
}

type GetAppReqVO struct {
	UserID   string
	Language string
}

type GetAppResVO []*AppInfoVO

type AppInfoVO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Entry string `json:"entry"`
}

type PutAppByIdReqVO struct {
	ID     int       `uri:"id"`
	Name   string    `json:"name"`
	Entry  string    `json:"entry"`
	Access *AccessVO `json:"access"`
}

type AccessVO struct {
	Limited bool        `json:"limited"`
	Users   []*UserCard `json:"users"`
}

type GetAppConfigResVO []*AppConfigVO

type AppConfigVO struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Entry         string `json:"entry"`
	Enabled       bool   `json:"enabled"`
	LimitAccess bool   `json:"limit_access"`
}

type PutAppByIdSwitchReqVO struct {
	ID      int  `uri:"id"`
	Enabled bool `json:"enabled"`
}

type DeleteAppByIdReqVO struct {
	ID int `uri:"id"`
}
