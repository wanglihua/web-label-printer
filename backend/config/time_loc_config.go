package config

import (
	"time"
)

// var timeLoc *time.Location = time.FixedZone("UTC", 8*60*60)

func TimeLoc() *time.Location {
	// return timeLoc
	return time.Local
}
