package stag

import "reflect"

type FieldLevel struct {
	Parent  FieldIf
	Current FieldIf
}

type FieldIf interface {
	Parent() FieldIf
	Name() string
	//FQDN() string
	Type() string
	Value() reflect.Value
	//Set(any) error
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

func (obj *Field) ReflectKind() reflect.Kind {
	return obj.value.Kind()
}

func (obj *Field) Type() string {
	return obj.value.Kind().String()
	//return obj.field.Type.Name() // + " " + obj.field.Type.String()
}

func (obj *Field) Value() reflect.Value {
	return obj.value
}
