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

func TestOptions(t *testing.T) {
	t.Run("Set options with nil", func(t *testing.T) {
		var browser = new(browser).setOptions(nil, nil)
		assert.Len(t, browser.tagProcessorsFn, 0)
	})

	t.Run("Set options with nil and fonction", func(t *testing.T) {
		const expectedKey = "foo"
		var browser = new(browser).setOptions(nil, WithTagFn(expectedKey, func(tagContent string, field FieldIf) error { return nil }))
		assert.Len(t, browser.tagProcessorsFn, 1)
		assert.Contains(t, browser.tagProcessorsFn, expectedKey)
	})

	t.Run("Set options with no tag name", func(t *testing.T) {
		var browser = new(browser).setOptions(WithTagFn("", func(tagContent string, field FieldIf) error { return nil }))
		assert.Len(t, browser.tagProcessorsFn, 0)
	})

	t.Run("Set options with no tag fonction", func(t *testing.T) {
		var browser = new(browser).setOptions(nil, WithTagFn("foo", nil))
		assert.Len(t, browser.tagProcessorsFn, 0)
	})
}
