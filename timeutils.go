package goutils

import (
	"errors"
	"time"
)

func ParseTime(s string) (time.Time, error) {

	layouts := []string{
		time.UnixDate,
		"2006-01-02 15:04:05.999-0700",
		time.ANSIC,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
	}

	for _, l := range layouts {
		if t, err := time.Parse(l, s); err == nil {
			return t, nil
		}
	}

	return time.Unix(0, 0), errors.New("one never has enough layouts")
}
