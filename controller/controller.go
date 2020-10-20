package controller

import (
	"github.com/gin-gonic/gin"
	"meeting_root_server/model"
	"meeting_root_server/service"
	"meeting_root_server/util"
	"net/http"
	"strconv"
	"strings"
)

//MeetingController is a controller to handle
type AppController struct {

}

var User *model.SuperUser = nil

var appService = new(service.AppService)

const  pageNum = 10

var realTimeList = [31]string{"08:00","08:30","09:00","09:30","10:00","10:30","11:00","11:30","12:00","12:30","13:00","13:30","14:00","14:30","15:00","15:30",
"16:00","16:30","17:00","17:30","18:00","18:30","19:00","19:30","20:00","20:30","21:00","21:30","22:00","22:30","23:00"}

func (meetingController *AppController) UserController(engine *gin.Engine) {
	userGroup := engine.Group("/user")
	{
		userGroup.GET("/index",meetingController.IndexHandler)
		userGroup.POST("/login",meetingController.LoginHandler )
		userGroup.POST("/updatePassword",meetingController.UpdatePasswordHandler)
		userGroup.POST("/deleteUser",meetingController.DeleteUserHandler)
		userGroup.GET("/userList/:page/:username",meetingController.UserListHandler)
		userGroup.POST("/freezeUser",meetingController.FreezeUserHandler)
		userGroup.GET("/userEdit",meetingController.UserEditHandler)
		userGroup.POST("deleteSuperUser",meetingController.DeleteSuperUserHandler)
	}
}

func (meetingController *AppController) UserListHandler(c *gin.Context) {
	currentPage, _ := strconv.Atoi(c.Param("page"))
	username := c.Param("username")
	//这里处理分页信息,每页显示十条
	users := appService.QueryAllUserService(currentPage,pageNum,username)   //根据传进来的页数获取相应的十条记录
	//获取总的用户数
	allUser := appService.QueryAllUserService(1,0,username)
	userCount := len(allUser)
	var totalPage int
	if userCount % pageNum == 0 {
		totalPage = userCount / pageNum
	}else {
		totalPage = userCount /pageNum +1
	}
	if totalPage == 0{
		totalPage = totalPage + 1
	}
	//由于go的template只支持ranger循环，不支持for循环(可以使用第三方的库https://github.com/flosch/pongo2，这个支持for循环)
	//这里我们把每一个页码分开放到一个slice里面，在前端去遍历
	var pageSlice []int
	for i := 1;i <= totalPage; i++ {
		pageSlice = append(pageSlice, i)
	}
	 c.HTML(http.StatusOK,"userlist.html",gin.H{
	 	"users": users,
	 	"user" : User,
	 	"userCount": userCount,
	 	"totalPage" : totalPage,
	 	"currentPage" : currentPage,
	 	"pageSlice" : pageSlice,
	 	"pageNum" : pageNum,
	 	"username" : username,
	 })
}

func (meetingController *AppController) IndexHandler(c *gin.Context) {
	//获取会议室总数
	meetingRoomCount := appService.QueryMeetingRoomCountService()
	//获取用户总数
	userCount := appService.QueryUserCountService()
	//获取会议总数
	meetingCount := appService.QueryMeetingCountService()
	//获取最新的会议消息，系统取最近的十个
	latestMeetings := appService.QueryLatestMeetingsService()
	//获取热点会议室，系统取预约次数前十的会议室
	hotMeetingRooms := appService.QueryHotMeetingRoomsService()
	c.HTML(http.StatusOK,"index.html",gin.H{
		"user" : User,
		"meetingRoomCount" : meetingRoomCount,
		"userCount" : userCount,
		"meetingCount" : meetingCount,
		"latestMeetings" : latestMeetings,
		"hotMeetingRooms" : hotMeetingRooms,
	})
}

func (meetingController *AppController) DeleteUserHandler(c *gin.Context) {
	username := c.PostForm("username")
	util.JsonResult(c,appService.DeleteUserService(username))
}

func (meetingController *AppController) UpdatePasswordHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	util.JsonResult(c,appService.UpdatePasswordService(username,password))
}

func (meetingController *AppController) LoginHandler(c *gin.Context)  {
	username := c.PostForm("username")
	password := c.PostForm("password")
	result,user := appService.LoginService(username,password)
	util.JsonResult(c,result)
	User = user
}


func (meetingController *AppController) MeetingController(engine *gin.Engine) {
	meetingGroup := engine.Group("/meeting")
	{
		meetingGroup.GET("/queryMeetingRoom",meetingController.QueryMeetingRoomHandler)
		//meetingGroup.GET("/queryMeeting",meetingController.QueryMeetingHandler)
		meetingGroup.GET("/roomList/:page/:room",meetingController.QueryMeetingRoomHandler)
		meetingGroup.POST("/updateRoomName",meetingController.UpdateRoomNameHandler)
		meetingGroup.POST("/deleteMeetingRoom",meetingController.DeleteMeetingRoomHandler)
		meetingGroup.POST("/addMeetingRoom",meetingController.AddMeetingRoomHandler)
		meetingGroup.GET("/meetingList/:page/:roomName/:username/:borrowTime",meetingController.QueryMeetingHandler)
		meetingGroup.POST("/deleteMeeting/:id",meetingController.DeleteMeetingHandler)
		meetingGroup.GET("/meetingReport",meetingController.ShowMeetingReportHandler)
		meetingGroup.POST("/deleteMeetings",meetingController.DeleteMeetingsHandler)
	}
}

