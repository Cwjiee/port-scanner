package main

import (
	"fmt"
	"github.com/Cwjiee/port-scanner/port"
)

func main() {

	fmt.Println("port scanning")
	// port.ScanIcmpPort()
	hostname := "localhost"

	tcpResults := make(chan []port.ScanResult)
	udpResults := make(chan []port.ScanResult)
	wideTcpResults := make(chan []port.ScanResult)
	wideUdpResults := make(chan []port.ScanResult)

	go func() {
		defer close(tcpResults)
		tcpResults <- port.ScanAllPorts("tcp", hostname)
	}()

	go func() {
		defer close(udpResults)
		udpResults <- port.ScanAllPorts("udp", hostname)
	}()

	tcpScanResults := <-tcpResults
	udpScanResults := <-udpResults

	results := append(tcpScanResults, udpScanResults...)

	fmt.Println("scan results")
	for _, result := range results {
		fmt.Printf("%+v\n", result)
	}

	go func() {
		defer close(wideTcpResults)
		wideTcpResults <- port.WideScan("tcp", hostname)
	}()

	go func() {
		defer close(wideUdpResults)
		wideUdpResults <- port.WideScan("udp", hostname)
	}()

	wideTcpScanResults := <- wideUdpResults
	wideUdpScanResults := <- wideUdpResults

	wideResults := append(wideTcpScanResults, wideUdpScanResults...)
	
	fmt.Println("\nwide scan results")
	for _, result := range wideResults {
		fmt.Printf("%+v\n", result)
	}
}
