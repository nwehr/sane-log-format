package sane

import (
	"testing"
)

func TestFourDigitYear(t *testing.T) {
	expected := "2006"
	got := TimeFormat("Y")

	if got != expected {
		t.Errorf("Expected %s; got %s", expected, got)
	}
}

func TestTwoDigitYear(t *testing.T) {
	expected := "06"
	got := TimeFormat("y")

	if got != expected {
		t.Errorf("Expected %s; got %s", expected, got)
	}
}

func TestNumericMonth(t *testing.T) {
	expected := "1"
	got := TimeFormat("n")

	if got != expected {
		t.Errorf("Expected %s; got %s", expected, got)
	}
}

func TestNumericMonthLeadingZero(t *testing.T) {
	expected := "01"
	got := TimeFormat("m")

	if got != expected {
		t.Errorf("Expected %s; got %s", expected, got)
	}
}

func TestAbbreviatedMonth(t *testing.T) {
	expected := "Jan"
	got := TimeFormat("M")

	if got != expected {
		t.Errorf("Expected %s; got %s", expected, got)
	}
}

func TestMonth(t *testing.T) {
	expected := "January"
	got := TimeFormat("F")

	if got != expected {
		t.Errorf("Expected %s; got %s", expected, got)
	}
}

func TestDay(t *testing.T) {
	expected := "2"
	got := TimeFormat("j")

	if got != expected {
		t.Errorf("Expected %s; got %s", expected, got)
	}
}

func TestDayLeadingZero(t *testing.T) {
	expected := "02"
	got := TimeFormat("d")

	if got != expected {
		t.Errorf("Expected %s; got %s", expected, got)
	}
}

func TestAbbreviatedWeekday(t *testing.T) {
	expected := "Mon"
	got := TimeFormat("D")

	if got != expected {
		t.Errorf("Expected %s; got %s", expected, got)
	}
}

func TestWeekday(t *testing.T) {
	expected := "Monday"
	got := TimeFormat("l")

	if got != expected {
		t.Errorf("Expected %s; got %s", expected, got)
	}
}

func TestAmPmLowercase(t *testing.T) {
	expected := "pm"
	got := TimeFormat("a")

	if got != expected {
		t.Errorf("Expected %s; got %s", expected, got)
	}
}

func TestAmPmUppercase(t *testing.T) {
	expected := "PM"
	got := TimeFormat("A")

	if got != expected {
		t.Errorf("Expected %s; got %s", expected, got)
	}
}

func Test12Hour(t *testing.T) {
	expected := "3"
	got := TimeFormat("g")

	if got != expected {
		t.Errorf("Expected %s; got %s", expected, got)
	}
}

func Test12HourLeadingZero(t *testing.T) {
	expected := "03"
	got := TimeFormat("h")

	if got != expected {
		t.Errorf("Expected %s; got %s", expected, got)
	}
}

func Test24HourLeadingZero(t *testing.T) {
	expected := "15"
	got := TimeFormat("H")

	if got != expected {
		t.Errorf("Expected %s; got %s", expected, got)
	}
}

func TestMinute(t *testing.T) {
	expected := "04"
	got := TimeFormat("i")

	if got != expected {
		t.Errorf("Expected %s; got %s", expected, got)
	}
}

func TestSecond(t *testing.T) {
	expected := "05"
	got := TimeFormat("s")

	if got != expected {
		t.Errorf("Expected %s; got %s", expected, got)
	}
}

func TestTimezone(t *testing.T) {
	expected := "MST"
	got := TimeFormat("T")

	if got != expected {
		t.Errorf("Expected %s; got %s", expected, got)
	}
}

func TestGMTOffset(t *testing.T) {
	expected := "-0700"
	got := TimeFormat("O")

	if got != expected {
		t.Errorf("Expected %s; got %s", expected, got)
	}
}

func TestGMTOffsetAlternate(t *testing.T) {
	expected := "-07:00"
	got := TimeFormat("Q")

	if got != expected {
		t.Errorf("Expected %s; got %s", expected, got)
	}
}
