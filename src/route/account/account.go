package account

import (
	"../../wrapper/context"
)

func Router(e context.BizEcho) {
	e.BizGET("/user", func(c context.BizContext) error {
		return c.BizSendJson(&context.BizJSON{Body: 123})
	})
}
