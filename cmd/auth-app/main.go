package main

import (
	"database/sql"
	"main/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.InitConfig()
	var db *sql.DB = utils.GetDB()
	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":" + utils.HTTP_PORT)
}
