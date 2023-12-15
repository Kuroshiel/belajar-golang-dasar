package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	for _, args := range args {
		fmt.Println(args)
	}

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error", err.Error())
	}
}
