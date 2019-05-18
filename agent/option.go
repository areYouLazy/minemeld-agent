package agent

import (
	"flag"
)

var (
	//Opt exports parsed options object
	Opt *Options
)

//Options holds value of flags from CLI
type Options struct {
	URLFile       *string
	FetchInsecure *bool
	Port          *int
}

//ParseOptions parse given options from CLI
func ParseOptions() {
	opt := &Options{
		URLFile:       flag.String("url-file", "urls.json", "Path of the JSON file containing urls"),
		FetchInsecure: flag.Bool("fetch-insecure", false, "Set to true to ignore certificate errors while fetching MineMeld URLs"),
		Port:          flag.Int("webserver-port", 9000, "Specify port for WebServer"),
	}

	//parse fags
	flag.Parse()

	//expose options object
	Opt = opt
}
