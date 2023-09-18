package main

import (
	"flag"
	"fmt"
)

func main() {
	pathPtr := flag.String("path", "", "path of binary to simulate")
	flag.Parse()
	fmt.Println("path flag: ", *pathPtr)
}
