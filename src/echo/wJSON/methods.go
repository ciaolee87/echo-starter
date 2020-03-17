package wJSON

func (j *JSON) FillDefault() {
	if j.Code == nil {
		defCode := 200
		j.Code = &defCode
	}

	if j.Msg == nil {
		defMsg := "success"
		j.Msg = &defMsg
	}
}

func (j *JSON) FillErrorDefault() {
	if j.Code == nil {
		defCode := 500
		j.Code = &defCode
	}

	if j.Msg == nil {
		defMsg := "server_error"
		j.Msg = &defMsg
	}
}
