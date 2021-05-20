package command

import (
	"fmt"
	"github.com/RalapZ/dingtalkopenldap/model"
	"log"
	"time"
)

func InitLdap(){
	//fmt.Println(model.AuthConf, model.Listenconfig, model.Authconfig, model.Ldapconfig, model.DBconfig)
	//model.InitLdapConnection()
	//AppKey := "dingicgycuisw7lrqcj8"
	//AppSecret := "pUTLevcbvk9EWdJKF7kzj5-Txc_7CHdyrO9eFiObLj7Qb6F3Y7q8YfXnwhtIFTvI"
	//Url_token := "https://oapi.dingtalk.com/gettoken?appkey=" + AppKey + "&appsecret=" + AppSecret
	//Url_token:=model.GetTokenUrl+"?appkey="+ model.Authconfig.AppKey+"&appsecret=" + model.Authconfig.AppSecret
	go model.ScheduleUpdateSub()
	model.GetToken("GET")
	log.Println(model.Token)

	UrlDepSubId := model.GetListSubIdUrl+"?access_token=" + model.Token
	//model.GetListSubId("POST",1 , UrlDepSubId)
	model.InitListSubId("POST",1 , UrlDepSubId)
	for k,v :=range model.DepListDetailInfo{
		fmt.Println("print",k,v)
	}
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
