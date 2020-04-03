package wJSON

type JSON struct {
	Code *int         `json:"code"`
	Msg  *string      `json:"msg"`
	Body *interface{} `json:"body"`
}
