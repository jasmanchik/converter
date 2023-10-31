package writer

import (
	"fmt"
	"io"
	"learn/internal/converter/models"
	"strings"
)

const staticCSVHeadersNum = 3

type CsvWriter struct {
	IO      io.Writer
	headers map[string]int
}

func (c *CsvWriter) WriteHead(logs []models.LogRow) error {
	c.headers = make(map[string]int)
	var hStr = "id,latitude,longitude"

	for _, l := range logs {
		for tagName := range l.Tags {
			_, ok := c.headers[tagName]
			if !ok {
				c.headers[tagName] = staticCSVHeadersNum + len(c.headers)
				hStr += "," + tagName
			}
		}
	}

	var sb strings.Builder
	sb.WriteString(hStr + "\n")
	_, err := c.IO.Write([]byte(sb.String()))
	if err != nil {
		return err
	}
	return nil
}

func (c *CsvWriter) WriteRow(log models.LogRow) error {
	lat := fmt.Sprintf("%g", log.Lat)
	lon := fmt.Sprintf("%g", log.Lon)

	row := make([]string, staticCSVHeadersNum+len(c.headers))
	row[0] = log.Id
	row[1] = lat
	row[2] = lon

	for tagName, val := range log.Tags {
		headerPos := c.headers[tagName]
		switch val.(type) {
		case int:
			row[headerPos] = fmt.Sprintf("%d", val)
		case float32:
			row[headerPos] = fmt.Sprintf("%g", val)
		case float64:
			row[headerPos] = fmt.Sprintf("%g", val)
		case string:
			row[headerPos] = fmt.Sprintf("%s", val)
		case bool:
			row[headerPos] = fmt.Sprintf("%t", val)
		}
	}

	var sb strings.Builder
	sb.WriteString(strings.Join(row, ",") + "\n")

	_, err := c.IO.Write([]byte(sb.String()))
	if err != nil {
		return err
	}

	return nil
}
