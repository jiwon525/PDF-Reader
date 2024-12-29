import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		filePath := "./" + file.Filename
		c.SaveUploadedFile(file, filePath)

		text, err := parsePDF(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Process text into structured data
		data := [][]string{
			{"Column1", "Column2"},
			{"Row1Value1", "Row1Value2"},
		}

		excelPath := "./output.xlsx"
		err = writeToExcel(data, excelPath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.File(excelPath)
	})

	r.Run(":8080")
}
