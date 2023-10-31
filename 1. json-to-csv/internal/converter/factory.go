package converter

import (
	"errors"
	"io"
	"learn/internal/converter/reader"
	"learn/internal/converter/writer"
)

var ErrUnexpectedExt = errors.New("unexpected file's extension")

func GetReader(r io.Reader, ext string) (Readable, error) {
	if ext == "json" {
		return &reader.JsonReader{IO: r}, nil
	} else if ext == "protobuf" {
		return &reader.ProtoReader{IO: r}, nil
	} else if ext == "csv" {
		return &reader.CsvReader{IO: r}, nil
	} else if ext == "yaml" {
		return &reader.YamlReader{IO: r}, nil
	}

	return nil, ErrUnexpectedExt
}

func GetWriter(w io.Writer, ext string) (Writable, error) {

	if ext == "json" {
		return &writer.JsonWriter{IO: w}, nil
	} else if ext == "protobuf" {
		return &writer.ProtoWriter{IO: w}, nil
	} else if ext == "csv" {
		return &writer.CsvWriter{IO: w}, nil
	} else if ext == "yaml" {
		return &writer.YamlWriter{IO: w}, nil
	}

	return nil, ErrUnexpectedExt
}
