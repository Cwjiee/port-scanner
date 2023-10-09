package port

import (
	"net"
	"time"
	"strconv"
	/* "os"
	"fmt"
	"github.com/tatsushid/go-fastping" */
)

type ScanResult struct {
	Port   string
	Status string
}

/*
func ScanIcmpPort() {
	
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}	

	p.AddIPAddr(ra)
	
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
	}

	p.OnIdle = func() {
		fmt.Println("finish")
	}

	err = p.Run()
	if err != nil {
		fmt.Println(err)
	}
}
*/

func ScanPort(protocol, hostname string, port int) ScanResult {
	
	result := ScanResult{Port: strconv.Itoa(port) + string("/") + protocol}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		result.Status = "Closed"
		return result
	}

	conn.Close()
	result.Status = "Open"
	return result
}

func ScanAllPorts(hostname string) []ScanResult {
	
	var results []ScanResult

	for i := 0; i <= 1024; i++ {
		results = append(results, ScanPort("tcp", hostname, i))
	}

	for i := 0; i <= 1024; i++ {
		results = append(results, ScanPort("udp", hostname, i))
	}

	return results
}

func WideScan(hostname string) []ScanResult {

	var results []ScanResult

	for i := 0; i <= 49152; i++ {
		results = append(results, ScanPort("tcp", hostname, i))
	}

	for i := 0; i <= 49152; i++ {
		results = append(results, ScanPort("udp", hostname, i))
	}

	return results
}
