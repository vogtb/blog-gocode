package logger

import (
	"testing"
)

func TestHandlesNonExistingItems(t *testing.T) {
	log := Logger{}
	if "thing" != log.ReturnSame("thing") {
		t.Error("Failed")
	}
}
