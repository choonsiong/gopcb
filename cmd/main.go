package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
)

var wg = new(sync.WaitGroup)

func main() {
	bufferSize := flag.Int("bufferSize", 100, "Buffer size")
	useBuffer := flag.Bool("useBuffer", false, "Use buffer write")

	flag.Usage = func() {
		usage := `usage: gopcb [-flag] ... file1 file2 ...`
		fmt.Println(usage)
		fmt.Println()
		flag.PrintDefaults()
		fmt.Println()
		example := `Examples:
$ gopcb data/file1.json data/file2.json`
		fmt.Println(example)
	}

	flag.Parse()
	flags := flag.Args()

	if len(flags) == 0 {
		flag.Usage()
	}

	count := len(flags)
	wg.Add(count)

	// Range over all the input files
	for _, f := range flags {
		_, err := os.Stat(f) // Is the file exists?
		if err != nil {
			log.Println(err)
			wg.Done()
			continue
		}

		mtd, err := parse(f)
		if err != nil {
			log.Println(err)
			wg.Done()
			continue
		}

		if *useBuffer {
			go mtd.bufferOut(*bufferSize)
		} else {
			go mtd.out()

		}
	}

	wg.Wait()
}
