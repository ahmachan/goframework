package routers
import (
	//"goframework/framework"
	"goframework/ace"
	"goframework/api"
	//"goframework/utils"
        "fmt"
)


func NewRouter() *ace.Ace {
        aceInstance := ace.New()
	g := aceInstance.Group("/user", func(context *ace.C) {
        	fmt.Println("before route func")
        	context.Next()
        })

        g.GET("/", func(context *ace.C) {
 
                fmt.Println("context:",context,"\n")                
	
                limit := context.MustQueryInt("limit",0)
		fmt.Println("limit:",limit,"\n")
                apiRes := api.GetUserList()

                context.JSON(200, apiRes)
        })

        // /user/:uid
        g.GET("/:uid/details", func(context *ace.C) {
                fmt.Println("context:",context,"\n")                
		//uid2 := context.MustQueryInt("uid",0) //at url query   ,or at form data
                uid := context.MustPathInt("uid",0) //at path link
                
		fmt.Println("userId:",uid,"\n")
                limit := context.MustQueryInt("limit",0)
		fmt.Println("limit:",limit,"\n")
                apiRes := api.GetHello(uid)

                context.JSON(200, apiRes)
        })

        // /user/:name
        g.POST("/:uid", func(context *ace.C) {
        	context.JSON(200, map[string]string{"TEST": "POST METHOD"})
        })

        // user/12345/tags
        g.GET("/:uid/tags",func(context *ace.C){
               userId := context.Param("uid")
               context.JSON(200, map[string]string{"TAGS": "block:"+userId})
        })

    return aceInstance

}
