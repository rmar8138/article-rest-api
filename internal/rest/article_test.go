package rest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidDate(t *testing.T) {
	date := "2021-22-07"
	valid := validDate(date)
	assert.True(t, valid)

	date = "20212207"
	valid = validDate(date)
	assert.False(t, valid)
}

func TestValidUnhyphenatedDate(t *testing.T) {
	date := "20212207"
	valid := validUnhyphenatedDate(date)
	assert.True(t, valid)

	date = "120212207"
	valid = validUnhyphenatedDate(date)
	assert.False(t, valid)
}

func TestToHyphenatedDate(t *testing.T) {
	date := "20212207"
	hyphenated := toHyphenatedDate(date)
	assert.Equal(t, "2021-22-07", hyphenated)
}
