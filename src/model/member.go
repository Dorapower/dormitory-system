package model

import (
	"dormitory-system/src/database"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type GroupsUser struct {
	ID        int `gorm:"primaryKey;autoIncrement"`
	Uid       int
	IsCreator int `gorm:"default:0"`
	GroupId   int
	IsDel     int `gorm:"default:0"`
	JoinTime  int
	LeaveTime int
	Status    int `gorm:"default:0"`
}

// GetUserGroup : check if user has a group
func GetUserGroup(uid int) int {
	var db = database.MysqlDb
	var groupUser GroupsUser
	result := db.Model(GroupsUser{}).Where("uid = ? and is_del = ?", uid, 0).First(&groupUser)
	if result.Error == gorm.ErrRecordNotFound {
		return 0
	}
	return groupUser.GroupId
}

// JoinGroup : join a group by invite code
//
//				 return 0 : success
//	                 	1 : already have a group
//			            2 : invite code wrong
//						3 : gender wrong
//			            4 : group members are full
func JoinGroup(uid int, inviteCode string) int {
	var db = database.MysqlDb

	// check if already have a group
	var groupId int
	groupId = GetUserGroup(uid)
	if groupId != 0 {
		return 1
	}

	// get group's inviteCode
	var code string
	db.Model(Groups{}).Select("invite_code").Where("id = ?", groupId).Scan(&code)
	if code != inviteCode {
		return 2
	}

	// check gender
	user := GetUserByUid(uid)
	var groupGender int
	subQuery1 := db.Model(Groups{}).Select("id").Where("invite_code = ? and is_del = ?", inviteCode, 0)
	subQuery2 := db.Model(GroupsUser{}).Select("uid").Where("group_id = (?) and is_creator = ?", subQuery1, 1)
	db.Model(Users{}).Select("gender").Where("uid = (?)", subQuery2).Scan(&groupGender)
	if groupGender != user.Gender {
		return 3
	}

	// check current group members if up to max
	var currentCnt int64
	db.Model(&GroupsUser{}).Where("group_id = ?", groupId).Count(&currentCnt)
	key, _ := strconv.Atoi(GetSystemConfigByKey("group_num").KeyValue)
	if int(currentCnt) == key {
		return 4
	}

	// check if user has quited the group before
	var memId int
	db.Model(GroupsUser{}).Select("id").Where("uid = ? and group_id = ? and is_del = ?", uid, groupId, 1).Scan(&memId)
	if memId != 0 {
		db.Model(GroupsUser{}).Where("uid = ? and group_id = ? and is_del = ?", uid, groupId, 1).Updates(map[string]interface{}{"is_del": 0, "leave_time": 0})
		return 0
	}

	var mem = GroupsUser{
		Uid:       uid,
		GroupId:   groupId,
		JoinTime:  int(time.Now().Unix()),
		LeaveTime: 0,
	}
	db.Create(&mem)
	return 0
}

func QuitGroup(uid int) bool {
	var groupId int
	groupId = GetUserGroup(uid)
	if groupId == 0 {
		return false
	}

	var db = database.MysqlDb
	result := db.Model(GroupsUser{}).Where("uid = ? and group_id = ?", uid, groupId).Updates(map[string]interface{}{"is_del": 1, "leave_time": int(time.Now().Unix())})
	if result.Error != nil {
		return false
	}
	return true
}

type MemberApi struct {
	StudentID   string `json:"student_id" gorm:"column:studentid"`
	StudentName string `json:"student_name" gorm:"column:name"`
}

type MyGroupApi struct {
	GroupID    int         `json:"group_id"`
	GroupName  string      `json:"group_name"`
	InviteCode string      `json:"invite_code"`
	Members    []MemberApi `json:"members"`
}

// GetMyGroup : get my group member's information (include myself)
func GetMyGroup(uid int) (myGroup MyGroupApi) {
	var db = database.MysqlDb

	// get group information
	var group Groups
	db.Where("id = (?)", db.Model(GroupsUser{}).Select("group_id").Where("uid = ? and is_del = ?", uid, 0)).First(&group)
	myGroup.GroupID = group.ID
	myGroup.GroupName = group.Name
	myGroup.InviteCode = group.InviteCode

	// get members' information
	rows, _ := db.Model(GroupsUser{}).Where("group_id = ? and is_del = ?", group.ID, 0).Rows()
	for rows.Next() {
		var mem GroupsUser
		err := db.ScanRows(rows, &mem)
		if err != nil {
			return MyGroupApi{}
		}
		var memApi MemberApi
		db.Model(StudentInfo{}).Select("student_info.studentid, users.name").Joins("left join users on student_info.uid = users.uid").Where("student_info.uid = ?", mem.Uid).Scan(&memApi)
		myGroup.Members = append(myGroup.Members, memApi)
	}
	return
}

// TransferGroup : transfer group to other
func TransferGroup(uid int, sId string) bool {
	var db = database.MysqlDb
	var groupId int
	groupId = GetUserGroup(uid)
	if groupId == 0 {
		return false
	}

	// check if sId(studentId) is right
	var otherUid int
	db.Model(StudentInfo{}).Select("uid").Where("studentid = ?", sId).Scan(&otherUid)
	if otherUid == 0 {
		return false
	}
	otherGroupId := GetUserGroup(otherUid)
	if otherGroupId == 0 || groupId != otherGroupId {
		return false
	}

	db.Model(GroupsUser{}).Where("uid = ? and is_del = ?", uid, 0).Update("is_creator", 0)
	db.Model(GroupsUser{}).Where("uid = ? and is_del = ?", otherUid, 0).Update("is_creator", 1)
	return true
}
