package gigasecond

import "time"

const testVersion = 4

func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Duration(1e9)*time.Second)
	return t
}
