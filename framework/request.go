//	请求解析,获取get,post,json参数,签名加密,链式操作,并且参数验证
//  依赖情况:
//          "github.com/astaxie/beego/validation" 基于beego的拦截器(已经集成)
//          "github.com/labstack/echo" 依赖于echo

package framework

import (
	"strconv"
        "goframework/ace"
        "goframework/framework/validation"
	//"github.com/labstack/echo"
	//"github.com/wenzhenxi/phalgo/validation"
	"errors"
	//"crypto/md5"
	//"encoding/hex"
	//"encoding/base64"
	//"regexp"
	"fmt"
)

type Request struct {
	Context    ace.C
	params     *param
	//Jsonparam  *Jsonparam
	valid      validation.Validation
	//Json       *Js
	//Encryption bool
	//Des        Des
	jsonTag    bool
	Debug      bool
}


type param struct {
	key string
	val string
	min int
	max int
}

// 初始化request
func NewRequest(c ace.C) *Request {

	R := new(Request)
	R.Context = c
	//增加debug参数的匹配
	if R.Param("__debug__").SetDefault("").GetString() == "" {
		R.Debug = false
	} else {
		R.Debug = true
	}
	return R
}

// 清理参数
func (this *Request)Clean() {

	this.params = new(param)
	//this.Jsonparam = new(Jsonparam)
}

// 返回报错信息
func (this *Request)GetError() error {

	if this.valid.HasErrors() {
		for _, v := range this.valid.Errors {
			return errors.New(v.Message + v.Key)
		}
	}

	return nil
}




//--------------------------------------------------------获取参数-------------------------------------


// 获取Get参数
func (this *Request)GetParam(key string) *Request {

	this.Clean()
	str := this.Context.QueryParam(key)
	this.params.val = str
	this.params.key = key
	this.jsonTag = false

	return this
}

// 获取post参数
func (this *Request)PostParam(key string) *Request {

	this.Clean()
	str := this.Context.FormValue(key)
	this.params.val = str
	this.params.key = key
	this.jsonTag = false

	return this
}

// 获取请求参数顺序get->post
func (this *Request)Param(key string) *Request {

	var str string
	this.Clean()
	str = this.Context.QueryParam(key)

	if str == "" {
		str = this.Context.FormValue(key)
	}

	this.params.val = str
	this.params.key = key

	return this
}

func (this *Request)SetDefault(val string) *Request {
	if this.jsonTag == true {
		//defJson := fmt.Sprintf(`{"index":"%s"}`, val)
		//this.Jsonparam.val = *Json(defJson).Get("index")
		/*		fmt.Println(defJson)
				fmt.Println(this.Jsonparam.val.Tostring())*/
	} else {
		this.params.val = val
	}
	return this
}

//----------------------------------------------------过滤验证------------------------------------

// GET,POST或JSON参数是否必须
func (this *Request)Require(b bool) *Request {

	this.valid.Required(this.getParamVal(), this.getParamKey()).Message("缺少必要参数,参数名称:")
	return this
}

// 设置参数最大值
func (this *Request)Max(i int) *Request {

	this.params.max = i
	return this
}

//设置参数最小值
func (this *Request)Min(i int) *Request {

	this.params.min = i
	return this
}


//--------------------------------------------GET,POST获取参数------------------------------------

// 获取参数 string类型 适用于GET或POST参数
func (this *Request)GetString() string {

	var str string

	str = this.getParamVal()


	return str
}

// 获取参数 int类型 适用于GET或POST参数
func (this *Request)GetInt() int {
	var (
		i int
		err error
	)

	if this.getParamVal() == "" {
		i = 0
	} else {
		i, err = strconv.Atoi(this.getParamVal())
		if err != nil {
			this.valid.SetError(this.getParamKey(), "参数异常!参数不是int类型,参数名称:")
		}
	}

	return i
}

// 获取并且验证参数 float64类型 适用于GET或POST参数
func (this *Request)GetFloat() float64 {

	var (
		i float64
		err error
	)

	if this.getParamVal() == "" {
		i = 0
	} else {
		i, err = strconv.ParseFloat(this.getParamVal(), 64)
		if err != nil {
			this.valid.SetError(this.getParamKey(), "此参数无法转换为float64类型,参数名称:")
		}
	}

	return i
}

// 返回解析参数的Val
func (this *Request)getParamVal() string {

	if this.jsonTag {
		return "" //this.Jsonparam.val.Tostring()
	} else {
		return this.params.val
	}
}

// 反回解析参数的Key
func (this *Request)getParamKey() string {
	if this.jsonTag {
		return "" //this.Jsonparam.key
	} else {
		return this.params.key
	}
}


// 捕获panic异样防止程序终止 并且记录到日志
func (this *Request)ErrorLogRecover() {

	if err := recover(); err != nil {
                /*
		this.Context.Response().Write([]byte("系统错误!具体原因:" + TurnString(err)))
		LogError(err, map[string]interface{}{
			"URL.Path":this.Context.Request().URL().Path(),
			"QueryParams":this.Context.QueryParams(),
		})
                */
	}
}
