package main

import (
	"flag"
	"fmt"
	"github.com/jhaoheng/printhello"
	"os"
)

func main() {
	name := flag.String("name", "", "")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\n\n")
		fmt.Fprintf(os.Stderr, "Usage : go run cmd.go [command]\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if len(*name) > 0 {
		printhello.Printhello()
		fmt.Println(*name)
	}
}
