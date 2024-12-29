package main

import (
	"log"
	"net/http"
	"pdf-to-excel/exporter"
	"pdf-to-excel/parser"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		filePath := "./" + file.Filename
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		text, err := parser.ParsePDF(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		data := [][]string{
			{"Extracted Content"},
			{text},
		}

		excelPath := "./output.xlsx"
		if err := exporter.WriteToExcel(data, excelPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.File(excelPath)
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
