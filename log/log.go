package log

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var (
	//OUTPUT flag holds output reference
	OUTPUT string

	//DEBUG flag is true if debug is requested from command-line flags
	DEBUG bool

	//COLORS flag is true if output should be colored
	COLORS bool

	//fh holds file if a custom output is specified
	fh *os.File

	//outputWriter holds the Writer reference for Output
	outputWriter *bufio.Writer

	//dateTimeFormat holds format for Date/Time output
	dateTimeFormat = "2006-01-02 15:04:05"

	//labels maps to translate logLevel into
	//a human readable label
	labels = map[int]string{
		0: "DBG",
		1: "INF",
		2: "WRN",
		3: "CRI",
		4: "FAT",
	}

	//colors maps to translate logLevel into
	//a human readable color based on severity
	colors = map[int]string{
		0: "\033[0;34m", // blue
		1: "\033[0;32m", // green
		2: "\033[0;33m", // yellow
		3: "\033[0;31m", // red
		4: "\033[0;37m", // gray
	}

	//resetCode holds code to reset output color
	resetCode = "\033[0;0m" // reset

	//dimCode holds code for dim output
	dimCode = "\033[0;2m" // dim

	//boldCode holds code for bold output
	boldCode = "\033[0;1m" //bold
)

//log default routine to get output
func log(level int, format string, args ...interface{}) {
	//Check if we have to return debug output
	if level == 0 && DEBUG == false {
		return
	}

	//defines temp variables
	var label, datetime, payload, message string

	//Get raw time for output
	rawTime := time.Now()

	//Parse time with our format
	datetime = rawTime.Format(dateTimeFormat)

	//Check if we need colored output
	if COLORS {
		label = "[" + colors[level] + labels[level] + resetCode + "]"
	} else {
		label = "[" + labels[level] + "]"
	}

	//Generate payload with given data
	payload = fmt.Sprintf(format, args...)

	//Generate clean output
	message = datetime + " " + label + " " + payload + "\n"

	//Check if we have a writer, or use os.Stdout
	switch {
	case OUTPUT == "":
		//instantiate os.Stdout as the Output Writer
		outputWriter = bufio.NewWriter(os.Stdout)
	case OUTPUT != "":
		fh, err := os.OpenFile(OUTPUT, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			//fall back to Stdout if we cannot write to the given output
			fmt.Printf("Cannot write to given output %s\n", OUTPUT)
			fmt.Printf("Falling back to Stdout\n")
			//instantiate os.Stdout as the Output Writer
			outputWriter = bufio.NewWriter(os.Stdout)
		} else {
			outputWriter = bufio.NewWriter(fh)
		}
	}

	//defer writer flush
	defer outputWriter.Flush()

	//Print Output to Writer
	_, err := fmt.Fprintf(outputWriter, message)
	if err != nil {
		fmt.Printf("Cannot write to output file: %s", OUTPUT)
	}
}

//dim return given output dimmed
func dim(s string) string {
	payload := dimCode + s + resetCode
	return payload
}

//bold return ginve output bolded
func bold(s string) string {
	payload := boldCode + s + resetCode
	return payload
}
