package command

import (
	"github.com/RalapZ/dingtalkopenldap/model"
	log "github.com/sirupsen/logrus"
	"time"
)

func SchedulerTimeFunc(){
	for {
		select{
		case <-time.After(time.Duration(model.Defaultconfig.SchedulerTime)*time.Second):
			//model.ScheduTimeChan<-true
			log.Infof("开始更新检测部门信息")
			UrlDepSubId := model.GetListSubIdUrl+"?access_token=" + model.Token
			model.UpdataDepListIdAndDepListDetailInfo("POST",1 , UrlDepSubId)//更新department列表信息
			log.Info("开发更新检测用户信息")
			model.CheckUserInfo("POST",0,model.GetUserListUrl+"?access_token="+model.Token)
		//default:
			//time.Sleep(5*time.Second)
			//time.Sleep(5*time.Second)
		}
	}
}


func ScheduleUpdateSub(){
	for{
		select {
		case <-time.After(time.Duration(model.Defaultconfig.SchedulerTime)):
			UrlDepSubId := model.GetListSubIdUrl+"?access_token=" + model.Token
			model.UpdataDepListIdAndDepListDetailInfo("POST",1 , UrlDepSubId)
		}
	}
}

//goroutine检测user info是否更新
func ChangeUserInfoFunc(){
		go func() {
			for {
				select {
				case userid := <-model.UserChangeChan:
					//fmt.Println("ChangeUserInfoFunc","user id change ", userid)
					model.LDAPservice.ModifyUserinfo(userid)
				default:
					time.Sleep(100*time.Millisecond)
				}
			}
		}()
}
