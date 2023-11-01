package writer

import (
	"io"
	"learn/internal/converter/models"
)

type YamlWriter struct {
	IO io.Writer
}

func (c *YamlWriter) WriteHead(logs []models.LogRow) error {
	return nil
}

func (c *YamlWriter) WriteRow(log models.LogRow) error {
	return nil
}
