package command

import (
	"fmt"
	"github.com/RalapZ/dingtalkopenldap/model"
	"time"
)

func SchedulerTimeFunc(){
	for {
		select{
		case <-time.After(time.Duration(model.Defaultconfig.SchedulerTime)):
			//model.ScheduTimeChan<-true
			model.CheckUserInfo("POST",0,model.GetUserListUrl+"?access_token="+model.Token)
		default:
			time.Sleep(5*time.Second)
		}
	}
}


func ChangeUserInfoFunc(){
	go func() {
		for {
			select {
			case userid := <-model.UserChangeChan:
				fmt.Println("user id change ", userid)
			default:
				time.Sleep(time.Millisecond)
			}
		}
	}()
}
