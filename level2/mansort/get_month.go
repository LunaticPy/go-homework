package main

import (
	"fmt"
	"strings"
)

var shortMonthNames = []string{
	"Jan",
	"Feb",
	"Mar",
	"Apr",
	"May",
	"Jun",
	"Jul",
	"Aug",
	"Sep",
	"Oct",
	"Nov",
	"Dec",
}

var longMonthNames = []string{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

func getMonth(m string) int {
	m = strings.ToLower(m)
	fmt.Println(m)
	for i := range shortMonthNames {
		if m == strings.ToLower(shortMonthNames[i]) || m == strings.ToLower(longMonthNames[i]) {
			return i + 1
		}
	}
	return 0
}
