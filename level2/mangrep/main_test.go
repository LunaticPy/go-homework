package main

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	testCases := []struct {
		Name, In, Out, Pattern string
		Atr                    map[string]interface{}
	}{
		{
			Name:    "main",
			In:      "data.txt",
			Out:     "test/main.txt",
			Pattern: "is",
			Atr:     make(map[string]interface{}),
		},
		{
			Name:    "A",
			In:      "data.txt",
			Out:     "test/A.txt",
			Pattern: "is",
			Atr: map[string]interface{}{
				"A": 2,
			},
		},
		{
			Name:    "B",
			In:      "data.txt",
			Out:     "test/B.txt",
			Pattern: "is",
			Atr: map[string]interface{}{
				"B": 2,
			},
		},
		{
			Name:    "C",
			In:      "data.txt",
			Out:     "test/C.txt",
			Pattern: "is",
			Atr: map[string]interface{}{
				"C": 2,
			},
		},
		{
			Name:    "count",
			In:      "data.txt",
			Out:     "test/count.txt",
			Pattern: "is",
			Atr: map[string]interface{}{
				"c": true,
			},
		},
		{
			Name:    "ignore",
			In:      "data.txt",
			Out:     "test/i.txt",
			Pattern: "is",
			Atr: map[string]interface{}{
				"I": true,
			},
		},
		{
			Name:    "verte",
			In:      "data.txt",
			Out:     "test/v.txt",
			Pattern: "is",
			Atr: map[string]interface{}{
				"v": true,
			},
		},
		{
			Name:    "F",
			In:      "data.txt",
			Out:     "test/F.txt",
			Pattern: "is",
			Atr: map[string]interface{}{
				"F": true,
			},
		},
		{
			Name:    "num",
			In:      "data.txt",
			Out:     "test/num.txt",
			Pattern: "is",
			Atr: map[string]interface{}{
				"num": true,
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
			rez := MyGrep(tc.Pattern, content)
			AtrReboot()
			assert.Equal(t, control, rez)
		})
	}

}

func AtrSwich(atr map[string]interface{}) {
	for i, val := range atr {
		switch i {
		case "I":
			I = val.(bool)
		case "v":
			v = val.(bool)
		case "F":
			F = val.(bool)
		case "num":
			num = val.(bool)
		case "c":
			c = val.(bool)
		case "A":
			A = val.(int)
		case "B":
			B = val.(int)
		case "C":
			C = val.(int)
		}
	}

}

func AtrReboot() {
	I, v, F, num, c = false, false, false, false, false
	A, B, C = 0, 0, 0
}
