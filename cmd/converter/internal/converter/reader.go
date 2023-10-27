package converter

import (
	"bufio"
	"encoding/json"
	"log"
)

const maxCap = 1_000_000
const exitCmd = "exit"

// ReadData todo read data
func (c *CsvReader) ReadData() {
}

func (c *CsvReader) GetLogs() *[]Log {
	return &c.Logs
}

// ReadData todo read data
func (c *YAMLReader) ReadData() {
}

func (c *YAMLReader) GetLogs() *[]Log {
	return &c.Logs
}

func (c *JSONReader) ReadData() {
	c.Logs = make([]Log, 0, maxCap)
	scanner := bufio.NewScanner(c.IO)
	for scanner.Scan() {
		text := scanner.Bytes()
		if string(text) == exitCmd {
			break
		}
		logSt := Log{}
		err := json.Unmarshal(text, &logSt)
		if err != nil {
			log.Println("decode error json string:", err)
			continue
		}
		c.Logs = append(c.Logs, logSt)
	}
}

func (c *JSONReader) GetLogs() *[]Log {
	return &c.Logs
}

// ReadData todo read data
func (c *ProtoReader) ReadData() {
}

func (c *ProtoReader) GetLogs() *[]Log {
	return &c.Logs
}
