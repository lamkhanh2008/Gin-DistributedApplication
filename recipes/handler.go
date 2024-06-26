package recipes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

func NewRecipeHandler(c *gin.Context) {
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	recipe.Id = xid.New().String()
	recipe.PublishedAt = time.Now()
	// recipes = append(recipes, recipe)
	c.JSON(http.StatusOK, recipe)
}
