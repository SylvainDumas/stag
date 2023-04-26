package stag_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/SylvainDumas/stag"
	"github.com/stretchr/testify/assert"
)

// _______________________ Struct test _______________________

const tagTest = "tag_test"

type structFieldTest struct {
	Fn func(string, stag.FieldIf) error
}

func (obj *structFieldTest) Tag() string {
	return tagTest
}

func (obj *structFieldTest) Do(tagContent string, field stag.FieldIf) error {
	return obj.Fn(tagContent, field)
}

// _______________________ _______________________

func TestBrowseNil(t *testing.T) {
	// Check for nil entry
	assert.NoError(t, stag.Browse(nil))
}

func TestParseStructVar(t *testing.T) {
	var var1 struct {
		person struct {
			Age int `tag_test:"42"`
		}
	}

	var expectedValueFn = func(tagContent string, field stag.FieldIf) error {
		assert.Equal(t, "42", tagContent, "tag content")

		// Check parent parent
		assert.Nil(t, field.Parent().Parent(), "field parent has no parent")

		// Check parent
		assert.Equal(t, "person", field.Parent().Name(), "field parent name")
		assert.Equal(t, reflect.Struct.String(), field.Parent().Type(), "field parent type")

		// Check field
		assert.Equal(t, "Age", field.Name(), "field name")
		assert.Equal(t, reflect.Int.String(), field.Type(), "field type")
		assert.EqualValues(t, 0, field.Value().Int(), "field value")
		return nil
	}
	assert.NoError(t, stag.Browse(&var1, stag.WithTagActor(&structFieldTest{Fn: expectedValueFn})))
}

func TestParseStructVarPointer(t *testing.T) {
	type identity struct {
		Age int `tag_test:"42"`
	}

	var person struct {
		id *identity
	}

	assert.NoError(t, stag.Browse(&person, stag.WithTagActor(&structFieldTest{Fn: func(s string, fi stag.FieldIf) error {
		return errors.New("must not be called")
	}})))

	var expectedAge = 25
	person.id = &identity{Age: expectedAge}
	var expectedValueFn = func(tagContent string, field stag.FieldIf) error {
		assert.Equal(t, "42", tagContent, "tag content")

		// Check parent parent
		assert.Nil(t, field.Parent().Parent(), "field parent has no parent")

		// Check parent
		assert.Equal(t, "id", field.Parent().Name(), "field parent name")
		assert.Equal(t, reflect.Pointer.String(), field.Parent().Type(), "field parent type")

		// Check field
		assert.Equal(t, "Age", field.Name(), "field name")
		assert.Equal(t, reflect.Int.String(), field.Type(), "field type")
		assert.EqualValues(t, expectedAge, field.Value().Int(), "field value")
		return nil
	}
	assert.NoError(t, stag.Browse(&person, stag.WithTagActor(&structFieldTest{Fn: expectedValueFn})))
}

func TestParseStructHerited(t *testing.T) {
	type herited struct {
		age int `tag_test:"42"`
	}

	var varWithHerited struct {
		herited
	}

	var expectedValueFn = func(tagContent string, field stag.FieldIf) error {
		assert.Equal(t, "42", tagContent, "tag content")

		// Check parent parent
		assert.Nil(t, field.Parent(), "field has no parent")

		// Check field
		assert.Equal(t, "age", field.Name(), "field name")
		assert.Equal(t, reflect.Int.String(), field.Type(), "field type")
		assert.EqualValues(t, 0, field.Value().Int(), "field value")
		return nil
	}
	assert.NoError(t, stag.Browse(&varWithHerited, stag.WithTagActor(&structFieldTest{Fn: expectedValueFn})))
}
