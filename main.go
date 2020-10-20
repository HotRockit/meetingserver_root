package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"meeting_root_server/controller"
	"meeting_root_server/router"
	"meeting_root_server/util"
	"net/http"
	"strings"
)

func AddFunc(a int,b int) int{
	return a+b
}
func SubFunc(a int,b int) int{
	return a-b
}

func MultiFunc(a int,b int) int{
	return a*b
}

func DivFunc(a int,b int) int{
	return a/b
}

func main() {

	//读取配制文件
	cfg, err := util.ParseConfig("./config/app.json")
	if err != nil {
		log.Print(err.Error())
		panic(err)
		return
	}

	//初始化数据库
	_,err = util.InitDataBase(cfg)
	if err != nil {
		log.Print(err.Error())
		panic(err)
		return
	}

	app := gin.Default()

	//登陆状态验证
	//app.Use(AuthMiddleWare())
	
	//添加自定义模板函数
	app.SetFuncMap(template.FuncMap{
		"sub" : SubFunc,
		"add" : AddFunc,
		"mul" : MultiFunc,
		"div" : DivFunc,
	})

	app.Static("/user/static","resource/static")
	app.Static("/meeting/static","resource/static")
	app.Static("/static","resource/static")
	app.LoadHTMLGlob("resource/templates/*")
	app.GET("/meeting_root_server", func(context *gin.Context) {
		context.HTML(http.StatusOK,"login.html",nil)
	})
	appRouter := new(router.AppRouter)
	appRouter.UserRouter(app)
	appRouter.MeetingRouter(app)
	app.Run(cfg.AppHost + ":" + cfg.AppPort)


}

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		//这种方式在请求静态资源文件时也会验证就会出现问题,所以就判断请求路径里面有没有static，如果有的话说明请求的是静态资源文件就直接放行
		if url := c.Request.URL.String(); url =="/meeting_root_server" || strings.Contains(url,"static"){
			c.Next()
			return
		}else if controller.User == nil {
			c.Redirect(http.StatusPermanentRedirect,"/meeting_root_server")
			c.Next()
			return
		}else{
			c.Next()
			return
		}
	}
}


