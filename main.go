package main

import (
	"fmt"
	"time"

	"github.com/Cwjiee/port-scanner/port"
	"github.com/Cwjiee/port-scanner/timer"
)

func main() {
	defer timer.TimeTrack(time.Now())
	fmt.Println("port scanning")
	// port.ScanIcmpPort()
	hostname := "localhost"

	// wideTcpResults := make(chan []port.ScanResult)
	// wideUdpResults := make(chan []port.ScanResult)

	// go func() {
	// 	defer close(wideTcpResults)
	// 	wideTcpResults <- port.WideScan("tcp", hostname)
	// }()

	// go func() {
	// 	defer close(wideUdpResults)
	// 	wideUdpResults <- port.WideScan("udp", hostname)
	// }()

	// wideTcpScanResults := <- wideTcpResults
	wideTcpScanResults := port.WideScan("tcp", hostname)
	// wideUdpScanResults := <- wideUdpResults

	// wideResults := append(wideTcpScanResults, wideUdpScanResults...)
	
	fmt.Println("\nwide scan results")
	for _, result := range wideTcpScanResults {
		fmt.Printf("%+v\n", result)
	}
}
