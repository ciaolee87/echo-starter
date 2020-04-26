package bizEcho

import "errors"

type BizError struct {
	Error    error
	Body     *BizJSON
	Callback func(ctx *BizContext)
}

var (
	Error2000RequestIdNotFound = BizError{
		Error: errors.New("RequestId not found!!"),
		Body: &BizJSON{
			Code: 2000,
			Msg:  "Generate RequestId is fail",
		},
		Callback: func(ctx *BizContext) {
		},
	}
)
