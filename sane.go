// Package sane converts time formats from php-style to go-style
package sane

func TimeFormat(saneFormat string) (goFormat string) {
	goFormat = saneFormat

	for i := len(goFormat) - 1; i > -1; i-- {
		insane := mappedChars[string(goFormat[i])]

		if len(insane) != 0 {
			goFormat = goFormat[:i] + insane + goFormat[i+1:]
		}
	}

	return goFormat
}

var mappedChars = map[string]string{
	// year
	"y": "06",
	"Y": "2006",

	// month
	"n": "1",
	"m": "01",
	"M": "Jan",
	"F": "January",

	// day
	"j": "2",
	"d": "02",
	"D": "Mon",
	"l": "Monday",

	// time
	"a": "pm",
	"A": "PM",
	"g": "3",
	"h": "03",
	"H": "15",
	"i": "04",
	"s": "05",
	"v": "000",
	"u": "000000",

	// timezone
	"T": "MST",
	"O": "-0700",
	"Q": "-07:00",
}
