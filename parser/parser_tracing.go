package parser

import (
	"fmt"
	"strings"
)

var traceLevel = 0

const traceIdentPlaceholder = "\t"

func identLevel() string {
	return strings.Repeat(traceIdentPlaceholder, traceLevel-1)
}

func tracePrint(fs string) {
	fmt.Printf("%s%s\n",
		identLevel(),
		fs,
	)
}

func incIdentifier() { traceLevel++ }
func decIdentifier() { traceLevel-- }

func trace(message string) string {
	incIdentifier()
	tracePrint("BEGIN " + message)
	return message
}

func untrace(message string) {
	tracePrint("END " + message)
	decIdentifier()
}
