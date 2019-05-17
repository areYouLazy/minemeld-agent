package log

import "flag"

//Init initialize log environment parsing given flags
func Init() {
	flag.StringVar(&OUTPUT, "log-output", "", "Set the output interface for log")
	flag.BoolVar(&DEBUG, "log-debug", false, "Set true to print debug messages")
	flag.BoolVar(&COLORS, "log-colors", true, "Set to false to turn of colored log output")

	//flag.Parse()
}
