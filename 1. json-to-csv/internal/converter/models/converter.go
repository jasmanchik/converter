package models

import (
	"learn/internal/converter"
	"log"
)

type Converter struct {
}

func (c *Converter) Convert(r converter.Readable, w converter.Writable) {
	logs := r.GetLogs()
	err := w.WriteHead(*logs)
	if err != nil {
		log.Fatalln("can't write heads: ", err)
		return
	}
	for _, row := range *logs {
		err := w.WriteRow(row)
		if err != nil {
			log.Fatalln("can't write row: ", err)
			return
		}
	}
}
