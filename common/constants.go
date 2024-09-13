package common

import (
	"github.com/biocrosscoder/flex/typed/collections/arraylist"
	"github.com/biocrosscoder/flex/typed/collections/dict"
)

const (
	LanguageZHCN = "zh_cn"
	LanguageENUS = "en_us"
	User         = "user"
	Department   = "department"
	UserGroup    = "group"
	Channel      = "demo_service"
)

var (
	Languages    = arraylist.Of(LanguageZHCN, LanguageENUS)
	UserTypeTag  = arraylist.Of(User, Department, UserGroup)
	UserTypeCode = dict.Dict[string, int]{User: 0, Department: 1, UserGroup: 2}
)

const (
	UserRole_NormalUser = 1 << iota
	UserRole_AppAdmin
	UserRole_SystemAdmin
	UserRole_SuperAdmin = UserRole_AppAdmin | UserRole_SystemAdmin
)

const (
	LogLevel_INFO = iota
	LogLevel_WARN
)

const (
	OpType_CREATE = iota
	OpType_ADD
	OpType_SET
	OpType_DELETE
	OpType_EDIT
)
