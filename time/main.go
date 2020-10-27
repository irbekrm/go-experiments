package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	s := roundCurrentTime()
	fmt.Printf("'%T' '%v'\n", s, s)
	t := timeFromRFC3339Timestamp(s)
	fmt.Printf("'%T' '%v'\n", t, t)
}

// returns a string representing an RFC3339 timestamp of the current time (wall clock) to a minute's precision
func roundCurrentTime() string {
	t := time.Now()
	// Time.Round also removes the monotonic clock reading
	return t.Round(time.Minute).Format(time.RFC3339)
}

func timeFromRFC3339Timestamp(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		log.Fatalf("error parsing time: %v", err)
	}
	return t
}
