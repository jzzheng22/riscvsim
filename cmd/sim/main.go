package main

import (
	"flag"
	"fmt"

	"github.com/jzzheng22/riscvsim/pkg/simulator"
)

func main() {
	pathPtr := flag.String("path", "", "path of binary to simulate")
	flag.Parse()
	fmt.Println("path flag: ", *pathPtr)

	// TODO: Not sure if this is the best thing to return from Simulate()
	exitCode, err := simulator.Simulate(*pathPtr)
	if err != nil {
		fmt.Println("Error in execution: ", err)
	}
	fmt.Println("Simulation exit with code ", exitCode)
}
