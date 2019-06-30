package agent

import "regexp"

//Validate check if list output is valid
func Validate(list []byte) bool {
	//General Apache validation
	errorString := regexp.MustCompile(`Not\sFound`)
	if res := errorString.Match(list); res == true {
		return false
	}

	//MineMeld Response validation
	errorString = regexp.MustCompile(`Unknown\sfeed`)
	if res := errorString.Match(list); res == true {
		return false
	}

	//Return Valid response
	return true
}
