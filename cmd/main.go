/*
MIT License

Copyright (c) 2022 Lee Choon Siong

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
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
