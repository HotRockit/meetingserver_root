package router

import (
	"github.com/gin-gonic/gin"
	"meeting_root_server/controller"
)

//AppRouter is a struct to handler router
type AppRouter struct {

}

var appController = new(controller.AppController)

func (appRouter *AppRouter) UserRouter(engine *gin.Engine){
	appController.UserController(engine)
}

func (appRouter *AppRouter) MeetingRouter(engine *gin.Engine) {
	appController.MeetingController(engine)
}
