package controller

import (
	domain "api/internal/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CollectController struct {
	csvService domain.ICsvService
}

func NewCollectController(csvService domain.ICsvService) *CollectController {
	return &CollectController{
		csvService: csvService,
	}
}

// CollectData handles the POST /upload endpoint for CSV files.
func (cc *CollectController) CollectData(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
		return
	}

	// Check if the file is too large
	const maxFileSize = 20 * 1024 * 1024 // 20 MB
	if file.Size > maxFileSize {
		c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": "File is too large"})
		return
	}

	// Open the file
	openedFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not open file"})
		return
	}
	defer openedFile.Close()

	// This is a placeholder for where you'd check the token's validity and permissions
	// You'll need to replace this with your actual auth logic
	// if !isValidToken(c.GetHeader("Authorization")) {
	// 	c.JSON(http.StatusForbidden, gin.H{"error": "Invalid or missing token"})
	// 	return
	// }

	if err := cc.csvService.ProcessCsv(openedFile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process CSV"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "CSV processed successfully!"})
}
