package stag

import "reflect"

type FieldIf interface {
	Parent() FieldIf
	Name() string
	Type() string
	Value() reflect.Value
}

// _____________ Impl _____________

type Field struct {
	parent FieldIf
	field  reflect.StructField
	value  reflect.Value
}

func (obj *Field) Parent() FieldIf {
	return obj.parent
}

func (obj *Field) Name() string {
	return obj.field.Name
}

func (obj *Field) Type() string {
	return obj.value.Kind().String()
}

func (obj *Field) Value() reflect.Value {
	return obj.value
}
