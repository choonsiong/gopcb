package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	flag.Parse()
	flags := flag.Args()

	if len(flags) == 0 {
		log.Fatal("error no input file")
	}

	// Range over all the input files
	for _, f := range flags {
		_, err := os.Stat(f) // Is the file exists?
		if err != nil {
			log.Println(err)
			continue
		}

		mtd, err := parse(f)
		if err != nil {
			log.Println(err)
			continue
		}

		err = mtd.out()
		if err != nil {
			log.Println(err)
			continue
		}
	}
}
