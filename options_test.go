package stag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptions(t *testing.T) {
	t.Run("Set options with nil", func(t *testing.T) {
		var opts options
		opts.apply(nil, nil)
		assert.Len(t, opts.tagProcessorsFn, 0)
	})

	t.Run("Set options with nil and fonction", func(t *testing.T) {
		var opts options
		const expectedKey = "foo"
		opts.apply(nil, WithTagFn(expectedKey, func(tagContent string, field FieldIf) error { return nil }))
		assert.Len(t, opts.tagProcessorsFn, 1)
		assert.Contains(t, opts.tagProcessorsFn, expectedKey)
	})

	t.Run("Set options with no tag name", func(t *testing.T) {
		var opts options
		opts.apply(WithTagFn("", func(tagContent string, field FieldIf) error { return nil }))
		assert.Len(t, opts.tagProcessorsFn, 0)
	})

	t.Run("Set options with no tag fonction", func(t *testing.T) {
		var opts options
		opts.apply(nil, WithTagFn("foo", nil))
		assert.Len(t, opts.tagProcessorsFn, 0)
	})
}
