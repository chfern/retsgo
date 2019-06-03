package core

type HttpException struct {
	Code  int
	Error string
}

func NewHttpException(code int, err string) *HttpException {
	return &HttpException{
		Code:  code,
		Error: err,
	}
}
