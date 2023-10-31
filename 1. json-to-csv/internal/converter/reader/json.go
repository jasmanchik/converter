package reader

import (
	"bufio"
	"encoding/json"
	"io"
	"learn/internal/converter/models"
	"log"
)

const maxCap = 1_000_000
const exitCmd = "exit"

type JsonReader struct {
	IO   io.Reader
	Logs []models.LogRow
}

func (c *JsonReader) ReadData() {
	c.Logs = make([]models.LogRow, 0, maxCap)
	scanner := bufio.NewScanner(c.IO)
	for scanner.Scan() {
		text := scanner.Bytes()
		if string(text) == exitCmd {
			break
		}
		logSt := models.LogRow{}
		err := json.Unmarshal(text, &logSt)
		if err != nil {
			log.Println("decode error json string:", err)
			continue
		}
		c.Logs = append(c.Logs, logSt)
	}
}

func (c *JsonReader) GetLogs() *[]models.LogRow {
	return &c.Logs
}
