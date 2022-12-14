package model

import (
	"dormitory-system/src/database"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type Groups struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	Name       string
	InviteCode string
	Describe   string
	IsDel      int `gorm:"default:0"`
	Status     int `gorm:"default:0"`
}

type CreatGroupApi struct {
	InviteCode string `json:"invite_code"`
	TeamId     int    `json:"team_id"`
}

// generate a random invite_code
func generateCode() string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 5; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// CreatGroup : creat a group, return group's id and invite_code
func CreatGroup(uid int, name, describe string) (CreatGroupApi, bool) {

	// check if already have a group
	groupId := GetUserGroup(uid)
	if groupId != 0 {
		return CreatGroupApi{}, false
	}

	var db = database.MysqlDb
	var cApi CreatGroupApi
	var group Groups
	var code string

	// check if code have already existed
	for {
		code = generateCode()
		result := db.Where("invite_code = ? and is_del = ?", code, 0).First(&group)
		// if not exist, jump out the loop
		if result.Error == gorm.ErrRecordNotFound {
			break
		}
	}

	// add to database
	group = Groups{}
	group.Name = name
	group.Describe = describe
	group.InviteCode = code
	db.Create(&group)

	cApi.TeamId = group.ID
	cApi.InviteCode = group.InviteCode

	var mem = GroupsUser{
		Uid:       uid,
		IsCreator: 1,
		GroupId:   group.ID,
		JoinTime:  int(time.Now().Unix()),
		LeaveTime: 0,
	}
	db.Create(&mem)
	return cApi, true
}

// DelGroup : delete a group by id
func DelGroup(uid int) bool {
	var db = database.MysqlDb

	// check if user have a group
	groupId := GetUserGroup(uid)
	if groupId == 0 {
		return false
	}

	var mem GroupsUser
	result := db.Where("group_id = ? and is_creator = ? and is_del = ?", groupId, 0, 0).First(&mem)

	// no member in group, delete it
	if result.Error == gorm.ErrRecordNotFound {
		db.Model(GroupsUser{}).Where("group_id = ? and is_creator = ?", groupId, 1).Updates(map[string]interface{}{"is_del": 1, "leave_time": int(time.Now().Unix())})
		db.Model(Groups{}).Where("id = ?", groupId).Update("is_del", 1)
		return true
	}
	return false
}
