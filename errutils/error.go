package errutils

import "log"

func ParseErr(e error, context string) {

	if e != nil {
		log.Fatalf("%s: %v\n", context, e)
	}
}
