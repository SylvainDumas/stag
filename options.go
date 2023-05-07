package stag

type optionFn func(*options)

type options struct {
	tagProcessorsFn map[string]TagProcessorFn
}

func (obj *options) apply(fn ...optionFn) {
	// Init
	if obj.tagProcessorsFn == nil {
		obj.tagProcessorsFn = make(map[string]TagProcessorFn)
	}
	// Apply options
	for _, v := range fn {
		if v != nil {
			v(obj)
		}
	}
}

// _______________________ options setter _______________________

func WithTagFn(tag string, fn TagProcessorFn) optionFn {
	return func(out *options) {
		if tag != "" && fn != nil {
			out.tagProcessorsFn[tag] = fn
		}
	}
}
