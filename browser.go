package stag

import (
	"reflect"
)

func Browse(dst any, options ...browserOption) error {
	return new(browser).setOptions(options...).browse(dst)
}

func BrowseSuggar(dst any, actors ...TagActorIf) error {
	var browserOptions []browserOption
	for _, v := range actors {
		if v != nil {
			browserOptions = append(browserOptions, WithTagActor(v))
		}
	}
	return new(browser).setOptions(browserOptions...).browse(dst)
}

// ________________________ config ________________________

type browserOption func(*browser)

func WithTagActor(actor TagActorIf) browserOption {
	return func(out *browser) {
		out.tagFn = append(out.tagFn, actor)
	}
}

//Toto option:
//- continueOnError (or set type error critical and not critical,)
//- lookup strategy ? (tag order)
//- option recursive or inDepth struct ?

// ________________________ browser ________________________

type browser struct {
	tagFn    []TagActorIf
	readOnly bool
}

func (obj *browser) setOptions(options ...browserOption) *browser {
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
	if len(obj.tagFn) == 0 {
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

		//fmt.Printf("%+v\n", fieldType)
		//fmt.Println(fieldValue.Kind(), fieldValue.Type(), fieldValue.CanSet())

		if fieldValue.IsValid() == false {
			continue
		}

		// Tag apply lookup for this field
		for _, v := range obj.tagFn {
			// tag actor struct with name
			// switch assert cast by type tag_reader, tag_writer, tag_deeper and call with specific fn

			if tagValue, found := fieldType.Tag.Lookup(v.Tag()); found {
				//fmt.Printf("%v(%v, %v) [%v]: %v - %v\n", fieldName, fieldType.Anonymous, fieldValue.CanSet(), fieldValue.Kind(), fieldValue, tagValue)
				if err := v.Do(tagValue, &Field{parent: parent, field: fieldType, value: fieldValue}); err != nil {
					return err
				}
			}
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
		//fmt.Printf("%v(%v, %v) [%v]: %v\n", fieldName, fieldType.Anonymous, fieldValue.CanSet(), fieldValue.Kind(), fieldValue)
		if err := obj.browseStructFields(parentField, structFieldValue); err != nil {
			return err
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
