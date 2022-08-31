package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trichardsonjr78/go-gin-api-medium/pkg/books"
	"github.com/trichardsonjr78/go-gin-api-medium/pkg/common/db"
	"github.com/trichardsonjr78/go-gin-api-medium/pkg/terraform_api"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	h := db.Init(dbUrl)

	books.RegisterRoutes(r, h)
	terraform_api.RegisterRoutes(r, h)
	//db.Init(dbUrl)

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"port":  port,
	// 		"dbUrl": dbUrl,
	// 	})

	// })

	r.Run(port)

}
