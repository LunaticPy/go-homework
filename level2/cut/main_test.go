package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCut(t *testing.T) {
	testCases := []struct {
		Name    string
		In, Out []string
		Atr     map[string]interface{}
	}{
		{
			Name: "main",
			In: []string{
				"1dat\tsad\tdddddd",
				"2dat\tsad\tdddddd",
				"3dat\tsad\tdddddd",
				"4dat\tsad\tdddddd",
			},
			Out: []string{
				"1dat\tsad\t",
				"2dat\tsad\t",
				"3dat\tsad\t",
				"4dat\tsad\t",
			},
			Atr: map[string]interface{}{
				"fields": "1,2",
			},
		},
		{
			Name: "separated",
			In: []string{
				"1dat sad dddddd",
				"2dat\tsad\tdddddd",
				"3dat\tsad\tdddddd",
				"4dat\tsad\tdddddd",
			},
			Out: []string{
				"2dat\tsad\t",
				"3dat\tsad\t",
				"4dat\tsad\t",
			},
			Atr: map[string]interface{}{
				"fields":    "1,2",
				"separated": true,
			},
		},
		{
			Name: "delim",
			In: []string{
				"1dat sad dddddd",
				"2dat sad dddddd",
				"3dat sad dddddd",
				"4dat sad dddddd",
			},
			Out: []string{
				"1dat sad ",
				"2dat sad ",
				"3dat sad ",
				"4dat sad ",
			},
			Atr: map[string]interface{}{
				"fields": "1,2",
				"delim":  " ",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			AtrSwich(tc.Atr)
			prep := parse(tc.In)
			rez := cut(prep)
			AtrReboot()
			assert.Equal(t, tc.Out, rez)
		})
	}

}

func AtrSwich(atr map[string]interface{}) {
	for i, val := range atr {
		switch i {
		case "separated":
			separated = val.(bool)
		case "fields":
			fields = val.(string)
		case "delim":
			delim = val.(string)
		}
	}

}

func AtrReboot() {
	separated = false
	fields = "-1"
	delim = "	"
}
