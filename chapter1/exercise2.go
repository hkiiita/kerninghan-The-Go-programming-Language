package ch1

import (
	"fmt"
	"os"
)

func BaseEnv() {
	//Using indexes
	println("Using indexes")
	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("index : %d -- value %s\n", i, os.Args[i])
	}

	//Using range
	println("Using range")
	for i, value := range os.Args {
		fmt.Printf("index : %d -- value %s\n", i, value)
	}

}
