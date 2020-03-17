package wJSON

func DefaultJson(body interface{}) *JSON {
	return Make(200, "success", &body)
}

func ErrorJson(body interface{}) *JSON {
	return Make(500, "server_error", &body)
}

func Make(code int, msg string, body *interface{}) *JSON {
	json := JSON{Code: &code, Msg: &msg, Body: body}
	return &json
}
