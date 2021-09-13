package dao

import (
	"log"
	"meeting_root_server/model"
	"meeting_root_server/util"
)

type AppDao struct {
	*util.GormEngine
}

func (appDao *AppDao) LoginDao(username string, password string) (uint, *model.SuperUser) {
	var user model.SuperUser
	appDao.Where("username = ?", username).Where("password = ?", password).First(&user)

	//可以使用下面这种方式判断结构体是否为空，但这个时候要求结构体的所有字段都是可比较的
	//if user == (model.User{}){
	//	return false
	//}
	//下面这种反射的方式就没有上面方法的局限性，但是反射的性能较低
	if user.UserIsEmpty() {
		return 1, &user
	} else {
		return 0, &user
	}
}

func (appDao *AppDao) QueryMeetingDao(currentPage int, num int, name string, username string, time string) []model.Meeting {
	var meetings []model.Meeting
	if num == 0 {
		if name != "no" && username != "no" && time != "no" {
			appDao.Where("user = ?", username).Where("borrow_time = ?", time).Where("meeting_room_name = ?", name).Find(&meetings)
		} else if name == "no" && username != "no" && time != "no" {
			appDao.Where("user = ?", username).Where("borrow_time = ?", time).Find(&meetings)
		} else if name != "no" && username == "no" && time != "no" {
			appDao.Where("borrow_time = ?", time).Where("meeting_room_name = ?", name).Find(&meetings)
		} else if name != "no" && username != "no" && time == "no" {
			appDao.Where("user = ?", username).Where("meeting_room_name = ?", name).Find(&meetings)
		} else if name == "no" && username == "no" && time != "no" {
			appDao.Where("borrow_time = ?", time).Find(&meetings)
		} else if name == "no" && username != "no" && time == "no" {
			appDao.Where("user = ?", username).Find(&meetings)
		} else if name != "no" && username == "no" && time == "no" {
			appDao.Where("meeting_room_name = ?", name).Find(&meetings)
		} else if name == "no" && username == "no" && time == "no" {
			appDao.Find(&meetings)
		} else {
			appDao.Find(&meetings)
		}
	} else {
		if name != "no" && username != "no" && time != "no" {
			appDao.Where("user = ?", username).Where("borrow_time = ?", time).Where("meeting_room_name = ?", name).Limit(num).Offset((currentPage - 1) * num).Find(&meetings)
		} else if name == "no" && username != "no" && time != "no" {
			appDao.Where("user = ?", username).Where("borrow_time = ?", time).Limit(num).Offset((currentPage - 1) * num).Find(&meetings)
		} else if name != "no" && username == "no" && time != "no" {
			appDao.Where("borrow_time = ?", time).Where("meeting_room_name = ?", name).Limit(num).Offset((currentPage - 1) * num).Find(&meetings)
		} else if name != "no" && username != "no" && time == "no" {
			appDao.Where("user = ?", username).Where("meeting_room_name = ?", name).Limit(num).Offset((currentPage - 1) * num).Find(&meetings)
		} else if name == "no" && username == "no" && time != "no" {
			appDao.Where("borrow_time = ?", time).Limit(num).Offset((currentPage - 1) * num).Find(&meetings)
		} else if name == "no" && username != "no" && time == "no" {
			appDao.Where("user = ?", username).Limit(num).Offset((currentPage - 1) * num).Find(&meetings)
		} else if name != "no" && username == "no" && time == "no" {
			appDao.Where("meeting_room_name = ?", name).Limit(num).Offset((currentPage - 1) * num).Find(&meetings)
		} else if name == "no" && username == "no" && time == "no" {
			appDao.Limit(num).Offset((currentPage - 1) * num).Find(&meetings)
		} else {
			appDao.Limit(num).Offset((currentPage - 1) * num).Find(&meetings)
		}
	}
	return meetings
}

func (appDao *AppDao) UpdatePasswordDao(username string, password string) uint {
	var user model.SuperUser
	appDao.Where("username = ?", username).First(&user)
	if err := appDao.Model(&user).Update("password", password).Error; err != nil {
		log.Print("更新失败")
		return 0
	}
	return 1
}

func (appDao *AppDao) DeleteUserDao(username string) uint {
	if err := appDao.Where("username = ?", username).Delete(&model.User{}).Error; err != nil {
		log.Print("删除失败")
		return 0
	}
	appDao.Where("user = ?", username).Delete(&model.Meeting{})
	return 1
}

func (appDao *AppDao) QueryMeetingRoomCountDao() uint {
	var count uint
	appDao.Model(&model.MeetingRoom{}).Count(&count)
	return count
}

func (appDao *AppDao) QueryUserCountDao() uint {
	var count uint
	appDao.Model(&model.User{}).Count(&count)
	return count
}

func (appDao *AppDao) QueryMeetingCountDao() uint {
	var count uint
	appDao.Model(&model.Meeting{}).Count(&count)
	return count
}

func (appDao *AppDao) QueryLatestMeetingsDao() []model.Meeting {
	var meetings []model.Meeting
	appDao.Order("meeting_id DESC").Limit(10).Find(&meetings)
	return meetings
}

