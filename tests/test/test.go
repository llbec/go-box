package main

import (
	"fmt"

	"github.com/pazzabec/go-box/command"
)

func main() {
	var ens command.Entrys
	ens.Parse()
	for i := 0; i < ens.Len(); i++ {
		fmt.Print("\t", ens.Args(i))
	}
	fmt.Println("\n\tFlag t = ", ens.Flags("t"))
	fmt.Println("\tFlag h = ", ens.Flags("h"))
}
