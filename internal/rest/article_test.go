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
