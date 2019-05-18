package agent

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/areYouLazy/minemeld-agent/log"
)

var (
	//IPv4List exposes ipv4 fetched from MineMeld
	IPv4List []string

	//IPv6List exposes ipv4 fetched from MineMeld
	IPv6List []string

	//FQDNList exposes names fetched form MineMeld
	FQDNList []string
)

//Fetch get list from minemeld url and returns request body
func Fetch(url string) ([]byte, error) {
	log.Info("Fetching list from %s", log.Bold(url))

	if *Opt.FetchInsecure == true {
		//this prevents error because of insecure certificate
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	//Query given URL
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	//Close body after call ends
	defer resp.Body.Close()

	//parse body to bytes
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//return call body
	return body, nil
}

//FetchEndpoints routine to fetch lists from MineMeld
func FetchEndpoints() {
	for {
		//check urls.json for changes
		Loader()

		log.Debug("Iterate Endpoints to fetch address list")
		//iterate endpoints to fetch address
		for _, v := range Endpoints {
			lst, err := Fetch(v.Endpoint)
			if err != nil {
				log.Critical("%s", err)
			} else if res := Validate(lst); res == true {
				log.Info("Parsing response from %s", log.Bold(v.Endpoint))
				switch {
				case v.Type == "ipv4":
					IPv4List = Parse(lst)
				case v.Type == "ipv6":
					IPv6List = Parse(lst)
				case v.Type == "fqdn":
					FQDNList = Parse(lst)
				}

				log.Debug("Found %s entry in list %s", log.Bold(strconv.Itoa(len(lst))), log.Bold(v.Endpoint))
			} else {
				log.Warning("Validation failed for Endpoint %s", log.Bold(v.Endpoint))
			}
		}

		time.Sleep(10 * time.Second)
	}
}
