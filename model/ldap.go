package model

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
	log "github.com/sirupsen/logrus"
	"gopkg.in/ldap.v2"
	"os"
)

//var LdapConn *ldap.Conn

var (
	LDAPconn    ldap.Conn
	LDAPservice LDAPService
)

type LDAPService struct {
	Conn      *ldap.Conn
	Config     LdapConfig
}

func InitLdapConnection(){
	LDAPconn, err := ldap.Dial("tcp", Ldapconfig.Addr)
	if err!=nil{
		log.Error("Ldap connected port error ",err)

		os.Exit(3)
	}
	log.Info("Ldap connected port successfully")
	LDAPservice=LDAPService{LDAPconn,Ldapconfig}
	err = LDAPservice.Conn.Bind(Ldapconfig.BindUserName, Ldapconfig.BindPassword)
	if err !=nil{
		//fmt.Println(err)
		log.Error("Ldap authentication error ",err)
		os.Exit(3)
		//panic(err)
	}
	log.Info("Ldap  authentication successfully")
}


func (ldapservice *LDAPService)AddGroupinfo(groupinfo DepDetailInfo) {
	var grouptempinfo string
	for _,k:=range StackDepmentinfo{
		grouptempinfo="ou="+k+","+grouptempinfo
	}
	//fmt.Println(grouptempinfo)
	sql := ldap.NewAddRequest(grouptempinfo+Ldapconfig.SearchDN)
	sql.Attribute("ou", []string{StackDepmentinfo[len(StackDepmentinfo)-1]})
	sql.Attribute("objectClass", []string{"organizationalUnit","top"})
	er := LDAPservice.Conn.Add(sql)
	if er!=nil{
		log.Errorf("AddGroup  %s error  ",groupinfo.Name)
		log.Error(er)
	}else{
		log.Infof("AddGroup %s  successfully",groupinfo.Name)
	}
	log.Infof("add group %v info into ldap successfully ",groupinfo.Name)
	//log.Println()
}


func (ldapservice *LDAPService)ModifyGroupinfo(groupinfo DepDetailInfo) {
	var grouptempinfo string
	for _,k:=range StackDepmentinfo{
		grouptempinfo="ou="+k+","+grouptempinfo
	}
	//fmt.Println(grouptempinfo)
	sql := ldap.NewAddRequest(grouptempinfo+Ldapconfig.SearchDN)
	sql.Attribute("ou", []string{StackDepmentinfo[len(StackDepmentinfo)-1]})
	sql.Attribute("objectClass", []string{"organizationalUnit","top"})
	er := LDAPservice.Conn.Add(sql)
	if er!=nil{
		log.Errorf("AddGroup  %s error  ",groupinfo.Name)
		log.Error(er)
	}else{
		log.Infof("AddGroup %s  successfully",groupinfo.Name)
	}
	log.Infof("modify group %v successfully",groupinfo.Name)
	//log.Println()
}


