package id_gen

import (
	"fmt"
	"testing"
	"time"
)

func TestNewEntityId(t *testing.T) {
	id := NewId()
	fmt.Println(id)
	if id == 0 {
		t.Errorf("Expected non-zero entity id, got 0")
	}

	latestId := id
	// for i := 0; i < 10; i++ {
	for range 10 {
		id = NewId()
		fmt.Println(id)

		if id <= latestId {
			t.Errorf("Expected new entity id to be greater or equal to latest id, got %d", id)
		}

		latestId = id
	}
}

func TestGetEpochTime(t *testing.T) {
	epochTime := time.Date(2005, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli()
	fmt.Println(epochTime) // print 1104537600000
}
