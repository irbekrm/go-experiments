package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	// Turn the current time into string and back to time.Time again
	s := roundCurrentTime()
	fmt.Printf("'%T' '%v'\n", s, s)
	t := timeFromRFC3339Timestamp(s)
	fmt.Printf("'%T' '%v'\n", t, t)

	// Get a time.Time for a particular date in a timezone
	l := rigaLocation()
	// time.Time representing time in Riga on 19th July 1986
	tr := time.Date(1986, time.July, 19, 0, 0, 0, 0, l)
	fmt.Printf("'%T' '%v'\n", tr, tr)

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

func rigaLocation() *time.Location {
	// from https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
	l, err := time.LoadLocation("Europe/Riga")
	if err != nil {
		log.Fatal(err)
	}
	return l
}
