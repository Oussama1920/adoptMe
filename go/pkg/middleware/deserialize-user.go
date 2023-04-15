package middleware

import (
	"context"
	"net/http"
	"strings"

	config "github.com/Oussama1920/adoptMe/go/pkg/config"
	"github.com/Oussama1920/adoptMe/go/pkg/db"
	utilis "github.com/Oussama1920/adoptMe/go/pkg/utilis"
	"github.com/gin-gonic/gin"
)

func DeserializeUser(dbHandler db.DbHandler, ctx context.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token string
		cookie, err := ctx.Cookie("token")

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		} else if err == nil {
			token = cookie
		}

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}
		var tokenConfig utilis.TokenConfig

		if err := config.GetDataConfiguration("service.token", &tokenConfig); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail to parse config", "message": err.Error()})
		}
		sub, err := utilis.ValidateToken(token, tokenConfig.TOKEN_SECRET)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		user, err := dbHandler.GetUserById(ctx, sub.(string))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
			return
		}

		ctx.Set("currentUser", user)
		ctx.Next()
	}
}
