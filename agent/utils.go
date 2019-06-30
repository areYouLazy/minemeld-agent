package agent

import (
	"encoding/json"

	"github.com/areYouLazy/minemeld-agent/log"
)

var (
	//AnchorList exposes all anchor declared in urls.json
	AnchorList []string
)

//CanAppendToAnchor returns true if given anchor is not in AnchorList
func CanAppendToAnchor(anchor string) bool {
	for _, v := range AnchorList {
		if v == anchor {
			return false
		}
	}

	return true
}

//GetAnchorListAsJSON populate AnchorList on-demand with all anchors
//from IP/FQDN lists
func GetAnchorListAsJSON() []byte {
	AnchorList = nil

	//iterate endpoints
	for _, v := range Endpoints {
		//verify to generate a list with unieuq entries
		if res := CanAppendToAnchor(v.Anchor); res == true {
			//Ensure anchor is not empty
			if len(v.Anchor) > 0 {
				AnchorList = append(AnchorList, v.Anchor)
			}
		}
	}

	//marshal to serve a JSON response
	res, err := json.Marshal(AnchorList)
	if err != nil {
		log.Debug("%s", err)
		return nil
	}

	return res
}
