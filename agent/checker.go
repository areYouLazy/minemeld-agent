package agent

import (
	"bytes"
	"net"
	"regexp"
	"strings"

	"github.com/areYouLazy/minemeld-agent/log"
)

//CheckIPv4 check if a given address is in IPv4List
func CheckIPv4(address, anchor string) bool {
	if anchor != "" {
		log.Debug("Received query for IPv4 address %s on list %s", log.Bold(address), log.Bold(anchor))
	} else {
		log.Debug("Received query for IPv4 address %s", log.Bold(address))
	}

	isInList := false

	IP := net.ParseIP(address)
	if IP == nil {
		log.Critical("Failed to parse IPv4 address %s", log.Bold(address))
		return isInList
	}

	//iterate list
	for _, payload := range IPv4List {
		//split payload in anchor and ip
		for k, v := range payload {
			//Check if line is in format of startIP-endIP
			if strings.Contains(v, "-") {
				vals := strings.Split(v, "-")
				StartIP := net.ParseIP(vals[0])
				EndIP := net.ParseIP(vals[1])
				if bytes.Compare(IP, StartIP) >= 0 && bytes.Compare(IP, EndIP) <= 0 {
					isInList = true
				}
				//Check if line is in format ip/mask
			} else if strings.Contains(v, "/") {
				_, net, _ := net.ParseCIDR(v)
				res := net.Contains(IP)
				if res == true {
					isInList = true
				}
			}

			//check if we have an anchor to match
			if anchor != "" {
				if anchor != k {
					isInList = false
				}
			}
		}
	}

	return isInList
}

//CheckIPv6 check if a given address is in IPv6List
func CheckIPv6(address string) bool {
	log.Debug("Received query for IPv6 address %s", log.Bold(address))

	isInList := false

	IP := net.ParseIP(address)
	if IP == nil {
		log.Critical("Failed to parse IPv6 address %s", log.Bold(address))
		return isInList
	}

	for _, v := range IPv6List {
		if strings.Contains(v, "-") {
			vals := strings.Split(v, "-")
			StartIP := net.ParseIP(vals[0])
			EndIP := net.ParseIP(vals[1])
			if bytes.Compare(IP, StartIP) >= 0 && bytes.Compare(IP, EndIP) <= 0 {
				isInList = true
			}
		}
	}

	return isInList
}

//CheckFQDN check if a given address is in FQDNList
func CheckFQDN(address string) bool {
	log.Debug("Received query for FQDN address %s", log.Bold(address))

	isInList := false

	for _, v := range FQDNList {
		reg := regexp.MustCompile(address)
		if reg.MatchString(v) {
			isInList = true
		}
	}

	return isInList
}
