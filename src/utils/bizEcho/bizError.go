package bizEcho

type BizError struct {
	Error    error
	Body     *BizJSON
	Callback *func()
}
