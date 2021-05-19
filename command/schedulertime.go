package command

import (
	"github.com/RalapZ/dingtalkopenldap/model"
	"time"
)

func SchedulerTimeFunc(){
	for {
		select{
		case <-time.After(time.Duration(model.Defaultconfig.SchedulerTime)):
			model.ScheduTimeChan<-true
		default:
			time.Sleep(5*time.Second)
		}
	}
}
