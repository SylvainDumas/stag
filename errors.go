package stag

import (
	"fmt"
	"reflect"
)

type BadValueErr struct {
	RefValue  reflect.Value
	IsPointer bool
}

func (obj *BadValueErr) Error() string {
	//return "struct value expected"
	return fmt.Sprintf("struct value expected, got %s (pointer: %v)", obj.RefValue.Kind(), obj.IsPointer)
}
