package luhn

import "fmt"

type Endpoint struct {
	Protocol string
	Host     string
	Port     int
}

var endpoints []Endpoint = []Endpoint{
	{
		Protocol: "HTTP",
		Host:     "localhost",
		Port:     80},
	{
		Protocol: "SSH",
		Host:     "10.10.5.9.xip.io",
		Port:     22}}

func ListEndpoints(startIndex int) {
	for index := range endpoints[startIndex:] {
		endpoint := endpoints[index]
		fmt.Printf("Priority: %d Procotol: %s Address: %s:%d\n",
			index, endpoint.Protocol, endpoint.Host, endpoint.Port)
	}
}
