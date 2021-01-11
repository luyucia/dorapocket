package controllers

type Response struct {
	Code int
	Msg string
	Data interface{}
}

func NewSuccessResponse()  Response {
	return Response{0,"ok",nil}
}

func NewDataResponse(data interface{})  Response {
	return Response{0,"ok",data}
}

func NewErrorResponse(msg string)  Response {
	return Response{500,msg,nil}
}

