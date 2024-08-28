package ch1

import (
	"fmt"
	"os"
)

func BaseEnv() {
	//Using indexes
	var res string
	var sep = " "
	for i := 0; i < len(os.Args); i++ {
		res += res + sep + os.Args[i]
	}
	fmt.Println(res)

	//Using range
	res = ""
	for _, value := range os.Args {
		res += sep + value
	}
	fmt.Println(res)

	//Direct
	fmt.Println(os.Args)

}
