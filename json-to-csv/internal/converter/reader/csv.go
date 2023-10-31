package reader

import (
	"io"
	"learn/internal/converter/models"
)

type CsvReader struct {
	IO   io.Reader
	Logs []models.LogRow
}

// ReadData todo read data
func (c *CsvReader) ReadData() {
}

func (c *CsvReader) GetLogs() *[]models.LogRow {
	return &c.Logs
}
