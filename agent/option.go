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
	IP      *string
	URLFile *string
}

//ParseOptions parse given options from CLI
func ParseOptions() {
	opt := &Options{
		URLFile: flag.String("url-file", "urls.json", "Path of the JSON file containing urls"),
	}

	//parse fags
	flag.Parse()

	//expose options object
	Opt = opt
}
