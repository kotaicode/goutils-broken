package goutils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTime(t *testing.T) {
	stamps := map[string]error{
		"Wed Jun 17 15:18:02 GMT 2020": nil,
		"2020-06-17 17:18:00.541+0200": nil,
		"2020-06-17 25:61:90.541-0200": errors.New("one never has enough layouts"),
		"not really a timestamp thing": errors.New("one never has enough layouts"),
	}

	for s, err := range stamps {
		_, e := ParseTime(s)
		assert.Equal(t, err, e)
	}
}
