/*
Package sane impliments a function for converting
php-style date/time format strings into go-style format strings
*/
package sane

var mappedChars map[string]string

// TimeFormat converts a sane, php-style format string to the equivilent go format
// TimeFormat("Y-m-d H:i:s O") returns "2006-01-02 15:04:05 MST"
func TimeFormat(saneFormat string) string {
	format := saneFormat

	for i := len(format) - 1; i >= 0; i-- {
		insane := mappedChars[string(format[i])]

		if len(insane) != 0 {
			format = format[:i] + insane + format[i+1:]
		}
	}

	return format
}

func init() {
	mappedChars = map[string]string{
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
}
