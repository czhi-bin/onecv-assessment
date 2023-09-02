package main

import (
	"github.com/czhi-bin/onecv-assessment/db"
	"github.com/czhi-bin/onecv-assessment/router"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()

	r := gin.Default()

	router.RegisterRoutes(r)

	r.Run("127.0.0.1:18000")
}
