package serial_code

import (
	"encoding/binary"
	"math"
	"strings"
	"uuid"

	"printer/core/id_gen"

	"github.com/lithammer/shortuuid/v5"
)

var encoderChars = "23456789ABCDEFGHJKLMNPQRSTUVWXYabcdefghijkmnpqrstuvwxy"
var encoder = shortuuid.NewEncoder(encoderChars)

// CreateSerialCode creates an increasing serial code, which can be used for various codes, including order numbers, invitation codes, and uploaded image filenames.
func CreateSerialCode() string {
	var newId = id_gen.NewId()
	var newIdBytes = make([]byte, 8)
	binary.LittleEndian.PutUint64(newIdBytes, uint64(newId))

	var l = len(newIdBytes)
	var newIdByteArray [16]byte

	for i := range 16 {
		newIdByteArray[i] = 0
	}

	var minLen = int(math.Min(float64(l), 16))
	for i := range minLen {
		newIdByteArray[16-i-1] = newIdBytes[i]
	}

	var code = encoder.Encode(uuid.UUID(newIdByteArray))
	return strings.TrimLeft(code, string(encoderChars[0]))

	// return serialCodeEncoder.Encode(uuid.NewV7()) -- this implementation also works and is simpler, but the code is longer, 22 characters
}
