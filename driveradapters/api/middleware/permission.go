package middleware

import (
	"main/errors"
	"net/http"

	"github.com/biocrosscoder/flex/typed/collections/set"
	"github.com/gin-gonic/gin"
)

func PermissionValidator(allow set.Set[int]) gin.HandlerFunc {
	return func(c *gin.Context) {
		roles, exist := c.Get("user_roles")
		if !exist {
			resp := errors.GeneralBizErr(http.StatusForbidden, "user_roles not exist")
			c.AbortWithStatusJSON(http.StatusForbidden, resp)
			return
		}
		if set.Of(roles.([]int)...).IsDisjoint(allow) {
			resp := errors.GeneralBizErr(http.StatusForbidden, "user has no permission")
			c.AbortWithStatusJSON(http.StatusForbidden, resp)
		}
	}
}
