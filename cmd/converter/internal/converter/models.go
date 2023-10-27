package converter

import (
	"io"
)

type Log struct {
	Id   string         `json:"id"`
	Lat  float64        `json:"latitude"`
	Lon  float64        `json:"longitude"`
	Tags map[string]any `json:"tags"`
}

type Writable interface {
	SaveData()
}

type CsvWriter struct {
	IO     io.Writer
	Reader Readable
}

type YamlWriter struct {
	IO     io.Writer
	Reader Readable
}

type ProtoWriter struct {
	IO     io.Writer
	Reader Readable
}

type JSONWriter struct {
	IO     io.Writer
	Reader Readable
}

type Readable interface {
	ReadData()
	GetLogs() *[]Log
}

type CsvReader struct {
	IO   io.Reader
	Logs []Log
}

type JSONReader struct {
	IO   io.Reader
	Logs []Log
}

type YAMLReader struct {
	IO   io.Reader
	Logs []Log
}

type ProtoReader struct {
	IO   io.Reader
	Logs []Log
}
