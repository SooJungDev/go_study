package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for _, ip := range hosts {
		stringIp := make([]string, 4)
		for i, numIp := range ip {
			stringIp[i] = strconv.Itoa(int(numIp))
		}

		ipAddr := strings.Join(stringIp, ".")
		fmt.Println(ipAddr)
	}
}
