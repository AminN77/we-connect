// Package csv provides the concurrent parsing the csv files. ParseFileConcurrent gets a file and a custom
// marshaller, creates a goroutine pool with goroutines as many as your Cpu cores. These goroutines then
// consume the src channel for incoming csv line and unmarshal with provided unmarshal impl.
// TODO: R&D on reading csv file concurrently via file.Seek
package csv

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"runtime"
	"sync"
)

type Csv struct {
}

func New() *Csv {
	return &Csv{}
}

func (c *Csv) ParseFileConcurrent(file *os.File, marshaller Marshaller) {
	var wg sync.WaitGroup
	reader := csv.NewReader(file)
	src := make(chan []string)

	// extract headers
	_, err := reader.Read()
	if err == io.EOF {
		panic("EOF")
	} else if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func(src chan []string, i int, wg *sync.WaitGroup) {
			defer wg.Done()
			for {
				select {
				case d, ok := <-src:
					if !ok {
						return
					}

					if err := marshaller.Unmarshal(d); err != nil {
						log.Println(err)
					}

				}
			}
		}(src, i, &wg)
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		src <- record
	}

	close(src)
	wg.Wait()
}
