package gigasecond

import (
	"time"
)

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	date := t.Add(1e9 * time.Second)
	return date
}
