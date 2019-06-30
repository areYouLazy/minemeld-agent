package agent

import (
	"strings"
)

//TODO: We can do patter parsing

//Parse turns a list of IP from Minemeld as []string or error
func Parse(lst []byte) []string {
	stringList := string(lst)
	addressList := strings.Split(stringList, "\n")
	for k, v := range addressList {
		if len(v) == 0 {
			addressList = append(addressList[:k], addressList[k+1:]...)
		}
	}

	return addressList
}
