package middleware

import (
	httpAdapter "main/drivenadapters/http"
	apiErr "main/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TokenParser(auth httpAdapter.AuthorizationClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			resp := apiErr.GeneralBizErr(http.StatusUnauthorized, "token is missing")
			c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			return
		}
		info, err := auth.ParseToken(token)
		if err != nil {
			resp := apiErr.GeneralBizErr(http.StatusUnauthorized, err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			return
		}
		c.Set("user_id", info.UserID)
		c.Set("user_name", info.UserName)
		c.Set("user_roles", info.UserRoles)
	}
}
