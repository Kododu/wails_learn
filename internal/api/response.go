package api

import (
	"changeme/internal/constant"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func ReplayBasicError(err error) Response {
	return Response{
		Code:    5001,
		Message: err.Error(),
		Data:    constant.EmptyJson,
	}
}

func ReplaySuccess(data any) Response {
	return Response{
		Code:    1000,
		Message: "Success",
		Data:    data,
	}
}

func ReplaySuccessEmpty() Response {
	return Response{
		Code:    1000,
		Message: "Success",
		Data:    constant.EmptyJson,
	}
}

type ResponseDataBuilder struct {
	inner map[string]interface{}
}

// BuildResponse 返回一个响应构造器,用于设置 Response 的data
func BuildResponse(key string, data interface{}) *ResponseDataBuilder {
	builder := ResponseDataBuilder{
		inner: make(map[string]interface{}),
	}
	builder.inner[key] = data
	return &builder
}

// Put 添加更多的内容
func (builder *ResponseDataBuilder) Put(key string, data interface{}) *ResponseDataBuilder {
	builder.inner[key] = data
	return builder
}

// Replay 返回的响应必定是1000
func (builder *ResponseDataBuilder) Replay() Response {
	return Response{
		Code:    1000,
		Message: "Success",
		Data:    builder.inner,
	}
}
