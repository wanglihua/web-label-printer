package id_gen

import (
	"sync"
	"time"
)

// timestampBegin is the epoch: 2025-01-01 00:00:00 UTC, used to reduce the high bits of the id.
// var timestampBegin = time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli() // it's actually 1735689600000
const timestampBegin int64 = 1735689600000

var (
	lastTimestamp int64 = -1
	sequence      int64 = 0
	mu            sync.Mutex
)

// NewId generates a unique, roughly time-ordered int64 id.
// The lower 12 bits are a sequence number (0–4095), and the higher bits are the
// number of milliseconds since 2025-01-01 UTC. It is safe for concurrent use.
func NewId() int64 {
	mu.Lock()
	defer mu.Unlock()

	var timestamp = getTimestamp()

	if timestamp != lastTimestamp {
		sequence = 0
	}

	sequence++

	if sequence >= 4096 { // 2^12 = 4096
		sequence = 1
		// for timestamp <= lastTimestamp { // if the clock goes backwards, this would wait forever
		for timestamp == lastTimestamp {
			timestamp = getTimestamp()
		}
	}

	lastTimestamp = timestamp

	return (timestamp << 12) | sequence
}

func getTimestamp() int64 {
	return time.Now().UTC().UnixMilli() - timestampBegin
}
