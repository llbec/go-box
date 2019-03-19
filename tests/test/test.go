package main

import (
	"fmt"
	"reflect"

	"github.com/pazzabec/go-box/order"
)

func main() {
	var ens order.Entrys
	ens.Parse()
	for i := 0; i < ens.Len(); i++ {
		fmt.Print("\t", ens.Args(i))
	}
	fmt.Println("\n\tFlag t = ", ens.Flags("t"))
	fmt.Println("\tFlag h = ", ens.Flags("h"))
	fmt.Println("type:", reflect.TypeOf(ens))
}
