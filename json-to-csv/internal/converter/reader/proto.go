package reader

import (
	"io"
	"learn/internal/converter/models"
)

type ProtoReader struct {
	IO   io.Reader
	Logs []models.LogRow
}

// ReadData todo read data
func (c *ProtoReader) ReadData() {
}

func (c *ProtoReader) GetLogs() *[]models.LogRow {
	return &c.Logs
}
