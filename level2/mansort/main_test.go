package main

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	testCases := []struct {
		Name, In, Out string
		Atr           map[string]interface{}
	}{
		{
			Name: "main",
			In:   "content/main.txt",
			Out:  "test/main.txt",
			Atr:  make(map[string]interface{}),
		},
		{
			Name: "reverse",
			In:   "content/main.txt",
			Out:  "test/reverse.txt",
			Atr: map[string]interface{}{
				"reverse": true,
			},
		},
		{
			Name: "numeric",
			In:   "content/num.txt",
			Out:  "test/num.txt",
			Atr: map[string]interface{}{
				"numeric": true,
			},
		},
		{
			Name: "unic",
			In:   "content/unic.txt",
			Out:  "test/unic.txt",
			Atr: map[string]interface{}{
				"unic": true,
			},
		},
		{
			Name: "b",
			In:   "content/main.txt",
			Out:  "test/b.txt",
			Atr: map[string]interface{}{
				"b": true,
			},
		},
		{
			Name: "isSorted",
			In:   "content/main.txt",
			Out:  "test/issorted.txt",
			Atr: map[string]interface{}{
				"c": true,
			},
		},
		{
			Name: "natural",
			In:   "content/natural.txt",
			Out:  "test/natural.txt",
			Atr: map[string]interface{}{
				"h": true,
			},
		},
		{
			Name: "column",
			In:   "content/col.txt",
			Out:  "test/col.txt",
			Atr: map[string]interface{}{
				"col": 2,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			content, err := ioutil.ReadFile(tc.In)
			assert.Nil(t, err)
			out, err := ioutil.ReadFile(tc.Out)
			assert.Nil(t, err)
			control := strings.Split(string(out), "\n")
			AtrSwich(tc.Atr)
			rez := MySort(content)
			AtrReboot()
			assert.Equal(t, control, rez)
		})
	}

}

func AtrSwich(atr map[string]interface{}) {
	for i, val := range atr {
		switch i {
		case "reverse":
			reverse = val.(bool)
		case "numeric":
			numeric = val.(bool)
		case "unic":
			unic = val.(bool)
		case "mnth":
			mnth = val.(bool)
		case "b":
			b = val.(bool)
		case "c":
			c = val.(bool)
		case "h":
			h = val.(bool)
		case "col":
			col = val.(int)
		case "colSeparator":
			colSeparator = val.(string)
		}
	}

}

func AtrReboot() {
	reverse, numeric, unic, mnth, b, c, h = false, false, false, false, false, false, false
	col = 0
	colSeparator = " "
}
