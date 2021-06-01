package dao

import (
	"Areyouanxiety/model"
	"Areyouanxiety/tool"
	"fmt"
)

type MemberDao struct {

}

//根据ID查找
func (md *MemberDao) QueryMemberByID(id int64) *model.Member{
	var member model.Member
	err := tool.DB.Where("id=?", id).Find(&member).Error
	if err != nil{
		return nil
	}
	return &member
}

//根据用户名和密码查询
func (md *MemberDao) Query(name string, password string) *model.Member{
	var member model.Member
	password = tool.EncoderSha256(password)
	err := tool.DB.Where("user_name=? and password=?", name, password).First(&member).Error
	if err != nil {
		panic(err.Error())
	}
	return &member
}

//新用户的数据库插入操作
func (md *MemberDao)InsertMember(member model.Member)int64{
	err:= tool.DB.Create(&member)
	if err!=nil{
		fmt.Println("*********************")
		fmt.Println(err)
		return 0
	}
	return member.Id
}

//更新用户头像
func (md *MemberDao) UpdateMemberAvatar(userId int64, filename string) int64{
	member := model.Member{Avatar: filename}
	err := tool.DB.Where("id=?", userId).Update(&member).Error
	if err != nil {
		panic(err.Error())
	}
	return member.Id
}