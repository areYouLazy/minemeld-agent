package agent

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/areYouLazy/minemeld-agent/log"
)

var (
	//Endpoints exposes Endpoint loaded from urls.json
	Endpoints []*Target
)

//Target structure to holds type/url from MineMeld
type Target struct {
	Type     string `json:"type"`
	Endpoint string `json:"endpoint"`
}

//Loader loads urls.json file and parse it to get IP/FQDN lists
func Loader() {
	log.Debug("Reading url file %s", log.Bold(*Opt.URLFile))

	//load json file
	jsonFile, err := os.Open(*Opt.URLFile)
	if err != nil {
		log.Fatal("Error reading URL File %s", log.Bold(*Opt.URLFile))
	}

	//defer file closure
	defer jsonFile.Close()

	//convert to bytes to read the file
	byteValue, _ := ioutil.ReadAll(jsonFile)

	//unmarshal file content to our list
	json.Unmarshal(byteValue, &Endpoints)

	for k, v := range Endpoints {
		if v.Type == "ipv4" || v.Type == "ipv6" || v.Type == "fqdn" {
			continue
		} else {
			log.Warning("Unknown Endpoint Type %s for endpoint %s. Check your URLFile for errors", log.Bold(v.Type), log.Bold(v.Endpoint))
			log.Debug("Endpoint %s removed from memory because of invalid Type %s", log.Bold(v.Endpoint), log.Bold(v.Type))
			Endpoints = append(Endpoints[:k], Endpoints[k+1:]...)
		}
	}
}
