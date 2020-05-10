package bizEcho

type BizJSON struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Body *interface{} `json:"body,omitempty"`
}

func NewJSON() *BizJSON {
	json := BizJSON{
		Code: 200,
		Msg:  "success",
		Body: nil,
	}
	return &json
}

func NewErrorJSON() *BizJSON {
	json := BizJSON{
		Code: 500,
		Msg:  "Server Error",
		Body: nil,
	}

	return &json
}
