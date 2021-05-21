package command

import (
	"fmt"
	"github.com/RalapZ/dingtalkopenldap/model"
	"time"
)

func Start(){
	go model.ScheduleUpdateSub()//启动定时更新任务
	fmt.Println("test")
	//go SchedulerTimeFunc()
	model.GetToken("GET")
	UrlDepSubId := model.GetListSubIdUrl+"?access_token=" + model.Token
	model.InitListSubId("POST",1 , UrlDepSubId)
	time.Sleep(1000*time.Microsecond)
	model.GetListUserInfoMap("POST",0,model.GetUserListUrl+"?access_token="+model.Token)
	getuserdetailinfo:=model.GetUserDetailUrl+"?access_token="+model.Token

	for _,v:= range model.Userlist{
		model.GetUserDetailInfo("POST",v,getuserdetailinfo)
	}
	go model.UpdateUserInfo()
	for _,v:=range model.Userlist{
		model.LDAPservice.AddUserinfo(v)
	}

}

