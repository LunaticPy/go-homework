package base

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	rTime := time.Now()
	testTime := Get()

	assert.Equal(t, rTime.Local().Hour(), testTime.Local().Hour())
	assert.Equal(t, rTime.Local().Minute(), testTime.Local().Minute())
	assert.Equal(t, rTime.Local().Second(), testTime.Local().Second())
}
