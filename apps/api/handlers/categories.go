package handlers

import (
	"aurora/db"
	"aurora/db/models"
	"aurora/handlers/dto"
	"aurora/services/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func CreateCategory(c *gin.Context) {
	body := c.MustGet("body").(dto.CreateCategoryDto)

	if body.ParentId != nil {
		if _, err := uuid.Parse(body.ParentId.String()); err != nil {
			utils.ErrorResponse(c, http.StatusBadRequest, "parentId is malformed")
			return
		}
	}

	category := &models.Category{
		Name:     body.Name,
		ParentId: body.ParentId,
	}

	res := db.Client.Create(&category)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.Status(http.StatusCreated)
}
