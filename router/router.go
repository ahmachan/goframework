package routers
import (
	"goframework/framework"
	"goframework/ace"
	"goframework/api"
	//"goframework/Domain"
        "fmt"
)


func NewRouter() *ace.Ace {
        aceInstance := ace.New()
	g := aceInstance.Group("/user", func(context *ace.C) {
        	fmt.Println("before route func")
        	context.Next()
        })

        g.GET("/", func(context *ace.C) {
                apiRes := map[string]string{"TEST": "home"}     
                //apiRes := api.GetHello(2)
                
                fmt.Println(apiRes)
                //if (err!=nil){
                //   fmt.Println(err)
                //}        	
                //c.JSON(status int, v interface{})
        	context.JSON(200, apiRes)
        })

        // /user/:name
        g.GET("/:uid", func(context *ace.C) {
                /*
                userIdstr := context.Param("uid") //return string	
                userId := framework.TurnInt(userIdstr)
                apiRes := api.GetHello(userId)
                //apiRes := map[string]string{"TEST": userIdstr}
                fmt.Println("id: "+userIdstr)
                fmt.Println("userId: ",userId)
       
                //c.JSON(status int, v interface{})
        	context.JSON(200, apiRes)
                */


                request := framework.NewRequest(context)
		response := framework.NewResponse(context)
		defer request.ErrorLogRecover()

		uid := request.GetParam("uid").GetInt()

		//参数过滤error处理
		if err := request.GetError(); err != nil {
			return response.RetError(err, -1)
		}

		//user, err := this.Domain.User.GetUserInfo(id)
		//if err != nil {
		//	return Response.RetError(err, 400)
		//}
                apiRes := api.GetHello(uid)

		return response.RetSuccess(apiRes)
        })

        // /user/:name
        g.POST("/:uid", func(context *ace.C) {
        	context.JSON(200, map[string]string{"TEST": "POST METHOD"})
        })

        g.GET("/:uid/tags",func(context *ace.C){
               userId := context.Param("uid")
               context.JSON(200, map[string]string{"TAGS": "block:"+userId})
        })

    return aceInstance

}