func (ldapservice *LDAPService)AddUserinfo(userid string){
	var userldappath string
	for _,k:=range DepListDetailInfo[UserListDetailInfo[userid].DeptOrderList[0].DeptId].LdapDepPath{
		userldappath="ou="+k+","+userldappath
	}
	username := UserListDetailInfo[userid].Name
	convert := pinyin.Convert(username, nil)
	var fullname string
	var secondname string
	for _,k:=range convert{
		fullname=fullname+k[0]
	}
	for i:=1;i<len(convert);i++{
		secondname=secondname+convert[i][0]
	}
	firstname:=convert[0][0]
	usertempinfo := ldap.NewAddRequest("mail="+UserListDetailInfo[userid].Email+","+userldappath+Ldapconfig.SearchDN)
	usertempinfo.Attribute("userPassword", []string{"123456"})
	usertempinfo.Attribute("mail", []string{UserListDetailInfo[userid].Email})
	usertempinfo.Attribute("cn", []string{fullname})
	usertempinfo.Attribute("givenName", []string{firstname})
	usertempinfo.Attribute("sn", []string{UserListDetailInfo[userid].Name})
	//fmt.Println(userid,UserListDetailInfo[userid].Mobile)
	usertempinfo.Attribute("mobileTelephoneNumber", []string{UserListDetailInfo[userid].Mobile})
	usertempinfo.Attribute("uid", []string{UserListDetailInfo[userid].Userid})
	usertempinfo.Attribute("objectClass", []string{"inetOrgPerson","organizationalPerson"})
	er := LDAPservice.Conn.Add(usertempinfo)
	if er!=nil{
		log.Errorf("Add user  %s something error  ",UserListDetailInfo[userid].Name)
		log.Error(er)
	}else{
		log.Infof("Add user %s  successfully",UserListDetailInfo[userid].Name)
	}
	log.Infof("add user %v into ldap successfully",UserListDetailInfo[userid].Name)
	////usertempinfo.Attribute("uidNumber", []string{UserListDetailInfo[userid].Userid})
	//usertempinfo.Attribute("uidNumber", []string{"1"})
	//usertempinfo.Attribute("gidNumber", []string{strconv.Itoa(UserListDetailInfo[userid].DeptOrderList[0].DeptId)})
	//usertempinfo.Attribute("userPassword", []string{"123456"})
	//usertempinfo.Attribute("homeDirectory", []string{"/home/"+fullname})
	//usertempinfo.Attribute("mail", []string{UserListDetailInfo[userid].Email})
	//usertempinfo.Attribute("cn", []string{fullname})
	//usertempinfo.Attribute("givenName", []string{firstname})
	//usertempinfo.Attribute("sn", []string{secondname})
	//usertempinfo.Attribute("mobileTelephoneNumber", []string{"15711823061"})
	//usertempinfo.Attribute("title", []string{UserListDetailInfo[userid].Title})
	//usertempinfo.Attribute("uid", []string{UserListDetailInfo[userid].Userid})
	//usertempinfo.Attribute("objectClass", []string{"inetOrgPerson","organizationalPerson","posixAccount"})
	//er := LDAPservice.Conn.Add(usertempinfo)
	//if er!=nil{
	//	log.Println(er)
	//}

}



func (ldapservice *LDAPService)ModifyUserinfo(userid string){
	var userldappath string
	fmt.Println("ModifyUserinfo",userid)
	for _,k:=range DepListDetailInfo[UserListDetailInfo[userid].DeptOrderList[0].DeptId].LdapDepPath{
		userldappath="ou="+k+","+userldappath
	}
	username := UserListDetailInfo[userid].Name
	convert := pinyin.Convert(username, nil)
	var fullname string
	var secondname string
	for _,k:=range convert{
		fullname=fullname+k[0]
	}
	for i:=1;i<len(convert);i++{
		secondname=secondname+convert[i][0]
	}
	firstname:=convert[0][0]
	fmt.Println("ModifyUserinfo  done1")
	usertempinfo := ldap.NewModifyRequest("mail="+UserListDetailInfo[userid].Email+","+userldappath+Ldapconfig.SearchDN)
	//usertempinfo := ldap.NewModifyRequest("mail="+UserListDetailInfo[userid].Email+","+userldappath+Ldapconfig.SearchDN)
	usertempinfo.Replace("mail", []string{UserListDetailInfo[userid].Email})
	usertempinfo.Replace("cn", []string{fullname})
	usertempinfo.Replace("givenName", []string{firstname})
	usertempinfo.Replace("sn", []string{UserListDetailInfo[userid].Name})
	//fmt.Println(userid,UserListDetailInfo[userid].Mobile)
	usertempinfo.Replace("mobileTelephoneNumber", []string{UserListDetailInfo[userid].Mobile})
	usertempinfo.Replace("objectClass", []string{"inetOrgPerson","organizationalPerson"})
	fmt.Println("ModifyUserinfo  done2")
	er := LDAPservice.Conn.Modify(usertempinfo)
	fmt.Println("ModifyUserinfo  done222")
	if er!=nil{
		log.Errorf("modifyuserinfo  %s something error  %v",UserListDetailInfo[userid].Name,er)
		//log.Error(er)
	}else{
		log.Infof("Modify user %s  successfully",UserListDetailInfo[userid].Name)
	}
	fmt.Println("ModifyUserinfo  done3")
}







