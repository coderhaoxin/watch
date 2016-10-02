package main

import "fmt"
import "os"

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
