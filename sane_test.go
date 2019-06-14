package sane

import (
	"testing"
)

func TestSimpleDate(t *testing.T) {
	expected := "2006-01-02"
	got := TimeFormat("Y-m-d")

	if got != expected {
		t.Errorf("Expected %s; got %s", expected, got)
	}
}

func TestSimpleTimestamp(t *testing.T) {
	expected := "2006-01-02 15:04:05 MST"
	got := TimeFormat("Y-m-d H:i:s T")

	if got != expected {
		t.Errorf("Expected %s; got %s", expected, got)
	}
}
