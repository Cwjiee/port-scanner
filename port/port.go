package port

import (
	"net"
	"time"
	"strconv"
)

type ScanResult struct {
	Port int 
	Status string
}

func ScanPort(protocol, hostname string, port int) ScanResult {
	
	result := ScanResult{Port: port}
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

	return results
}


