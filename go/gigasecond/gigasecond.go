//Package gigasecond provides function that can get date and add to the date 10^9 seconds and return the new date
package gigasecond

import (
	"time"
)

//AddGigasecond function add to passed date 10^9 seconds
func AddGigasecond(t time.Time) time.Time {
	const timeToAdd = time.Duration(1e9) * time.Second
	t = t.Add(timeToAdd)
	return t
}
