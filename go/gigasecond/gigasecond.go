package gigasecond

import (
	"math"
	"time"
)

const (
	TestVersion = 2
)

var Birthday = time.Date(1983, time.June, 11, 11, 0, 0, 0, time.UTC)

var gigaseconds = time.Duration(math.Pow(10, 9)) * time.Second

func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigaseconds)
}
