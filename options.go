package stag

func WithTagFn(tag string, fn TagProcessorFn) browserOption {
	return func(out *browser) {
		if tag != "" && fn != nil {
			out.tagProcessorsFn[tag] = fn
		}
	}
}
