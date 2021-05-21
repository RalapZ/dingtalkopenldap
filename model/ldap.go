package model

import (
	"fmt"
	//"github.com/RalapZ/dingtalkopenldap/tools"
	"github.com/mozillazg/go-pinyin"
	log "github.com/sirupsen/logrus"
	"os"

	//"github.com/go-ldap/ldap/v3"
	"gopkg.in/ldap.v2"
)

//var LdapConn *ldap.Conn

var LDAPservice LDAPService

type LDAPService struct {
	Conn     *ldap.Conn
	Config     LdapConfig
}

func InitLdapConnection(){
	conn, err := ldap.Dial("tcp", Ldapconfig.Addr)
	if err!=nil{
		log.Error("Ldap connected port error ")

		os.Exit(3)
	}
	log.Info("Ldap connected port successfully")
	LDAPservice=LDAPService{conn,Ldapconfig}
	err = LDAPservice.Conn.Bind(Ldapconfig.BindUserName, Ldapconfig.BindPassword)
	if err !=nil{
		fmt.Println(err)
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