func (meetingController *AppController) DeleteMeetingHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	util.JsonResult(c,appService.DeleteMeetingService(id))
}


func (meetingController *AppController) QueryMeetingRoomHandler(c *gin.Context) {
	currentPage, _ := strconv.Atoi(c.Param("page"))
	room := c.Param("room")
	//这里处理分页信息,每页显示十条
	rooms := appService.QueryAllMeetingRoomService(currentPage,pageNum,room)   //根据传进来的页数获取相应的十条记录
	//获取总的会议室数
	allRoom := appService.QueryAllMeetingRoomService(1,0,room)
	roomCount := len(allRoom)
	var totalPage int
	if roomCount % pageNum == 0 {
		totalPage = roomCount / pageNum
	}else {
		totalPage = roomCount /pageNum +1
	}
	if totalPage == 0{
		totalPage = totalPage + 1
	}
	//由于go的template只支持ranger循环，不支持for循环(可以使用第三方的库https://github.com/flosch/pongo2，这个支持for循环)
	//这里我们把每一个页码分开放到一个slice里面，在前端去遍历
	var pageSlice []int
	for i := 1;i <= totalPage; i++ {
		pageSlice = append(pageSlice, i)
	}
	c.HTML(http.StatusOK,"roomlist.html",gin.H{
		"rooms": rooms,
		"user" : User,
		"roomCount": roomCount,
		"totalPage" : totalPage,
		"currentPage" : currentPage,
		"pageSlice" : pageSlice,
		"pageNum" : pageNum,
		"room" : room,
	})
}

func (meetingController *AppController) QueryMeetingHandler(c *gin.Context) {
	currentPage, _ := strconv.Atoi(c.Param("page"))
	roomName := c.Param("roomName")
	username := c.Param("username")
	borrowTime := c.Param("borrowTime")
	//这里处理分页信息,每页显示十条
	meetings := appService.QueryAllMeetingService(currentPage,pageNum,roomName,username, borrowTime) //根据传进来的页数获取相应的十条记录
	//获取总的用户数
	allMeetings := appService.QueryAllMeetingService(1,0,roomName,username, borrowTime)
	meetingCount := len(allMeetings)
	var totalPage int
	if meetingCount % pageNum == 0 {
		totalPage = meetingCount / pageNum
	}else {
		totalPage = meetingCount /pageNum +1
	}
	if totalPage == 0{
		totalPage = totalPage + 1
	}
	//由于go的template只支持ranger循环，不支持for循环(可以使用第三方的库https://github.com/flosch/pongo2，这个支持for循环)
	//这里我们把每一个页码分开放到一个slice里面，在前端去遍历
	var pageSlice []int
	for i := 1;i <= totalPage; i++ {
		pageSlice = append(pageSlice, i)
	}
	c.HTML(http.StatusOK,"meetinglist.html",gin.H{
		"meetings": meetings,
		"user" : User,
		"meetingCount": meetingCount,
		"totalPage" : totalPage,
		"currentPage" : currentPage,
		"pageSlice" : pageSlice,
		"rooms" : appService.QueryAllMeetingRoomService(1,0,"no"),
		"realTimeList" : realTimeList,
		"roomName" : roomName,
		"username" : username,
		"borrowTime" : borrowTime,
		"pageNum" : pageNum,
	})
}

func (meetingController *AppController) FreezeUserHandler(c *gin.Context) {
	username := c.PostForm("username")
	util.JsonResult(c,appService.FreezeUserService(username))
}

func (meetingController *AppController) UserEditHandler(c *gin.Context) {
	c.HTML(http.StatusOK,"useredit.html",gin.H{
		"user":User,
	})
}

func (meetingController *AppController) DeleteSuperUserHandler(c *gin.Context) {
	username := c.PostForm("username")
	util.JsonResult(c,appService.DeleteSuperUserService(username))
}

func (meetingController *AppController) UpdateRoomNameHandler(c *gin.Context) {
	roomName := c.PostForm("roomName")
	newRoomName := c.PostForm("newRoomName")
	util.JsonResult(c,appService.UpdateRoomNameService(roomName,newRoomName))
}

func (meetingController *AppController) DeleteMeetingRoomHandler(c *gin.Context) {
	roomName := c.PostForm("roomName")
	util.JsonResult(c,appService.DeleteMeetingRoomService(roomName))
}

func (meetingController *AppController) AddMeetingRoomHandler(c *gin.Context) {
	roomName := c.PostForm("roomName")
	util.JsonResult(c,appService.AddMeetingRoomService(roomName))
}

func (meetingController *AppController) ShowMeetingReportHandler(c *gin.Context) {
	//获取每一个会议室的会议数
	result := appService.ShowMeetingReportService()
	c.HTML(http.StatusOK,"meetingsreport.html",gin.H{
		"user" : User,
		"result" : result,
	})
}

func (meetingController *AppController) DeleteMeetingsHandler(c *gin.Context) {

	result := c.PostForm("ids")
	list := strings.Split(result,",")
	for i:=0;i<len(list);i++{
		if list[i]!="undefined" {
			id,_ := strconv.Atoi(list[i])
			res := appService.DeleteMeetingService(id)
			if res == 0{
				util.JsonResult(c,0)
				return
			}
		}
	}
	util.JsonResult(c,1)
}
