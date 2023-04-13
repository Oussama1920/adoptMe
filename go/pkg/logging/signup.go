package logging

import (
	"context"
	"net/http"

	db "github.com/Oussama1920/adoptMe/go/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SignUp(dbHandler db.DbHandler, ctx context.Context, logger *logrus.Logger) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var newUser db.User
		// Call BindJSON to bind the received JSON to
		if err := c.BindJSON(&newUser); err != nil {
			logger.Error(err)
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failed to Parse User"})
			return
		}
		if err := dbHandler.AddUser(ctx, newUser); err != nil {
			logger.Error(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to insert user", "error": err.Error()})
			return
		}
		//let's save the user
		c.IndentedJSON(http.StatusOK, gin.H{"message": "New User added successfully"})
	}

	return gin.HandlerFunc(fn)
}
