package log

import "os"

//Debug returns output in Debug mode
func Debug(format string, args ...interface{}) {
	log(0, format, args...)
}

//Info returns output in Info mode
func Info(format string, args ...interface{}) {
	log(1, format, args...)
}

//Warning returns output in Warning mode
func Warning(format string, args ...interface{}) {
	log(2, format, args...)
}

//Critical returns output in Critical mode
func Critical(format string, args ...interface{}) {
	log(3, format, args...)
}

//Fatal returns output in Fatal mode, and call os.Exit(1)
func Fatal(format string, args ...interface{}) {
	log(4, format, args...)
	os.Exit(1)
}

//Bold returns output in Bold mode
func Bold(s string) string {
	return bold(s)
}

//Dim returns output in Dim mode
func Dim(s string) string {
	return dim(s)
}
