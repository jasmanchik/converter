package converter

import (
	"fmt"
	"log"
	"strings"
)

const staticCSVHeadersNum = 3

func (c *CsvWriter) SaveData() {
	var h = make(map[string]int)
	var hStr = "id,latitude,longitude"

	for _, l := range *c.Reader.GetLogs() {
		for tagName := range l.Tags {
			_, ok := h[tagName]
			if !ok {
				h[tagName] = staticCSVHeadersNum + len(h)
				hStr += "," + tagName
			}
		}
	}

	var sb strings.Builder
	sb.WriteString(hStr + "\n")

	for _, l := range *c.Reader.GetLogs() {
		lat := fmt.Sprintf("%g", l.Lat)
		lon := fmt.Sprintf("%g", l.Lon)

		row := make([]string, staticCSVHeadersNum+len(h))
		row[0] = l.Id
		row[1] = lat
		row[2] = lon

		for tagName, val := range l.Tags {
			switch val.(type) {
			case int:
				row[h[tagName]] = fmt.Sprintf("%d", val)
			case float32:
				row[h[tagName]] = fmt.Sprintf("%g", val)
			case float64:
				row[h[tagName]] = fmt.Sprintf("%g", val)
			case string:
				row[h[tagName]] = fmt.Sprintf("%s", val)
			case bool:
				row[h[tagName]] = fmt.Sprintf("%t", val)
			}
		}

		sb.WriteString(strings.Join(row, ",") + "\n")
	}

	_, err := c.IO.Write([]byte(sb.String()))
	if err != nil {
		log.Fatalln("can't write in the file:", err)
	}
}

// SaveData todo запись данных
func (c *YamlWriter) SaveData() {
}

// SaveData todo запись данных
func (c *ProtoWriter) SaveData() {
}

// SaveData todo запись данных
func (c *JSONWriter) SaveData() {
}
