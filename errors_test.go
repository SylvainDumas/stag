package stag_test

import (
	"reflect"
	"testing"

	"github.com/SylvainDumas/stag"
	"github.com/stretchr/testify/assert"
)

func TestBadValueErr(t *testing.T) {
	var value = &stag.BadValueErr{}
	assert.EqualValues(t, "struct value expected (type=invalid pointer=false)", value.Error())

	value = &stag.BadValueErr{
		RefValue:  reflect.ValueOf(42),
		IsPointer: true,
	}
	assert.EqualValues(t, "struct value expected (type=int pointer=true)", value.Error())
}
