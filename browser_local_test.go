package stag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadOnly(t *testing.T) {
	// Check read only detection
	var bar struct{}
	var browser = new(browser)
	assert.NoError(t, browser.browse(bar))
	assert.True(t, browser.readOnly)

	assert.NoError(t, browser.browse(&bar))
	assert.False(t, browser.readOnly)
}
