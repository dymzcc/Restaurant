package controller

import (
	"Areyouanxiety/model"
	"Areyouanxiety/param"
	"Areyouanxiety/service"
	"Areyouanxiety/tool"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"time"
)

type MemberController struct {

}

func (mc *MemberController) Router(engine *gin.Engine){
	engine.GET("/api/captcha", mc.captcha)
	engine.POST("/api/login_pwd", mc.nameLogin)
	engine.POST("/api/upload/avator", mc.uploadAvator)
	engine.GET("/api/userinfo", mc.userInfo)
}

func (mc *MemberController) nameLogin(context *gin.Context){
	//1. 解析用户登录的参数
	var loginParam param.LoginParam
	err := tool.Decode(context.Request.Body, &loginParam)
	if err != nil{
		tool.Failed(context, "参数解析失败")
		return
	}

	//2. 验证验证码
	validate := tool.VertifyCaptcha(loginParam.Id, loginParam.Value)
	if !validate{
		tool.Failed(context, "验证码不正确，请重新验证")
		return
	}

	//3. 登录
	ms := service.MemberService{}
	member := ms.Login(loginParam.Name, loginParam.Password)
	if member.Id != 0{
		tool.Success(context, &member)
		//用户信息保存到session
		sess, _ := json.Marshal(member)
		err = tool.Setsession(context, "user" + string(member.Id), sess)
		if(err != nil){
			tool.Failed(context, "登录失败")
		}
		fmt.Println("登录成功")
		return
	}
	tool.Failed(context, "登录失败")
}

//生成验证码
func (mc *MemberController) captcha(contex *gin.Context)  {
	tool.GenerateCaptcha(contex)
}


func (mc *MemberController) uploadAvator(context *gin.Context){
	//1. 解析参数：file、user-id
	userId := context.PostForm("user_id")
	file, err := context.FormFile("avatar")
	if err != nil || userId == ""{
		tool.Failed(context, "参数解析失败")
		return
	}

	//2. 判断用户是否登录
	sess := tool.Getsession(context, "user_"+userId)
	if sess == nil{
		tool.Failed(context, "参数不合法")
	}
	var member model.Member
	json.Unmarshal(sess.([]byte), &member)

	//3. file 保存到本地
	//为了确保文件名唯一，命名时加入时间参数
	fileName := "./uploadfile/" + strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	err = context.SaveUploadedFile(file, fileName)
	if err != nil{
		tool.Failed(context, "头像更新失败")
		return
	}
	//然后将文件上传到fastdfs上
	fileId := tool.UploadFile(fileName)
	if fileId != ""{
		//上传成功，删除本地的文件
		os.Remove(fileName)
		//4. 将保存后的文件本地路径 保存到用户表中的头像字段
		memberServiece := service.MemberService{}
		//截取filename字符串，删去开头的 .
		path := memberServiece.UploadAvatar(member.Id, fileName[1:])
		if path != ""{
			tool.Success(context, tool.FileServerAddr() + "/" + path)
		}
		tool.Failed(context, "上传失败")
	}
}

func (mc *MemberController) userInfo(context *gin.Context){
	cookie, err := tool.CookieAuth(context)
	if err != nil{
		context.Abort()
		tool.Failed(context, "还为登录，请先登录")
		return
	}
	member := new(service.MemberService).GetUserInfo(cookie.Value)
	if member != nil{
		tool.Success(context, map[string]interface{}{
			"id":            member.Id,
			"user_name":     member.UserName,
			"mobile":        member.Mobile,
			"register_time": member.RegisterTime,
			"avatar":        member.Avatar,
			"balance":       member.Balance,
			"city":          member.City,
		})
	}
	tool.Failed(context,"获取用户信息失败")
}
