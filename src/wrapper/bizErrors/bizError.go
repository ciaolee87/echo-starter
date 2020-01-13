package bizErrors

import (
	"../context"
)

func c403() context.BizError {
	return context.BizError{
		Code:     403,
		Msg:      "auth-fail",
		Body:     nil,
		Callback: nil,
	}
}
