package routers
import (
	"phalgo-sample/framework"
	"phalgo-sample/ace"
	"phalgo-sample/api"
	//"phalgo-sample/Domain"
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
                userIdstr := context.Param("uid") //return string	
                userId := framework.TurnInt(userIdstr)
                apiRes := api.GetHello(userId)
                //apiRes := map[string]string{"TEST": userIdstr}
                fmt.Println("id: "+userIdstr)
                fmt.Println("userId: ",userId)
                //if (err!=nil){
                //   fmt.Println(err)
                //}
        	//context.JSON(200, map[string]string{"TEST": apiRes})
                //c.JSON(status int, v interface{})
        	context.JSON(200, apiRes)
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
