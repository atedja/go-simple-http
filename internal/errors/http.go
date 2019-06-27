package errors

func BadParameter(msg string) *Error {
	return &Error{
		Message:    msg,
		Type:       "ERR_BAD_PARAMETER",
		StatusCode: 400,
	}
}

func Unauthorized(msg string) *Error {
	return &Error{
		Message:    msg,
		Type:       "ERR_UNAUTHORIZED",
		StatusCode: 401,
	}
}

func Forbidden(msg string) *Error {
	return &Error{
		Message:    msg,
		Type:       "ERR_UNAUTHORIZED",
		StatusCode: 403,
	}
}

func NotFound(msg string) *Error {
	return &Error{
		Message:    msg,
		Type:       "ERR_NOT_FOUND",
		StatusCode: 404,
	}
}

func ServerError(msg string) *Error {
	return &Error{
		Message:    msg,
		Type:       "ERR_INTERNAL_SERVER",
		StatusCode: 500,
	}
}

func BadGateway(msg string) *Error {
	return &Error{
		Message:    msg,
		Type:       "ERR_BAD_GATEWAY",
		StatusCode: 502,
	}
}

func GatewayTimeout(msg string) *Error {
	return &Error{
		Message:    msg,
		Type:       "ERR_GATEWAY_TIMEOUT",
		StatusCode: 504,
	}
}
