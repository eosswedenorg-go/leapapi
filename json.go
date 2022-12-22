package leapapi

import (
	"time"

	jsontime "github.com/liamylian/jsontime/v2/v2"
)

var json = jsontime.ConfigWithCustomTimeFormat

func init() {
	// EOS Api does not specify timezone in timestamps (they are always UTC tho).
	jsontime.SetDefaultTimeFormat("2006-01-02T15:04:05", time.UTC)
}