func (appDao *AppDao) QueryHotMeetingRoomsDao() []model.MeetingCount {
	var count []model.MeetingCount
	//select meeting_room_name,count(*) from meetings group by meeting_room_name order by count(*)  desc;
	rows, err := appDao.Table("meetings").Select("meeting_room_name,count(*)").Group("meeting_room_name").Order("count(*) desc").Limit(10).Rows()
	defer rows.Close()
	if err != nil {
		log.Print("查询失败")
	}
	for rows.Next() {
		var c model.MeetingCount
		rows.Scan(&c.MeetingRoomName, &c.Count)
		count = append(count, c)
	}
	return count
}

func (appDao *AppDao) ShowMeetingReportDao() []model.MeetingCount {
	var count []model.MeetingCount
	//select meeting_room_name,count(*) from meetings group by meeting_room_name order by count(*)  desc;
	rows, err := appDao.Table("meetings").Select("meeting_room_name,count(*)").Group("meeting_room_name").Rows()
	defer rows.Close()
	if err != nil {
		log.Print("查询失败")
	}
	for rows.Next() {
		var c model.MeetingCount
		rows.Scan(&c.MeetingRoomName, &c.Count)
		count = append(count, c)
	}
	return count
}

func (appDao *AppDao) QueryMeetingRoomDao() []model.MeetingRoom {
	var rooms []model.MeetingRoom
	appDao.Find(&rooms)
	return rooms
}

func (appDao *AppDao) QueryAllUserDao(currentPage int, num int, username string) []model.User {
	var users []model.User
	if num == 0 {
		if username == "no" {
			appDao.Find(&users)
		} else {
			appDao.Where("username = ?", username).Find(&users)
		}
	} else {
		if username == "no" {
			appDao.Limit(num).Offset((currentPage - 1) * num).Find(&users)
		} else {
			appDao.Where("username = ?", username).Limit(num).Offset((currentPage - 1) * num).Find(&users)
		}
	}
	if len(users) == 0 { //说明通过用户名没取到,这时再看通过手机号查询
		if num == 0 {
			if username == "no" {
				appDao.Find(&users)
			} else {
				appDao.Where("phone = ?", username).Find(&users)
			}
		} else {
			if username == "no" {
				appDao.Limit(num).Offset((currentPage - 1) * num).Find(&users)
			} else {
				appDao.Where("phone = ?", username).Limit(num).Offset((currentPage - 1) * num).Find(&users)
			}
		}
	}
	return users
}

func (appDao *AppDao) FreezeUserDao(username string) uint {
	var user model.User
	appDao.Where("username = ?", username).First(&user)
	if user.State == 0 {
		if err := appDao.Model(&user).Update("state", 1).Error; err != nil {
			log.Print("更新失败")
			return 0
		}
	} else {
		if err := appDao.Model(&user).Update("state", 0).Error; err != nil {
			log.Print("更新失败")
			return 0
		}
	}
	return 1
}

func (appDao *AppDao) DeleteSuperUserDao(username string) uint {
	if err := appDao.Where("username = ?", username).Delete(&model.SuperUser{}).Error; err != nil {
		log.Print("删除失败")
		return 0
	}
	return 1
}

func (appDao *AppDao) QueryAllMeetingRoomDao(currentPage int, num int, room string) []model.MeetingRoom {
	var rooms []model.MeetingRoom
	if num == 0 {
		if room == "no" {
			appDao.Find(&rooms)
		} else {
			appDao.Where("room_name = ?", room).Find(&rooms)
		}
	} else {
		if room == "no" {
			appDao.Limit(num).Offset((currentPage - 1) * num).Find(&rooms)
		} else {
			appDao.Where("room_name = ?", room).Limit(num).Offset((currentPage - 1) * num).Find(&rooms)
		}
	}
	return rooms
}

func (appDao *AppDao) UpdateRoomNameDao(roomName string, newRoomName string) uint {
	var room model.MeetingRoom
	//首先判断输入的会议室名称是否已经存在
	appDao.Where("room_name = ?", newRoomName).First(&room)
	if !room.MeetingRoomIsEmpty() { //输入的会议室名称已存在
		return 0
	}
	appDao.Where("room_name = ?", roomName).First(&room)
	if err := appDao.Model(&room).Update("room_name", newRoomName).Error; err != nil {
		log.Print("更新失败")
		return 1
	}
	return 2
}

func (appDao *AppDao) DeleteMeetingRoomDao(roomName string) uint {
	if err := appDao.Where("room_name = ?", roomName).Delete(&model.MeetingRoom{}).Error; err != nil {
		log.Print("删除失败")
		return 0
	}
	return 1
}

func (appDao *AppDao) AddMeetingRoomDao(name string) uint {
	var room model.MeetingRoom
	appDao.Where("room_name = ?", name).First(&room)
	if !room.MeetingRoomIsEmpty() {
		log.Print("会议室已存在")
		return 0
	}
	if err := appDao.Create(&model.MeetingRoom{RoomName: name}).Error; err != nil {
		log.Print("插入失败")
		return 1
	}
	return 2
}

func (appDao *AppDao) DeleteMeetingDao(id int) uint {
	if err := appDao.Where("meeting_id = ?", id).Delete(&model.Meeting{}).Error; err != nil {
		log.Print("删除失败")
		return 0
	}
	return 1
}
