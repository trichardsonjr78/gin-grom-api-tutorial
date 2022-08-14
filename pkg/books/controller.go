package books

import (
	"gorm.io/gorm"
)

type handler strcut {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/books")
	routes.POST("/", h.AddBook)
	routes.GET("/", h.GetBooks)
	routes.GET("/:id", h.GetBook)
	routes.PUT("/:id", h.UpdateBook)
	routes.DELETE("/id", h.DeleteBook)
}