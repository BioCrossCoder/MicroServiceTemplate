package models

type UserObject struct {
	ID   string `json:"id"`
	Type uint8  `json:"type"`
}

type UserCard struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type ItemKey struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AuditLog struct {
	ID       string `json:"id"`
	Level    uint8  `json:"level"`
	OpType   string `json:"op_type"`
	Msg      string `json:"msg"`
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	Date     uint64 `json:"date"`
}
