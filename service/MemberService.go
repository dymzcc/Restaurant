package service

import (
	"Areyouanxiety/dao"
	"Areyouanxiety/model"
	"Areyouanxiety/tool"
	"strconv"
	"time"
)

type MemberService struct{

}

func (ms *MemberService) GetUserInfo(userId string)  *model.Member{
	id, err := strconv.Atoi(userId)
	if err != nil{
		return nil
	}
	return new(dao.MemberDao).QueryMemberByID(int64(id))
}


func (ms *MemberService) Login(name string, password string) *model.Member{
	//1. 使用用户名+密码查询用户信息
	member := new(dao.MemberDao).Query(name, password)
	if member.Id != 0{
		return member
	}

	//2. 如果用户不存在，作为新用户保存到数据库中
	user := model.Member{}
	user.UserName = name
	user.Password = tool.EncoderSha256(password)
	user.RegisterTime = time.Now().Unix()
	result := new(dao.MemberDao).InsertMember(user)
	user.Id = result
	return &user
}

func (ms *MemberService)UploadAvatar(userId int64, fileName string) string {
	result := new(dao.MemberDao).UpdateMemberAvatar(userId, fileName)
	if result == 0{
		return ""
	}
	return fileName
}