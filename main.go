package main

import (
	"fmt"
	"github.com/Cwjiee/port-scanner/port"
)

func main() {

	fmt.Println("port scanning")
	results := port.ScanAllPorts("localhost")
	fmt.Println(results)
}
