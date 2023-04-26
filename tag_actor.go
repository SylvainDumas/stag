package stag

// Field value visitor

type ActionReader interface {
	Read(any) error
}

type ActionWriter interface {
	Write(any) (any, error)
}

type TagActorIf interface {
	Tag() string
	Do(tagContent string, field FieldIf) error
}
