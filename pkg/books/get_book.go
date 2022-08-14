package AddBookRequestBody

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trichardsonjr78/go-gin-api-medium/pkg/common/models"
)

func (h handler) GetBooks(c *gin.Context) {
	//var books []models.Book
	id := c.Param("id")

	var book models.Book

	// if result := h.DB.Find(&books); result.Error != nil {
	// 	c.AbortWithError(http.StatusNotFound, result.Error)
	// 	return
	// }

	if result := h.DB.First(&book, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &books)
}
