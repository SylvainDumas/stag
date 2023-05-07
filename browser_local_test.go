package stag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadOnlyDetection(t *testing.T) {
	// Check read only detection
	var bar struct{}

	t.Run("Check read only", func(t *testing.T) {
		var browser = new(browser)
		assert.NoError(t, browser.browse(bar))
		assert.True(t, browser.readOnly)
	})

	t.Run("Check writable", func(t *testing.T) {
		var browser = new(browser)
		assert.NoError(t, browser.browse(&bar))
		assert.False(t, browser.readOnly)
	})
}
