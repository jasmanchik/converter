package writer

import (
	"encoding/json"
	"io"
	"learn/internal/converter/models"
	"strings"
)

type JsonWriter struct {
	IO io.Writer
}

func (c *JsonWriter) WriteHead(logs []models.LogRow) error {
	return nil
}

func (c *JsonWriter) WriteRow(log models.LogRow) error {
	var sb strings.Builder

	b, err := json.Marshal(log)
	if err != nil {
		return err
	}

	sb.Write(b)

	return nil
}
