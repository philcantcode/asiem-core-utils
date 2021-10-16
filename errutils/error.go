package errutils

import "log"

func ParseErr(context string, e error) {

	if e != nil {
		log.Fatalf("%s: %v\n", context, e)
	}
}
