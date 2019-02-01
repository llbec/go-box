package main

import (
	"flag"
	"fmt"
)

func main() {
	n := flag.Int("n", 0, "n test")
	flag.Parse()
	fmt.Println(*n, flag.Args(), flag.Parsed())
}
