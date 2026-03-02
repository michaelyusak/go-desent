package middleware

import (
	"michaelyusak/go-desent.git/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Auth struct {
	authService service.Auth
}

func NewAuth(
	authService service.Auth,
) Auth {
	return Auth{
		authService: authService,
	}
}

func (m Auth) AuthGuard() func(c *gin.Context) {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("authorization")
		if authorization == "" {
			logrus.Warn("[middleware][auth][AuthGuard] empty authorization")
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
				"message": "unauthorized",
			})
			return
		}

		token := strings.Split(authorization, " ")
		if len(token) != 2 || strings.ToLower(token[0]) != "bearer" {
			logrus.Warn("[middleware][auth][AuthGuard] invalid authorization")

			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
				"message": "unauthorized",
			})
			return
		}

		valid := m.authService.ValidateToken(c.Request.Context(), token[1])
		if !valid {
			logrus.WithField("token", token[1]).Warn("[middleware][auth][AuthGuard] token not registered")

			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
				"message": "unauthorized",
			})
			return
		}

		c.Next()
	}
}
