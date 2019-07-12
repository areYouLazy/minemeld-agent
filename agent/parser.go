package agent

import (
	"strings"
)

//TODO: We can do patter parsing

//Parse turns a list of IP from Endpoint as []string or error
func Parse(lst []byte) []string {
	//convert byte-slice to string
	stringList := string(lst)

	//split string in lines
	addressList := strings.Split(stringList, "\n")

	//remove empty entries to prevent wildcard matches
	for k, v := range addressList {
		if len(v) == 0 {
			addressList = append(addressList[:k], addressList[k+1:]...)
		}
	}

	//remove comment entries, that starts with #
	for k, v := range addressList {
		if res := strings.HasPrefix(v, "#"); res == true {
			addressList = append(addressList[:k], addressList[k+1:]...)
		}
	}

	return addressList
}
