package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mar-coding/fum-cloud-notification-report-2023/app/db"
	"github.com/mar-coding/fum-cloud-notification-report-2023/app/middleware"
)

func main() {
	sqlDB, err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()
	router := initRouter(sqlDB)
	err = router.Run(":1234")
	if err != nil {
		panic(err)
	}
}

func initRouter(sqlDB *sql.DB) *gin.Engine {
	router := gin.Default()
	api := router.Group("/reports")
	{
		secured := api.Group("/mail").Use(middleware.Auth())
		{
			secured.GET("", func(c *gin.Context) {
				middleware.HandleAllMailRequests(c, sqlDB)
			})
			secured.GET("/:requestId", func(c *gin.Context) {
				middleware.HandleMailRequestByID(c, sqlDB)
			})
			secured.GET("/configs/:configId", func(c *gin.Context) {
				middleware.HandleMailRequestByConfigID(c, sqlDB)
			})
		}
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
	})

	return router
}
