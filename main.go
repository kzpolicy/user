package main

import (
	"github.com/gin-contrib/cors"
	// "github.com/kzpolicy/user/controller"
	"github.com/kzpolicy/user/middleware"
	"local.packages/controller"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/volatiletech/sqlboiler/boil"
)

//go:generate sqlboiler --wipe mysql

func main() {
	r := gin.Default()

	// ミドルウェア
	r.Use(middleware.RecordUaAndTime)

	// CORS 対応
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	r.Use(cors.New(config))

	// ルーティング
	TodoRoute := r.Group("/api/v1")
	{
		v1 := TodoRoute.Group("/user")
		{
			v1.POST("/create", controller.CreateUser)
		}
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	r.Run()
}
