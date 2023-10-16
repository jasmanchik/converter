package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func getArg(name string) (string, bool) {
	for _, v := range os.Args[1:] {
		argName := strings.Split(v, "=")
		if argName[0] == name {
			return argName[1], true
		}
	}

	return "", false
}

type Log struct {
	Id   string         `json:"id"`
	Lat  float64        `json:"latitude"`
	Lon  float64        `json:"longitude"`
	Tags map[string]any `json:"tags"`
}

func (c *Log) ToSlice() []string {
	var fields = make([]string, 0, 3+len(c.Tags))
	fields = append(fields, c.Id, fmt.Sprintf("%f", c.Lat), fmt.Sprintf("%f", c.Lon))

	for _, val := range c.Tags {
		tmp := ""
		switch val.(interface{}).(type) {
		case int:
			tmp = fmt.Sprintf("%d", val)
		case float32:
			tmp = fmt.Sprintf("%g", val)
		case float64:
			tmp = fmt.Sprintf("%g", val)
		case string:
			tmp = fmt.Sprintf("%s", val)
		case bool:
			tmp = fmt.Sprintf("%t", val)
		}

		fields = append(fields, tmp)
	}

	return fields
}

func ScanConsole(logCh *chan Log, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(*logCh)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Bytes()
		if string(text) == "exit" {
			return
		}

		logSt := Log{}
		err := json.Unmarshal(text, &logSt)
		if err != nil {
			log.Println("не смог сделать анмаршал json строки:", err)
			continue
		}
		*logCh <- logSt
	}
}

func ScanFile(logCh *chan Log, wg *sync.WaitGroup, filePath string) {
	defer wg.Done()
	defer close(*logCh)

	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		log.Fatalln("не удалось открыть файл:", err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		logSt := Log{}
		err = json.Unmarshal(scanner.Bytes(), &logSt)
		if err != nil {
			log.Println("ошибка при декодировании json строки:", err)
			continue
		}
		*logCh <- logSt
	}
}

func WriteToCsv(logsCh chan Log, wg *sync.WaitGroup, csvPath string) {
	defer wg.Done()
	f, err := os.Create(csvPath)
	defer f.Close()

	if err != nil {
		log.Fatalln("не удалось создать файл:", err)
	}

	w := csv.NewWriter(f)
	defer w.Flush()
	w.Comma = ';'

	for {
		select {
		case logSt, ok := <-logsCh:
			if ok {
				if err := w.Write(logSt.ToSlice()); err != nil {
					log.Fatalln("не удалось записать в файл:", err)
				}
			} else {
				return
			}
		}
	}
}

func WriteToConsole(ch chan Log, wg *sync.WaitGroup) {
	defer wg.Done()
	out := bufio.NewWriter(os.Stdout)
loop:
	for {
		select {
		case logSt, ok := <-ch:
			if !ok {
				break loop
			}
			_, err := fmt.Fprintln(out, logSt.ToSlice())
			if err != nil {
				log.Println("не смог вывести результат в консоль:")
			}
			out.Flush()
		}
	}
}

func main() {
	logsCh := make(chan Log)
	wg := sync.WaitGroup{}
	wg.Add(2)

	jsonPath, ok := getArg("config")
	if ok {
		if _, err := os.Stat(jsonPath); err != nil {
			log.Fatalln("не удалось прочитать файл:", err)
		}
		go ScanFile(&logsCh, &wg, jsonPath)
	} else {
		go ScanConsole(&logsCh, &wg)
	}

	csvPath, ok := getArg("output")
	if ok {
		go WriteToCsv(logsCh, &wg, csvPath)
	} else {
		go WriteToConsole(logsCh, &wg)
	}

	wg.Wait()
}
