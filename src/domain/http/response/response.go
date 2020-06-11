package response

type Response struct {
	Code     int64       `json:"code"`
	Messages []string    `json:"message"`
	Error    string      `json:"error"`
	Data     interface{} `json:"data"`
}

func (r Response) SetCote(code int64) Response           { r.Code = code; return r }
func (r Response) SetMessages(message []string) Response { r.Messages = message; return r }
func (r Response) SetError(error string) Response        { r.Error = error; return r }
func (r Response) SetData(data interface{}) Response     { r.Data = data; return r }
func (r Response) AddMessages(message string) Response {
	r.SetMessages(append(r.Messages, message))
	return r
}

func (r Response) Make(code int64, messages []string, error string, data interface{}) Response {
	return r.SetCote(code).SetMessages(messages).SetError(error).SetData(data)
}

func (r Response) Success(data interface{}, messages ...string) Response {
	return r.Make(1, messages, "", data)
}

func (r Response) BadRequest(err error, messages ...string) Response {
	return r.Make(-1, messages, err.Error(), nil)
}

func (r Response) InValid(err error, messages ...string) Response {
	return r.Make(-2, messages, err.Error(), nil)
}

func (r Response) NotFound(err error, messages ...string) Response {
	return r.Make(-3, messages, err.Error(), nil)
}

func (r Response) TryAgain(err error, messages ...string) Response {
	return r.Make(-4, messages, err.Error(), nil)
}

func (r Response) InternalError(err error, messages ...string) Response {
	return r.Make(-5, messages, err.Error(), nil)
}
