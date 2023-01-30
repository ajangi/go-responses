package responses

import (
	"net/http"
)

type Result string

const (
	Error   Result = "ERROR"
	Success Result = "SUCCESS"
)

type Message []string

type Response struct {
	Messages   map[string]Message `json:"messages"`
	StatusCode int                `json:"status_code"`
	Result     Result             `json:"result"`
	Data       interface{}        `json:"data"`
}

func NotFoundError(resource string) Response {
	return Response{
		Messages:   map[string]Message{resource: {resource + " not found!"}},
		StatusCode: 404,
		Result:     Error,
		Data:       nil,
	}
}

func BadRequestError(error string) Response {
	return Response{
		Messages:   map[string]Message{"api": {error}},
		StatusCode: 400,
		Result:     Error,
		Data:       nil,
	}
}

func ErrorResponse(error map[string]Message) Response {
	return Response{
		Messages:   error,
		StatusCode: 400,
		Result:     Error,
		Data:       nil,
	}
}

func InternalServerErrorResponse() Response {
	return Response{
		Messages:   map[string]Message{"server": {"Internal server error!"}},
		StatusCode: 500,
		Result:     Error,
		Data:       nil,
	}
}

func MethodNotAllowedErrorResponse() Response {
	return Response{
		Messages:   map[string]Message{"method": {"Method not allowed!"}},
		StatusCode: 405,
		Result:     Error,
		Data:       nil,
	}
}

func UnAuthorizedError() Response {
	return Response{
		Messages:   map[string]Message{"auth": {"UnAuthorized!"}},
		StatusCode: http.StatusUnauthorized,
		Result:     Error,
		Data:       nil,
	}
}

func SuccessResponse() Response {
	return Response{
		Messages:   map[string]Message{},
		StatusCode: 200,
		Result:     Success,
		Data:       nil,
	}
}

func SuccessResponseWithData(data interface{}) Response {
	return Response{
		Messages:   map[string]Message{},
		StatusCode: 200,
		Result:     Success,
		Data:       data,
	}
}
