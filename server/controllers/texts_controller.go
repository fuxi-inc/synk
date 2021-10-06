package synk

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TextsController(c *gin.Context) {
	var json struct {
		Raw string
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"content": json.Raw})
	}

}
