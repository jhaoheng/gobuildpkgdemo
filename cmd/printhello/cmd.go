package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jhaoheng/gobuildpkgdemo/pkg"
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
		pkg.Printhello()
		fmt.Println(*name)
	}
}
