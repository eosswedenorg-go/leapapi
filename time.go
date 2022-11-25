package leapapi

import "time"

func fromTS(ts int64) time.Time {
	return time.Unix(ts/1000, ts%1000).UTC()
}
