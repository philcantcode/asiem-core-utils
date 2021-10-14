package errutils

import "log"

func parseErr(e error, context string) {

	if e != nil {
		log.Fatalf("%s: %v\n", context, e)
	}
}
