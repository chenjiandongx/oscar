package modules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModuler(t *testing.T) {
	assert.True(t, true)
	assert.Equal(t, "hi!", "hi!")
	assert.NotEqual(t, "Leonardos?", "No!")
}
