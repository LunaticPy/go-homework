package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_makeSet(t *testing.T) {
	data := []string{
		"листок",
		"тяпка",
		"пятка",
		"пятак",
		"Пятак",
		"пятак",
		"слиток",
		"столик",
	}

	Out := map[string][]string{
		"листок": {"листок", "слиток", "столик"},
		"тяпка":  {"пятак", "пятка", "тяпка"},
	}

	rez := makeSet(data)
	assert.Equal(t, Out, rez)
}
