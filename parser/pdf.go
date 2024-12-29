import (
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func parsePDF(filePath string) (string, error) {
	// Extract text from the PDF file
	text, err := api.ExtractTextFile(filePath, nil, nil)
	if err != nil {
		return "", fmt.Errorf("failed to extract text: %w", err)
	}
	return text, nil
}
