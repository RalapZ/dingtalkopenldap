package command

import (
	"fmt"
	"github.com/RalapZ/dingtalkopenldap/model"
	"time"
)

func Start(){
	//go ScheduleUpdateSub()//启动定时更新部门信息任务
	//go ChangeUserInfoFunc()
	model.GetToken("GET")
	UrlDepSubId := model.GetListSubIdUrl+"?access_token=" + model.Token
	model.InitListSubId("POST",1 , UrlDepSubId)
	//fmt.Println()
	time.Sleep(1000*time.Microsecond)
	model.GetListUserInfoMap("POST",0,model.GetUserListUrl+"?access_token="+model.Token)
	getuserdetailinfo:=model.GetUserDetailUrl+"?access_token="+model.Token

	for _,v:= range model.Userlist{
		model.GetUserDetailInfo("POST",v,getuserdetailinfo)
	}
	//fmt.Println("main",model.DepListId)
	go SchedulerTimeFunc() //启动定时器
	//go model.UpdateUserInfo()
	for _,v:=range model.Userlist{
		model.LDAPservice.AddUserinfo(v)
	}
	fmt.Println("start end")

}

