package time_token

import (
	"fmt"
	"testing"
)

func TestCheckToken(t *testing.T) {
	ok, err := CheckToken("202202171311421")
	if err != nil {
		t.Fail()
	}

	if ok {

		t.Fail()
	}

	ok, err = CheckToken("202202171311421db96c1f690373f36f411ad509e55785")
	if err != nil {
		t.Fail()
	}

	if ok {

		t.Fail()
	}

	ok, err = CheckToken(CreateToken())
	if err != nil {
		t.Fail()
	}

	if !ok {
		t.Fail()
	}

}

func TestCreateToken(t *testing.T) {
	var token = CreateToken()
	fmt.Println(token) // 20260725103423dbd461f7ff2cebf65763ce2f7129f68f67bfc9285eb9c233922de13f62b22fc6

}
