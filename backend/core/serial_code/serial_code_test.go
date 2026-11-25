package serial_code

import (
	"fmt"
	"testing"
)

func TestCreateSerialCode(t *testing.T) {
	var code1 = CreateSerialCode()
	fmt.Println(code1)

	var code2 = CreateSerialCode()
	fmt.Println(code2)
}
