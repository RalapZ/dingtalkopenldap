package model

var (
	Token string
	METHOD = "get"
)

type Tokenstr struct {
	Errcode      int    `json:"errocd"`
	Access_token string `json:"access_token"`
	Errmsg       string `json:"errmsg"`
	Expires_in   int    `json:"expires_in"`
	//info []interface{}
}

type ListSubInfo struct {
	Errcode    int       `json:"errcode"`
	Errmsg     string    `json:"errmsg"`
	Department []SubInfo `json:"department"`
}

type SubInfo struct {
	CreateDeptGroup bool   `json:"createDeptGroup"`
	Name            string `json:"name"`
	Id              int    `json:"id"`
	AutoAddUser     bool   `json:"autoAddUser"`
	Parentid        int    `json:"parentid"`
}

type DepDetailInfo struct {
	AutoAddUser         bool          `json:"auto_add_user"`
	CreateDeptGroup     bool          `json:"create_dept_group"`
	DeptGroupChatId     string        `json:"dept_group_chat_id"`
	DeptId              int           `json:"dept_id"`
	DeptPermits         []interface{} `json:"dept_permits"`
	GroupContainSubDept bool          `json:"group_contain_sub_dept"`
	HideDept            bool          `json:"hide_dept"`
	Name                string        `json:"name"`
	Order               int           `json:"order"`
	OrgDeptOwner        string        `json:"org_dept_owner"`
	OuterDept           bool          `json:"outer_dept"`
	OuterPermitDepts    []interface{} `json:"outer_permit_depts"`
	OuterPermitUsers    []interface{} `json:"outer_permit_users"`
	ParentId            int           `json:"parent_id"`
	UserPermits         []interface{} `json:"user_permits"`
}
type ResponseDepDetailInter struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Result  DepDetailInfo `json:"result"`
	RequestId string `json:"request_id"`
}
