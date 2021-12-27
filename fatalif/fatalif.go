// Package fatalif provides functions that will call log.Fatal with a
// contextual message if some condition is true.
package fatalif

import "log"

// If err is not null, log.Fatal with the error message and some context about
// the error.
func Error(err error, msg string) {
	if err != nil {
		log.Fatal("error ", msg, ": ", err)
	}
}
