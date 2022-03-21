package options

// SealOptions the global options
type SealOptions struct {
	EncodeHook EncodeHookFunc
	ExecLog    ExecLogFunc
	BuildLog   BuildLogFunc
}

// SealOptionsFunc SealOptions
type SealOptionsFunc func(opt *SealOptions)

// DefaultSealOptions only contains the EncodeHook
func DefaultSealOptions() *SealOptions {
	return &SealOptions{
		EncodeHook: Time2StringEncodeHook,
	}
}

// WithExecLogger set execlog func
func WithExecLogger(elog ExecLogFunc) SealOptionsFunc {
	return func(opt *SealOptions) {
		opt.ExecLog = elog
	}
}

// WithEncodeHook set encode hook. we can set nil to encode hook func if we want drop the default encode hook
func WithEncodeHook(ehook EncodeHookFunc) SealOptionsFunc {
	return func(opt *SealOptions) {
		opt.EncodeHook = ehook
	}
}

// WithBuildLogger set the sql build log
func WithBuildLogger(blog BuildLogFunc) SealOptionsFunc {
	return func(opt *SealOptions) {
		opt.BuildLog = blog
	}
}
