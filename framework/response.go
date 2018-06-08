//  返回json参数,默认结构code,msg,data
//  依赖情况:
//          "github.com/labstack/echo" 必须基于echo路由

package framework

import (
	"goframework/ace"
	"net/http"
)

type Response struct {
	Context   ace.C
	parameter *RetParameter
}

type RetParameter struct {
	Code int         `json:"code";xml:"code"`
	Msg  string      `json:"msg";xml:"msg"`
	Data interface{} `json:"data";xml:"data"`
}

const DefaultCode = 1

var HttpStatus = http.StatusOK

// 初始化Response
func NewResponse(c ace.C) *Response {

	R := new(Response)
	R.Context = c
	R.parameter = new(RetParameter)
	R.parameter.Data = nil
	return R
}

// 设置返回的Status值默认http.StatusOK
func (this *Response)SetStatus(i int) {
	HttpStatus = i
}

func (this *Response)SetMsg(s string) {
	this.parameter.Msg = s
}

func (this *Response)SetData(d interface{}) {
	this.parameter.Data = d
}

// 返回自定自定义的消息格式
func (this *Response)RetCustomize(code int, d interface{}, msg string) error {

	this.parameter.Code = code
	this.parameter.Data = d
	this.parameter.Msg = msg

	return this.Ret(this.parameter)
}

// 返回成功的结果 默认code为1
func (this *Response)RetSuccess(d interface{}) error {

	this.parameter.Code = DefaultCode
	this.parameter.Data = d

	return this.Ret(this.parameter)
}

// 返回失败结果
func (this *Response)RetError(err error, code int) error {

	this.parameter.Code = code
	this.parameter.Msg = err.Error()

	return this.Ret(this.parameter)
}

// 返回结果
func (this *Response)Ret(par interface{}) error {

	//switch RetType {
	//case 2:
	//	return this.Context.XML(HttpStatus, par)
	//default:
	//	return this.Context.JSON(HttpStatus, par)
	//}
        return this.Context.JSON(HttpStatus, par)
}

// 输出返回结果
func (this *Response)Write(b []byte) {
        this.Context.Response().Writer.Write(b)
	//_, e := this.Context.Response().Write(b)
        
	//if e != nil {
	//	print(e.Error())
	//}
}



