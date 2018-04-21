package handlers

// ResponseError as
type ResponseError struct {
	Error   bool
	Message string
}

// Error error
type Error struct {
	Error   bool
	Message string
}

// Response as
type Response struct {
	Error Error
	Data  interface{}
}

// Data data response
type Data struct {
	Data interface{}
}

func responseSuccess(message string, i interface{}) Response {
	return Response{Error{false, message}, i}
}

func responseError(message string) Response {
	return Response{Error{true, message}, nil}
}
