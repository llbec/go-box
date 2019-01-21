package command

import (
	"os"
	"strings"
)

//Entrys a command's args and flags for application
type Entrys struct {
	args  []string
	flags map[string]string
}

//Parse get os args, and analysis
func (ens *Entrys) Parse() {
	ens.args = os.Args[1:]
	ens.flags = make(map[string]string)
	for i := 0; i < len(ens.args); {
		if strings.HasPrefix(ens.args[i], "-") == true {
			ens.flags[ens.args[i][1:]] = ens.args[i+1]
			if i+2 == len(ens.args) {
				ens.args = ens.args[:i]
			} else {
				ens.args = append(ens.args[:i], ens.args[i+2:]...)
			}
		} else {
			i++
		}
	}
}

//Len get args's length
func (ens Entrys) Len() int {
	return len(ens.args)
}

//Args (int) get the position of th args
func (ens Entrys) Args(i int) string {
	if i >= len(ens.args) {
		return ""
	}
	return ens.args[i]
}

//Flags (key string) return value of the key
func (ens Entrys) Flags(key string) string {
	if v, exist := ens.flags[key]; exist {
		return v
	}
	return ""
}
