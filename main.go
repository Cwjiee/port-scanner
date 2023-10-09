package main

import (
	"fmt"
	"github.com/Cwjiee/port-scanner/port"
)

func main() {

	fmt.Println("port scanning")
	results := port.ScanAllPorts("localhost")
	// port.ScanIcmpPort()
	fmt.Println(results)

	widescanresults := port.WideScan("localhost")
	fmt.Println(widescanresults)
}
