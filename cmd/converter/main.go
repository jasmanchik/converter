package main

import (
	"flag"
	"fmt"
	"io"
	"learn/cmd/converter/internal/converter"
	"log"
	"os"
)

const defInExt = "json"
const defOutExt = "csv"

func main() {
	var config, output, cExt, oExt = "", "", "", ""

	flag.StringVar(&config, "config", "", "path of logger file")
	flag.StringVar(&output, "output", "", "path of new file")
	flag.StringVar(&cExt, "cExt", "", "config's' extension")
	flag.StringVar(&oExt, "oExt", "", "output's extension")
	flag.Parse()

	if cExt == "" {
		cExt = defInExt
	}
	rf := getIOReader(config)
	var r converter.Readable
	if cExt == "json" {
		r = &converter.JSONReader{IO: rf}
	} else if cExt == "protobuf" {
		r = &converter.ProtoReader{IO: rf}
	} else if cExt == "csv" {
		r = &converter.CsvReader{IO: rf}
	} else if cExt == "yaml" {
		r = &converter.YAMLReader{IO: rf}
	}
	r.ReadData()

	if oExt == "" {
		oExt = defOutExt
	}
	wf := getIOWriter(output)

	var w converter.Writable
	if oExt == "json" {
		w = &converter.JSONWriter{IO: wf, Reader: r}
	} else if oExt == "protobuf" {
		w = &converter.ProtoWriter{IO: wf, Reader: r}
	} else if oExt == "csv" {
		fmt.Println(321)
		w = &converter.CsvWriter{IO: wf, Reader: r}
	} else if oExt == "yaml" {
		w = &converter.YamlWriter{IO: wf, Reader: r}
	}
	w.SaveData()
}

func getIOReader(config string) io.Reader {
	if config != "" {
		f, err := os.Open(config)
		/*defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				log.Println("can't close file:", err)
			}
		}(f)*/

		if err != nil {
			log.Fatalln("can't open file:", err)
		}
		return f
	} else {
		return os.Stdin
	}
}

func getIOWriter(output string) io.Writer {
	if output != "" {
		f, err := os.Create(output)
		/*defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				log.Println("can't close file:", err)

			}
		}(f)*/

		if err != nil {
			log.Fatalln("can't create the file:", err)
		}
		return f
	} else {
		return os.Stdout
	}
}
