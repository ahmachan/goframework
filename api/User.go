package api

import (
    "goframework/utils"
	"goframework/Domain"
    //"encoding/json"
    "fmt"
    //"os"
)

/*
type ApiResData struct {
     object     interface{} `json:"object"`
     list       interface{} `json:"list"`
     extend     interface{} `json:"extend"`
}
*/

type ApiResult struct {
    Code     string `json:"code"`
    Message  string `json:"message"`
    Data     map[string]interface{} `json:"data"`
}

func GetHello(id int) map[string]interface{} {
    userList,err := Domain.GetHello(id)
    //userList,err := Domain.GetUserList()
    apiCode := "1"
    apiMsg := "OK"
    var result ApiResult
    if (err!=nil){
        fmt.Println(err)
        apiCode = "2001"
        apiMsg = "failed"
    }

    //注意，最后一项"extend:nil,"如果没有逗号，则}不能另起新行，否则会报错  
    /*
    resData := ApiResData{
        object:userList,
        list:nil,
        extend:nil,
    }
    */
        
 
    //result.Code = apiCode
    //result.Message= apiMsg
    //result.Data= utils.StructToMap2(userList)

    result = ApiResult{
        Code:apiCode,
        Message:apiMsg,
        Data:utils.StructToMap2(userList),
    }

    resMap:=utils.StructToMap2(result)
    fmt.Println("resData: ",resMap)

        
    return resMap
}


func GetUserList() map[string]interface{} {
    userList,err := Domain.GetUserList()
    
    apiCode := "1"
    apiMsg := "OK"
    var result ApiResult
    if (err!=nil){
        fmt.Println(err)
        apiCode = "2001"
        apiMsg = "failed"
    }
    fmt.Println("userList: ",userList)
    result = ApiResult{
        Code:apiCode,
        Message:apiMsg,
        Data:utils.StructToMap2(userList),
    }

    resMap:=utils.StructToMap2(result)
    fmt.Println("resData: ",resMap)

        
    return resMap
}
