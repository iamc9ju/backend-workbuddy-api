package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) *gin.Engine {
	if os.Getenv("STATE") != "DEV" {
		// r.Use(handle.CloudWatchMiddleware(awsconfig.LogGroupName, awsconfig.LogStreamName))
	}

	// database

	// db, err := postgres.NewPostgresDatagase()
	// if err != nil {
	// 	log.Fatalf("Cannot connect to database. '%v'", err)
	// }

	//repositories

	//service

	//controller

	return r
}
