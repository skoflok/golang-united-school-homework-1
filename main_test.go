package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMessage(t *testing.T) {
	var e string = "Hello :world_map:!"
	m := getMessage()
	assert.Equal(t, e, m)
	if m != "Hello :world_map:!" {
		t.Errorf("Expect `%s`, got `%s`", e, m)
	}
}
