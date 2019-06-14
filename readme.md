## Install ##

```
$ go get github.com/nwehr/sane-time-format
```

## Example Usage ##

```go
package main

import (
	"github.com/nwehr/sane-time-format"
	"time"
)

func main() {
	date := time.Now().Format(sane.TimeFormat("Y-m-d H:i:s.u T"))
	println(date) // outputs 2019-06-13 17:25:15.926299 PDT
}
```

## Format Characters ##

The go reference time is Mon Jan 2 15:04:05 MST 2006.

| Character | Description | Example |
| :-------- | :---------- | :------ |
| y | Two-digit year | 06 |
| Y | Four-digit year | 2006 |
| n | Numeric month | 1 |
| m | Numeric month (leading zero) | 01 |
| M | Abbreviated month | Jan |
| F | Month | January |
| j | Day of the month | 2 |
| d | Day of the month (leading zero) | 02 |
| D | Abbreviated day of the week | Mon |
| l | Day of the week | Monday |
| a | Ante/Post meridiem (lowercase) | pm |
| A | Ante/Post meridiem (uppercase) | PM |
| g | Hour (12-hour) | 3 |
| h | Hour (12-hour, leading zero) | 03 |
| H | Hour (24-hour, leading zero) | 15 |
| i | Minute (leading zero) | 04 |
| s | Second (leading zero) | 05 |
| T | Timzezone abbreviation | MST |
| O | GMT offset | -0700 |
| Q | GMT offset | -07:00 |