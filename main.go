package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
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
	store := cookie.NewStore([]byte("wdy970603"))
	store.Options(sessions.Options{Path: "/", MaxAge: 60*5})   //设置过期时间(秒数)
	app.Use(sessions.Sessions("mySession",store))

	//登陆状态验证
	app.Use(AuthMiddleWare())
	
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
		session := sessions.Default(c)
		user := session.Get("user")
		if url := c.Request.URL.String(); url =="/meeting_root_server" || url == "/user/login" || strings.Contains(url,"static") || user == true{
			c.Next()
		}else {
			//之前这里设置的是永久重定向，运行时发现cookie里面有用户信息，但是点击之前重定向的页面，效果还是重定向之后的页面
			//例如，点击了用户列表界面，但是cookie里面没有用户信息，重定向到了登陆界面，登录之后有了用户信息，但是点击用户列表界面，还是跳转到了登陆界面
			c.Redirect(http.StatusTemporaryRedirect,"/meeting_root_server")
			c.Abort()
			return
		}
	}
}


