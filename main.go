package main

import (
	"github.com/czhi-bin/onecv-assessment/config"
	"github.com/czhi-bin/onecv-assessment/db"
	"github.com/czhi-bin/onecv-assessment/router"
	"github.com/czhi-bin/onecv-assessment/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.InitLogger()
	db.Init(false)

	r := gin.Default()
	
	router.RegisterRoutes(r)

	utils.Logger.Info("Starting server")
	r.Run(config.LOCAL_HOST)
}
