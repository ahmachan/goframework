package api

import (
	"goframework/Domain"
        "encoding/json"
        "fmt"
        "os"
)

type ApiResData struct {
     object     interface{} `json:"object"`
     list       interface{} `json:"list"`
     extend     interface{} `json:"extend"`
}

type ApiResult struct {
     code     string `json:"code"`
     message  string `json:"message"`
     data     interface{} `json:"data"`
}


func GetHello(id int) interface{} {
     userList,err := Domain.GetHello(id)
     apiCode := "1"
     apiMsg := "OK"
     var result ApiResult
     if (err!=nil){
        fmt.Println(err)
        apiCode = "2001"
        apiMsg = "failed"
     }

     resData := ApiResData{
        object:userList,
        list:nil,
        extend:nil}//注意，如果没有逗号，则}不能另起新行，否则会报错  
     result.code = apiCode
     result.message= apiMsg
     result.data= resData
     fmt.Println("resData: ",resData)
     /* 
     result:=map[string]string{
               "code": apiCode,
               "message":apiMsg}
     */

    type ColorGroup struct {  
        ID     int  
        Name   string  
        Colors []string  
    }  
    group := ColorGroup{  
        ID:     1,  
        Name:   "Reds",  
        Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},  
    }  

     fmt.Println("result: ",group)
     /**/
     ret, err := json.Marshal(result)//转换成JSON返回的是[]byte
     //var ret []ApiResult
     //err:=json.Unmarshal(result, &ret) 
     if err != nil {
         fmt.Println("Error: ", err)
     } 
     fmt.Println(string(ret))
     os.Stdout.Write(ret) 
     
    return string(ret)
}

