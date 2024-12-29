import (
	"github.com/xuri/excelize/v2"
)

func writeToExcel(data [][]string, filePath string) error {
	f := excelize.NewFile()
	for rowIndex, row := range data {
		for colIndex, value := range row {
			cell, _ := excelize.CoordinatesToCellName(colIndex+1, rowIndex+1)
			f.SetCellValue("Sheet1", cell, value)
		}
	}
	return f.SaveAs(filePath)
}
