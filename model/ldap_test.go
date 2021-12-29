package model

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
	"gopkg.in/ldap.v2"
	"testing"
)


var test LDAPService

func TestInitLdapConnection(t *testing.T) {
	Ldapconfig=LdapConfig{
		Addr :"192.168.137.201:389",
		BindUserName:"cn=manager,dc=asinking,dc=com",
		BindPassword :"1qaz@WSX",
		SearchDN :"dc=asinking,dc=com",
	}
	InitLdapConnection()
	//searchRequest := ldap.NewSearchRequest(Ldapconfig.SearchDN,
	//	ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
	//	fmt.Sprintf("(&(objectClass=organizationalPerson)(uid=%s))", "op01"),
	//	[]string{"dn"},
	//	nil,)
	//sr, err := LDAPservice.Conn.Search(searchRequest)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//if len(sr.Entries) != 1 {
	//	log.Fatal("User does not exist or too many entries returned")
	//}
	//sr.Print()
	//InitLdapConnection()
	//LDAPservice.AddGroupinfo()
	//hans := "zhur"
	hans := "朱浩睿我的是"
	convert := pinyin.Convert(hans, nil)
	var fullname string
	var secondname string
 	for _,k:=range convert{
		fullname=fullname+k[0]
	}
	for i:=1;i<len(convert);i++{
		secondname=secondname+convert[i][0]
	}
	firstname:=convert[0][0]
	//sql := ldap.NewAddRequest("uid="+fullname+",ou=it,dc=asinking,dc=com")
	//////sql := ldap.NewAddRequest("ou=供应链组,ou=产品部,ou=Staff,ou=Groups,o=AsinKing,dc=asinking,dc=com")
	////
	//sql.Attribute("uidNumber", []string{"1010"})
	//sql.Attribute("gidNumber", []string{"1003"})
	//sql.Attribute("userPassword", []string{"123456"})
	//sql.Attribute("homeDirectory", []string{"/home/wujq"})
	//sql.Attribute("mail", []string{"test@asinking.com"})
	//sql.Attribute("cn", []string{fullname})
	//sql.Attribute("givenName", []string{firstname})
	//sql.Attribute("sn", []string{secondname})
	//sql.Attribute("mobileTelephoneNumber", []string{"15711823061"})
	////sql.Attribute("uid", []string{"test"})
	//sql.Attribute("title", []string{"运维工程师"})
	////sql.Attribute("objectClass", []string{"inetOrgPerson","shadowAccount", "posixAccount", "account"})
	//sql.Attribute("objectClass", []string{"inetOrgPerson","organizationalPerson","posixAccount"})
	//er := LDAPservice.Conn.Add(sql)
	//if er!=nil{
	//	fmt.Println(er)
	//}




	sql := ldap.NewAddRequest("mail="+fullname+"@asinking.com"+",ou=it,dc=asinking,dc=com")
	sql.Attribute("userPassword", []string{"123456"})
	sql.Attribute("mail", []string{"test@asinking.com"})
	sql.Attribute("cn", []string{fullname})
	sql.Attribute("givenName", []string{firstname})
	sql.Attribute("sn", []string{secondname})
	sql.Attribute("mobileTelephoneNumber", []string{"15711823061"})
	sql.Attribute("uid", []string{"23134234234123412341234123"})
	sql.Attribute("objectClass", []string{"inetOrgPerson","organizationalPerson"})
	er := LDAPservice.Conn.Add(sql)
	if er!=nil{
		fmt.Println(er)
	}

}
//
