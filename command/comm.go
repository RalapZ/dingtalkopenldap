package command

import (
	"github.com/RalapZ/dingtalkopenldap/model"
	log "github.com/sirupsen/logrus"

	"time"
)

func Init(){
	go model.ScheduleUpdateSub()//启动定时更新任务
	model.GetToken("GET")
	log.Println(model.Token)
	UrlDepSubId := model.GetListSubIdUrl+"?access_token=" + model.Token
	//model.GetListSubId("POST",1 , UrlDepSubId)
	model.InitListSubId("POST",1 , UrlDepSubId)
	//for k,v :=range model.DepListDetailInfo{
	//	fmt.Println("print",k,v)
	//}
	//GetListSubId

	time.Sleep(1000*time.Microsecond)
	log.Println(model.DepListId)

	model.GetListUserInfoMap("POST",0,model.GetUserListUrl+"?access_token="+model.Token)
	//fmt.Println("main",model.Userlist)
	getuserdetailinfo:=model.GetUserDetailUrl+"?access_token="+model.Token
	for _,v:= range model.Userlist{
		model.GetUserDetailInfo("POST",v,getuserdetailinfo)
	}
	//
	//for k,v:=range model.DepListDetailInfo{
	//	log.Println(k,v)
	//}
	for _,v:=range model.Userlist{
		model.LDAPservice.AddUserinfo(v)
	}
	//for k,v:=range model.UserListDetailInfo{
	//	log.Println("test",k,v)
	//}
}

