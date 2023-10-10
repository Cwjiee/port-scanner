package port

import (
	"net"
	"strconv"
	"sync"
	"time"
	/* "os"
	"fmt"
	"github.com/tatsushid/go-fastping" */)

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

func ScanPort(protocol, hostname string, port int, wg *sync.WaitGroup, results chan<- ScanResult) {
	defer wg.Done()

	result := ScanResult{Port: strconv.Itoa(port) + string("/") + protocol}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		result.Status = "Closed"
	} else {
		conn.Close()
		result.Status = "Open"
	}

	results <- result
}

func WideScan(protocol string, hostname string) []ScanResult {
	var wg sync.WaitGroup
	results := make(chan ScanResult)

	for i := 0; i <= 49152; i++ {
		wg.Add(1)
		go ScanPort(protocol, hostname, i, &wg, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var scanResults []ScanResult
	for result := range results {
		scanResults = append(scanResults, result)
	}
	return scanResults
}
