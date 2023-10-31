package main

import (
	"flag"
	"fmt"
	"learn/internal/converter"
	"learn/internal/converter/models"
	"log"
	"os"
)

const defInExt = "json"
const defOutExt = "csv"

func main() {
	//var _ FooInterface = (*MyFoo)(nil)
	var config, output, cExt, oExt = "", "", "", ""

	flag.StringVar(&config, "config", "", "path of logger file")
	flag.StringVar(&output, "output", "", "path of new file")
	flag.StringVar(&cExt, "cExt", "", "config's' extension")
	flag.StringVar(&oExt, "oExt", "", "output's extension")
	flag.Parse()

	if cExt == "" {
		cExt = defInExt
	}
	if oExt == "" {
		oExt = defOutExt
	}
	if cExt == oExt {
		log.Fatalf("can't convert %s to %s", cExt, oExt)
		return
	}
	rf := os.Stdin
	if config != "" {
		rf, err := os.Open(config)
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				log.Println("can't close file:", err)
			}
		}(rf)

		if err != nil {
			log.Fatalln("can't open file:", err)
		}
	}
	r, err := converter.GetReader(rf, cExt)
	if err != nil {
		fmt.Errorf("getting reader: %v", err)
	}

	wf := os.Stdout
	if output != "" {
		wf, err := os.Create(output)
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				log.Println("can't close file:", err)
			}
		}(wf)
		if err != nil {
			log.Fatalln("can't create the file:", err)
		}
	}
	w, err := converter.GetWriter(wf, oExt)
	if err != nil {
		fmt.Errorf("getting writer: %v", err)
	}

	c := models.Converter{}
	c.Convert(r, w)
}
