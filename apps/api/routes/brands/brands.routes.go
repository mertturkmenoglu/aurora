package brands

import (
	"aurora/services/db"
	"aurora/services/db/models"
	"aurora/services/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateBrand(c *gin.Context) {
	body := c.MustGet("body").(CreateBrandDto)
	brand := models.Brand{
		Name:        body.Name,
		Description: body.Description,
	}

	res := db.Client.Create(&brand)

	if res.Error != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, res.Error.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": brand,
	})
}

func GetBrandById(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is required")
		return
	}

	var brand models.Brand

	res := db.Client.First(&brand, "id = ?", id)

	if res.Error != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, res.Error.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": brand,
	})
}
