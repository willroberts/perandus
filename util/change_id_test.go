package util

import (
	"testing"
)

func TestGetLatestChangeID(t *testing.T) {
	out := GetLatestChangeID()
	if out == "" {
		t.Log("Failed to retrieve current change ID")
		t.FailNow()
	}
}
