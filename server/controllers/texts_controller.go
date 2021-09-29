package synk

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
)

func TextsController(c *gin.Context) {
	var json struct {
		Raw string
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Dir(exe)
	filename := uuid.New().String()
	fileErr := qrcode.WriteFile(json.Raw, qrcode.Medium, 256,
		filepath.Join(dir, "uploads", filename+".png"))
	fmt.Println(fileErr, filepath.Join(dir, "uploads", filename+".png"))
	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})

}
