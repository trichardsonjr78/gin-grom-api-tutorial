package terraform_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trichardsonjr78/go-gin-api-medium/pkg/common/models"
)

type AddTFOResourceBody struct {
	UUID      string `json:"uuid" gorm:"primaryKey"`
	CreatedBy string `json:"createdby"`
	CreatedAt string `json:"createdat"`
	UpdatedBy string `json:"updatedby"`
	UpdatedAt string `json:"updatedat"`
	DeletedBy string `json:"deletedby"`
	DeletedAt string `json:"deleetedat"`
}

func (h handler) Addtfo(c *gin.Context) {
	body := AddTFOResourceBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var tfo_resource models.TFOResource

	tfo_resource.UUID = body.UUID
	tfo_resource.CreatedAt = body.CreatedAt
	tfo_resource.UpdatedBy = body.UpdatedBy
	tfo_resource.UpdatedAt = body.UpdatedAt
	tfo_resource.DeletedBy = body.DeletedBy
	tfo_resource.DeletedAt = body.DeletedAt

	if result := h.DB.Create(&tfo_resource); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &tfo_resource)
}
