package model

import (
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
	LDAPservice.AddGroupinfo()

	//sql := ldap.NewAddRequest("uid=test,ou=people,ou=it,dc=asinking,dc=com")
	////sql := ldap.NewAddRequest("ou=供应链组,ou=产品部,ou=Staff,ou=Groups,o=AsinKing,dc=asinking,dc=com")
	//
	//sql.Attribute("uidNumber", []string{"1010"})
	//sql.Attribute("gidNumber", []string{"1003"})
	//sql.Attribute("userPassword", []string{"123456"})
	//sql.Attribute("homeDirectory", []string{"/home/wujq"})
	//sql.Attribute("cn", []string{"test"})
	//sql.Attribute("uid", []string{"test"})
	//sql.Attribute("objectClass", []string{"shadowAccount", "posixAccount", "account"})
	////sql.Attribute("objectClass", []string{"inetOrgPerson","organizationalPerson","person","top"})
	//er := LDAPservice.Conn.Add(sql)
	//if er!=nil{
	//	fmt.Println(er)
	//}

}
//
