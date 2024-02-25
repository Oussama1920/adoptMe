package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	utilis "github.com/Oussama1920/adoptMe/go/pkg/utilis"

	config "github.com/Oussama1920/adoptMe/go/pkg/config"
	db "github.com/Oussama1920/adoptMe/go/pkg/db"
	logging "github.com/Oussama1920/adoptMe/go/pkg/logging"
	pet "github.com/Oussama1920/adoptMe/go/pkg/pet"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	filename := filepath.Base(os.Args[0])
	configFilename := flag.String("config", "config/config.yaml", "config filename")
	configPaths := []string{*configFilename, "/etc/axis/" + filename, "/"}

	config.Init(*configFilename, configPaths)

	router := gin.Default()
	router.Use(corsMiddleware())

	logger := logrus.New()
	ctx, _ := context.WithCancel(context.Background())

	file, err := os.Create("logFile.txt")
	if err == nil {
		mw := io.MultiWriter(os.Stdout, file)
		logger.SetOutput(mw)
	} else {
		logger.Errorf("Failed to open log to file '%s', using default stderr", "logFile.txt")
	}
	logger.Infof("Set log file to '%s'", file.Name())
	dbWorker, err := initDB(ctx, logger)
	if err != nil {
		logger.Fatalf("failed to initialise db worker : %#v --> exit", err)
		os.Exit(1)
	}

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/signup", logging.SignUp(dbWorker, ctx, logger))
		v1.POST("/login", logging.Login(dbWorker, ctx, logger))
		v1.GET("/verifyemail/:verificationCode", logging.VerifyEmail(dbWorker, ctx, logger))

	}
	auth := router.Group("/auth")
	{
		users := auth.Group("/users")
		{
			auth.GET("/logout", logging.LogOut(dbWorker, ctx, logger))
			//	users.GET("/me", logging.GetUser(dbWorker, ctx, logger))
			users.GET("/me", corsAuthentication(dbWorker, ctx, logger), logging.GetMe)
			users.PUT("/me", corsAuthentication(dbWorker, ctx, logger), func(c *gin.Context) {
				logging.UpdateUser(c, dbWorker, logger)
			})
			//users.PUT("/me", middleware.DeserializeUser(dbWorker, ctx), logging.GetMe)
		}
	}
	pets := v1.Group("/pets")
	{
		pets.POST("/pet", corsAuthentication(dbWorker, ctx, logger), func(c *gin.Context) {
			logging.AddPet(c, dbWorker, logger)
		})
		pets.GET("/pet/:id", func(c *gin.Context) {
			logging.GetPet(c, dbWorker, logger)
		})
		pets.POST("/pet/search", func(c *gin.Context) {
			pet.SearchPet(c, dbWorker, logger)
		})

	}

	router.Run(":8080")
	logger.Info("serving on 8080 ...")
}

func initDB(ctx context.Context, appLog *logrus.Logger) (db.DbHandler, error) {
	var dbc db.DBConfig

	if err := config.GetDataConfiguration("service.database", &dbc); err != nil {
		return nil, fmt.Errorf("can't read DB configuration : %v", err.Error())
	}
	return db.NewDB(ctx, dbc, appLog)

}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight OPTIONS requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		// Continue to the next middleware or route handler
		c.Next()

	}
}

func corsAuthentication(dbHandler db.DbHandler, ctx context.Context, logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Middleware to verify authentication token
		// Get the authentication token from the request headers
		var token string
		cookie, err := c.Cookie("token")

		authorizationHeader := c.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		} else if err == nil {
			token = cookie
		}

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}
		var tokenConfig utilis.TokenConfig

		if err := config.GetDataConfiguration("service.token", &tokenConfig); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail to parse config", "message": err.Error()})
		}
		sub, err := utilis.ValidateToken(token, tokenConfig.TOKEN_SECRET)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		user, err := dbHandler.GetUserById(ctx, sub.(string))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
			return
		}
		logger.Infof("found user is %v", user)
		c.Set("currentUser", user)

		// Continue to the next middleware or route handler
		c.Next()

	}
}
