package ch1

import (
	"strings"
)


func UsingLoop(args []string) {
	var res string
	var sep = " "
	for i := 0; i < len(args); i++ {
		res += res + sep + args[i]
	}
	println(res)
}

func UsingJoin(args []string) {
	res := strings.Join(args, " ")
	println(res)
}
