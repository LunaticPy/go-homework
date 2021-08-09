package unpack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnpuck(t *testing.T) {
	testCases := []struct {
		line, asn string
	}{
		{line: "a4bc2d5e",
			asn: "aaaaabcccdddddde"},
		{line: "abcd",
			asn: "abcd"},
		{line: "45",
			asn: ""},
		{line: "",
			asn: ""},
		{line: "qwe\\4\\5",
			asn: "qwe45"},
		{line: "qwe\\45",
			asn: "qwe44444"},
		{line: "qwe\\\\5",
			asn: `qwe\\\\\`},
	}

	for _, tc := range testCases {
		t.Run(tc.line, func(t *testing.T) {
			assert.Equal(t, tc.asn, Unpack(tc.line))

		})
	}

}
