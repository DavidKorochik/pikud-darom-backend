package testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Just a check testing (doesn't belong to our app)
func TestAdd(t *testing.T) {
	inputA := 1
	inputB := 2
	expected := 3
	sum := inputA + inputB

	assert.Equal(t, expected, sum)
}
