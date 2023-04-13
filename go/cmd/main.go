package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	config "github.com/Oussama1920/adoptMe/go/pkg/config"
	db "github.com/Oussama1920/adoptMe/go/pkg/db"
	logging "github.com/Oussama1920/adoptMe/go/pkg/logging"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	filename := filepath.Base(os.Args[0])
	configFilename := flag.String("config", "config/config.yaml", "config filename")
	configPaths := []string{*configFilename, "/etc/axis/" + filename, "/"}

	config.Init(*configFilename, configPaths)

	router := gin.Default()
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
		logger.Fatal("failed to initialise db worker : %#v --> exit", err)
		os.Exit(1)
	}

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/signup", logging.SignUp(dbWorker, ctx, logger))
		v1.GET("/login", func(c *gin.Context) {
			//let's save the user
			c.IndentedJSON(http.StatusOK, gin.H{"message": "New User added successfully"})
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
