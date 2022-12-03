package model

import (
	"dormitory-system/src/database"
	"gorm.io/gorm"
	"strconv"
	"time"
)

const GroupMaxPeople = 4

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

// CheckUserGroup : check if user has a group
func CheckUserGroup(uid int) bool {
	var db = database.MysqlDb
	result := db.Model(GroupsUser{}).Select("id").Where("uid = ? and is_del = ?", uid, 0)
	if result.Error == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

// JoinGroup : join a group by invite code
//
//		 return 0 : success
//	            1 : invite code wrong
//				2 : gender wrong
//	            3 : group members are full
func JoinGroup(uid int, inviteCode string) int {
	var db = database.MysqlDb

	// get group's id, if the group exist
	var groupId int
	result := db.Model(Groups{}).Select("id").Where("invite_code = ? and is_del = ?", inviteCode, 0).Scan(&groupId)
	if result.Error == gorm.ErrRecordNotFound {
		return 1
	}

	// check gender
	user := GetUserByUid(uid)
	var groupGender int
	subQuery1 := db.Model(Groups{}).Select("id").Where("invite_code = ? and is_del = ?", inviteCode, 0)
	subQuery2 := db.Model(GroupsUser{}).Select("uid").Where("group_id = (?) and is_creator = ?", subQuery1, 1)
	db.Model(Users{}).Select("gender").Where("uid = (?)", subQuery2).Scan(&groupGender)
	if groupGender != user.Gender {
		return 2
	}

	// check current group members if up to max
	var currentCnt int64
	db.Model(&GroupsUser{}).Where("group_id = ?", groupId).Count(&currentCnt)
	groupNum, _ := strconv.Atoi(GetSystemConfigByKey("group_num").KeyValue)
	if int(currentCnt) == groupNum {
		return 3
	}

	// check if user has quited the group before
	result = db.Model(GroupsUser{}).Select("id").Where("uid = ? and group_id = ? and is_del = ?", uid, groupId, 1)
	if result.Error != gorm.ErrRecordNotFound {
		db.Model(GroupsUser{}).Where("uid = ? and group_id = ? and is_del = ?", uid, groupId, 1).Update("is_del", 0)
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

func QuitGroup(uid, groupId int) bool {
	var db = database.MysqlDb
	result := db.Model(GroupsUser{}).Where("uid = ? and group_id = ?", uid, groupId).Updates(map[string]interface{}{"is_del": 1, "leave_time": int(time.Now().Unix())})
	if result.Error != nil {
		return false
	}
	return true
}

type MemberApi struct {
	student_id   string
	student_name string
}

type MyGroupApi struct {
	group_id    int
	group_name  string
	invice_code string
	members     []MemberApi
}

// GetMyGroup : get my group member's information (include myself)
func GetMyGroup(uid int) (myGroup MyGroupApi) {
	var db = database.MysqlDb

	// get group information
	var group Groups
	db.Where("id = (?)", db.Model(GroupsUser{}).Select("group_id").Where("uid = ? and is_del = ?", uid, 0)).First(&group)
	myGroup.group_id = group.ID
	myGroup.group_name = group.Name
	myGroup.invice_code = group.InviteCode

	// get members' information
	rows, _ := db.Model(GroupsUser{}).Select("uid").Where("group_id = ? and is_del = ?", group.ID, 0).Rows()
	for rows.Next() {
		var memId int
		db.ScanRows(rows, &memId)
		var mem MemberApi
		db.Model(StudentInfo{}).Select("student_info.studentid, users.name").Joins("left join users on student_info.uid = users.uid").Where("uid = ?", memId).Scan(&mem)
		myGroup.members = append(myGroup.members, mem)
	}
	return
}

// TransferGroup : transfer group to other
func TransferGroup(uid int, sId string) bool {
	var db = database.MysqlDb
	// check if sId(studentId) is right
	var otherUid int
	result := db.Model(StudentInfo{}).Select("uid").Where("studentid = ?", sId).Scan(&otherUid)
	if result.Error != nil {
		return false
	}

	db.Model(GroupsUser{}).Where("uid = ? and is_del = ?", uid, 0).Update("is_creator", 0)
	db.Model(GroupsUser{}).Where("uid = ? and is_del = ?", otherUid, 0).Update("is_creator", 1)
	return true
}
