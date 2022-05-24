package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	flag.Parse()
	flags := flag.Args()

	if len(flags) == 0 {
		log.Fatal("no input file")
	}

	for _, f := range flags {
		mtd, err := parse(f)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(mtd)
	}
}
