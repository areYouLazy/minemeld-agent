package agent

import (
	"regexp"
	"strings"
)

//Validate check if list output is valid
func Validate(list []byte) bool {
	errorString := regexp.MustCompile(`Unknown\sfeed`)

	//if match is true, than web server response is NOT valid
	if res := errorString.Match(list); res == true {
		return false
	}

	return true
}

//Parse turns a list of IP from Minemeld as []string or error
func Parse(lst []byte) []string {
	errorString := regexp.MustCompile(`Unknown\sfeed`)
	res := errorString.Match(lst)
	if res == true {
		return nil
	}

	stringList := string(lst)
	addressList := strings.Split(stringList, "\n")

	return addressList
}
