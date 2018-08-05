package logger

import (
	"log"
	"runtime"
)

/* These functions are not currently used, but don't delete them yet!
	 They may be useful elsewhere.

// Gets a given line from a byte array.
func getLine(buf []byte, lineNumber int) string {
	bytesReader := bytes.NewReader(buf)
	bufReader := bufio.NewReader(bytesReader)
	var line []byte
	for i := 0; i < lineNumber; i++ {
		line, _, _ = bufReader.ReadLine()
	}
	return string(line)
}

// Gets the function name from a stack trace line.
func getFuncName(line string) string {
	periodIndex := strings.Index(line, ".")
	parenIndex := strings.Index(line, "(")
	return line[periodIndex+1 : parenIndex]
}

func getStackTraceFuncName(depth int) string {
	buf := make([]byte, 300) //TODO: Can this be more dynamic instead of hard-coding 300?
	bytes := runtime.Stack(buf, false)
	stackLine := getLine(buf[:bytes], depth*2)
	return getFuncName(stackLine)
}
*/

type any interface{}

/**
 * runtime.Frame is a struct that has many fields
 * including File, Line, and Function.
 * Function is a string that contains the
 * package name, a period, and a function name.
 */
func getFrame() runtime.Frame {
	pc := make([]uintptr, 15)
	n := runtime.Callers(3, pc) // going back 3 levels in the call stack
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return frame
}

// Log logs any number of values to stdout, including the
// calling file name, function name, and line number.
func Log(values ...any) {
	frame := getFrame()
	log.SetFlags(log.Lshortfile)
	log.Println(frame.Function, values)
}

// LogVar is similar to Log, but adds "=" between the two inputs.
func LogVar(name string, value any) {
	frame := getFrame() // has File, Line, and Function fields
	log.SetFlags(log.Lshortfile)
	log.Printf("%s %s=%+v\n", frame.Function, name, value)
}
