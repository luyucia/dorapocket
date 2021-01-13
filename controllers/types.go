package controllers

type Response struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
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

