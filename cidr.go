package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Usage: ./cidr <CIDR> [<CIDR> ...]")
		return
	}

	for _, cidrInput := range flag.Args() {
		_, ipNet, err := net.ParseCIDR(cidrInput)
		if err != nil {
			fmt.Printf("Error for %s: %s\n", cidrInput, err)
			continue
		}

		ips := getIPRange(ipNet)

		for _, ip := range ips {
			fmt.Println(ip)
		}
		fmt.Println()
	}
}

func getIPRange(ipNet *net.IPNet) []string {
	var ips []string

	for ip := ipNet.IP.Mask(ipNet.Mask); ipNet.Contains(ip); incrementIP(ip) {
		ips = append(ips, ip.String())
	}

	return ips
}

func incrementIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
