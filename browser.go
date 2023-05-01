package stag

import (
	"reflect"
)

type browserOption func(*browser)

func Browse(dst any, options ...browserOption) error {
	return new(browser).setOptions(options...).browse(dst)
}

// ________________________ browser ________________________

type TagProcessorFn func(tagContent string, field FieldIf) error

type browser struct {
	tagProcessorsFn map[string]TagProcessorFn
	readOnly        bool
}

func (obj *browser) setOptions(options ...browserOption) *browser {
	if obj.tagProcessorsFn == nil {
		obj.tagProcessorsFn = make(map[string]TagProcessorFn)
	}
	for _, v := range options {
		if v != nil {
			v(obj)
		}
	}
	return obj
}

func (obj *browser) browse(dst any) error {
	// Nil, nothing to do
	if dst == nil {
		return nil
	}

	// Check is a struct
	valueRef, isPointer := getReflectedValue(reflect.ValueOf(dst))
	if valueRef.Kind() != reflect.Struct {
		return &BadValueErr{RefValue: valueRef, IsPointer: isPointer}
	}
	obj.readOnly = !isPointer

	// If no tag actors, nothing to do
	if len(obj.tagProcessorsFn) == 0 {
		return nil
	}

	// Browse struct fields
	return obj.browseStructFields(nil, valueRef)
}

func (obj *browser) browseStructFields(parent FieldIf, valueRef reflect.Value) error {
	var valueRefType = valueRef.Type()

	for i := 0; i < valueRefType.NumField(); i++ {
		// Get struct field info
		var fieldType = valueRefType.Field(i)
		var fieldName = fieldType.Name
		var fieldValue = valueRef.FieldByName(fieldName)

		if fieldValue.IsValid() == false {
			continue
		}

		// Tag apply lookup for this field
		if err := obj.applyFieldTagProcessors(fieldType.Tag, &Field{parent: parent, field: fieldType, value: fieldValue}); err != nil {
			return err
		}

		// If not a direct or indirect value struct, pass to next field
		structFieldValue, _ := getReflectedValue(fieldValue)
		if structFieldValue.Kind() != reflect.Struct {
			continue
		}
		var parentField = parent
		// If not a herited struct field, create a parent on current field
		if fieldType.Anonymous == false {
			parentField = &Field{parent: parent, field: fieldType, value: fieldValue}
		}
		// Anonymous, ptr on struct, array of struct, map of struct
		if err := obj.browseStructFields(parentField, structFieldValue); err != nil {
			return err
		}
	}

	return nil
}

func (obj *browser) applyFieldTagProcessors(structTag reflect.StructTag, field FieldIf) error {
	for k, v := range obj.tagProcessorsFn {
		if tagValue, found := structTag.Lookup(k); found {
			if err := v(tagValue, field); err != nil {
				return err
			}
		}
	}

	return nil
}

// getReflectedValue returns real reflected value and a bool for pointer on
func getReflectedValue(value reflect.Value) (reflect.Value, bool) {
	if value.Kind() == reflect.Ptr {
		return value.Elem(), true
	}
	return value, false
}
