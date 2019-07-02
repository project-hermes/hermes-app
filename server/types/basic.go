package types

import "time"

func FromMilliseconds(milliseconds int64) time.Time {
	if milliseconds == 0 {
		return time.Time{}
	}

	return time.Unix(milliseconds/1000, milliseconds%1000*1000000)
}
