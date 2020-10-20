package service

import (
	"meeting_root_server/dao"
	"meeting_root_server/model"
	"meeting_root_server/util"
)

type AppService struct {

}


func (appService *AppService)LoginService(username string,password string) (uint,*model.SuperUser) {
	//下面这个变量虽然所有的方法里面都有而且一样，但是不能把它提到全局位置，然后共用。因为全局变量在加载文件的时候就初始化了的，但是
	//如果前面的方法没执行的话util.GlobalGormEngine是没有初始化的,是nil
	appDao := dao.AppDao{GormEngine: util.GlobalGormEngine}
	return appDao.LoginDao(username,password)
}


func (appService *AppService) QueryMeetingRoomService() []model.MeetingRoom {
	appDao := dao.AppDao{GormEngine: util.GlobalGormEngine}
	return appDao.QueryMeetingRoomDao()
}


func (appService *AppService) UpdatePasswordService(username string, password string) uint {
	appDao := dao.AppDao{GormEngine: util.GlobalGormEngine}
	return appDao.UpdatePasswordDao(username,password)
}

func (appService *AppService) DeleteUserService(username string) uint {
	appDao := dao.AppDao{GormEngine: util.GlobalGormEngine}
	return appDao.DeleteUserDao(username)
}

func (appService *AppService) QueryMeetingRoomCountService() uint {
	appDao := dao.AppDao{GormEngine: util.GlobalGormEngine}
	return appDao.QueryMeetingRoomCountDao()
}

func (appService *AppService) QueryUserCountService() uint {
	appDao := dao.AppDao{GormEngine: util.GlobalGormEngine}
	return appDao.QueryUserCountDao()
}

func (appService *AppService) QueryMeetingCountService() uint {
	appDao := dao.AppDao{GormEngine: util.GlobalGormEngine}
	return appDao.QueryMeetingCountDao()
}

func (appService *AppService) QueryLatestMeetingsService() []model.Meeting {
	appDao := dao.AppDao{GormEngine: util.GlobalGormEngine}
	return appDao.QueryLatestMeetingsDao()
}

func (appService *AppService) QueryHotMeetingRoomsService() []model.MeetingCount {
	appDao := dao.AppDao{GormEngine: util.GlobalGormEngine}
	//判断出前十个出现次数最多的会议室,要分会议室数目是否大于十分两种情况
	return appDao.QueryHotMeetingRoomsDao()
}

func (appService *AppService) QueryAllUserService(currentPage int,num int,username string) []model.User {
	appDao := dao.AppDao{GormEngine: util.GlobalGormEngine}
	return appDao.QueryAllUserDao(currentPage,num,username)
}

func (appService *AppService) FreezeUserService(username string) uint {
	appDao := dao.AppDao{GormEngine: util.GlobalGormEngine}
	return appDao.FreezeUserDao(username)
}

func (appService *AppService) DeleteSuperUserService(username string) uint {
	appDao := dao.AppDao{GormEngine: util.GlobalGormEngine}
	return appDao.DeleteSuperUserDao(username)
}

func (appService *AppService) QueryAllMeetingRoomService(currentPage int,num int,room string) []model.MeetingRoom {
	appDao := dao.AppDao{GormEngine: util.GlobalGormEngine}
	return appDao.QueryAllMeetingRoomDao(currentPage,num,room)
}

func (appService *AppService) UpdateRoomNameService(roomName string,newRoomName string) uint {
	appDao := dao.AppDao{GormEngine: util.GlobalGormEngine}
	return appDao.UpdateRoomNameDao(roomName,newRoomName)
}

func (appService *AppService) DeleteMeetingRoomService(roomName string) uint {
	appDao := dao.AppDao{GormEngine: util.GlobalGormEngine}
	return appDao.DeleteMeetingRoomDao(roomName)
}

func (appService *AppService) AddMeetingRoomService(roomName string) uint {
	appDao := dao.AppDao{GormEngine: util.GlobalGormEngine}
	return appDao.AddMeetingRoomDao(roomName)
}

func (appService *AppService) QueryAllMeetingService(currentPage int, num int, name string, username string, time string) []model.Meeting {
	appDao := dao.AppDao{GormEngine: util.GlobalGormEngine}
	return appDao.QueryMeetingDao(currentPage,num,name,username,time)
}

func (appService *AppService) DeleteMeetingService(id int) uint {
	appDao := dao.AppDao{GormEngine: util.GlobalGormEngine}
	return appDao.DeleteMeetingDao(id)
}

func (appService *AppService) ShowMeetingReportService() []model.MeetingCount{
	appDao := dao.AppDao{GormEngine: util.GlobalGormEngine}
	return appDao.ShowMeetingReportDao()
}