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
	//IPv4List exposes ipv4 fetched from Endpoint
	IPv4List []map[string]string

	//IPv6List exposes ipv4 fetched from Endpoint
	IPv6List []map[string]string

	//FQDNList exposes names fetched form Endpoint
	FQDNList []map[string]string
)

//Fetch get list from endpoint url and returns request body
func Fetch(url string) ([]byte, error) {
	log.Info("Fetching list from %s", log.Bold(url))

	if *Opt.FetchInsecure == true {
		//this prevents error because of insecure certificate
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	//set a global timeout for GET calls
	http.DefaultClient.Timeout = time.Duration(3 * time.Second)

	//Query given URL
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	//Close body after call ends
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, err
	}

	//parse body to bytes
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//return call body
	return body, nil
}

//FetchEndpoints routine to fetch lists from Endpoint
func FetchEndpoints() {
	for {
		//check urls.json for changes
		Loader()

		//reset lists to avoid duplicates
		IPv4List, IPv6List, FQDNList = nil, nil, nil

		log.Debug("Iterate Endpoints to fetch address list")
		//iterate endpoints to fetch address
		for _, v := range Endpoints {
			//fetch endpoint url
			lst, err := Fetch(v.Endpoint)
			if err != nil {
				log.Critical("%s", err)
				//validate response befor submission
			} else if res := Validate(lst); res == true {
				log.Info("Parsing response from %s", log.Bold(v.Endpoint))
				parsedList := Parse(lst)

				switch {
				case v.Type == "ipv4":
					for _, l := range parsedList {
						p := map[string]string{v.Anchor: l}
						IPv4List = append(IPv4List, p)
					}
				case v.Type == "ipv6":
					for _, l := range parsedList {
						p := map[string]string{v.Anchor: l}
						IPv6List = append(IPv6List, p)
					}
				case v.Type == "fqdn":
					for _, l := range parsedList {
						p := map[string]string{v.Anchor: l}
						FQDNList = append(FQDNList, p)
					}
				}

				log.Debug("Found %s entry in list %s", log.Bold(strconv.Itoa(len(parsedList))), log.Bold(v.Endpoint))
			} else {
				log.Warning("Validation failed for Endpoint %s response", log.Bold(v.Endpoint))
			}
		}

		time.Sleep(10 * time.Second)
	}
}
