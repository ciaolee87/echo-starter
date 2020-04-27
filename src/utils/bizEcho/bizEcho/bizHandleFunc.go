package bizEcho

type (
	BizHandleFunc func(*BizContext) error
)

var PageNotFoundHandler BizHandleFunc = func(ctx *BizContext) error {
	return ctx.BizJson(&BizJSON{
		Code: 404,
		Msg:  "Page Not Found",
		Body: nil,
	})
}
