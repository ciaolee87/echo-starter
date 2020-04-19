package bizEcho

type BizJSON struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Body *interface{} `json:"body"`
}

func NewJSON() *BizJSON {
	json := BizJSON{
		Code: 200,
		Msg:  "success",
		Body: nil,
	}
	return &json
}
