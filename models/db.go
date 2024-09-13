package models

type AppPO struct {
	ID          uint64 `gorm:"f_id"`
	Name        string `gorm:"f_name"`
	Entry       string `gorm:"f_entry"`
	Order       uint16 `gorm:"f_order"`
	Enabled     bool   `gorm:"f_enabled"`
	LimitAccess bool   `gorm:"f_limit_access"`
}

type AppUserPO struct {
	ID       uint64 `gorm:"f_id"`
	AppID    uint64 `gorm:"f_app_id"`
	UserID   string `gorm:"f_user_id"`
	UserType uint8  `gorm:"f_user_type"`
}
