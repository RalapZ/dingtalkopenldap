package main

import (
	"github.com/RalapZ/dingtalkopenldap/command"
	"github.com/RalapZ/dingtalkopenldap/model"
)


func init(){
	model.InitConfig()
	model.InitLdapConnection()
	//model.GetToken("GET")

	//log.
}

func main() {
	command.InitLdap()
	//fmt.Println("test",model.Userlist)
}