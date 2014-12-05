package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocallizeTime(t *testing.T) {
	dt := "Thu, 03 Jul 2014 10:10:01 -0700"
	localDt, _ := LocalizeTime(dt)
	// EST
	assert.Equal(t, localDt, "Thu, 3 Jul 2014 13:10:01 -0400")
}
