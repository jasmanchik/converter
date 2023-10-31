package reader

import (
	"io"
	"learn/internal/converter/models"
)

type YamlReader struct {
	IO   io.Reader
	Logs []models.LogRow
}

// ReadData todo read data
func (c *YamlReader) ReadData() {
}

func (c *YamlReader) GetLogs() *[]models.LogRow {
	return &c.Logs
}
