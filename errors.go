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
	return fmt.Sprintf("struct value expected (type=%s pointer=%v)", obj.RefValue.Kind(), obj.IsPointer)
}
