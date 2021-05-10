package model

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
)

var(
	DepListId             =make(map[int][]int)                  //department 子列表信息
	DepListDetailInfo     =make(map[int]DepDetailInfo)          //department 详细信息
	ChanageDepDetailch    chan *ChanageDepDetailInfo
	StackDepmentinfo    []string
)

type ResponseListSubInfo struct {
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
//department 详细信息
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

type ChanageDepDetailInfo struct{
	Action string
	Departinfo *DepDetailInfo
}

//department接口返回信息
type ResponseDepDetailInter struct {
	Errcode   int           `json:"errcode"`
	Errmsg    string        `json:"errmsg"`
	Result    DepDetailInfo `json:"result"`
	RequestId string        `json:"request_id"`
}

//department 子列表信息
type ResponseDepListSubId struct {
	Errcode int `json:"errcode"`
	Result  struct {
		DeptIdList []int `json:"dept_id_list"`
	} `json:"result"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"request_id"`
}




//初始化组织信息的结构内存map信息,   待补充并发初始化数据
func InitListSubId(method string,DepID int, url string) {
	body:= make(map[string]interface{})
	body["dept_id"]=DepID
	str := UrlRequest(method, url,&body)
	json_info := ResponseDepListSubId{}
	err := json.Unmarshal(str, &json_info)
	if err != nil {
		log.Println(err)
	}
	DepListId[DepID]=json_info.Result.DeptIdList //更新部门子部门map信息
	info, err := GetSubDetailInfo("POST", DepID)
	if err!=nil{
		log.Println(err)
		//return
	}
	if _,ok:=DepListDetailInfo[DepID];!ok{
		StackDepmentinfo= append(StackDepmentinfo, info.Result.Name)
		//fmt.Println(StackDepmentinfo)
		DepListDetailInfo[DepID] = info.Result
		LDAPservice.AddGroupinfo(DepListDetailInfo[DepID])
		//temp:=&ChanageDepDetailInfo{"Add",DepListDetailInfo[DepID]}
		//ChanageDepDetailch<-temp
	}
	//更新部门部门详细信息到map
	for _,v:=range json_info.Result.DeptIdList{
		InitListSubId(method,v,url)
		StackDepmentinfo= StackDepmentinfo[:len(StackDepmentinfo)-1]
	}
	//fmt.Println("stack info",StackDepmentinfo)
}

//func Add





//获取列表信息
func GetListSubId(method string,DepID int, url string) {
	body:= make(map[string]interface{})
	body["dept_id"]=DepID
	str := UrlRequest(method, url,&body)
	//body_json := err2
	json_info := ResponseDepListSubId{}
	err := json.Unmarshal(str, &json_info)
	if err != nil {
		log.Println(err)
	}
}




func GetListSub(method string, url string) []SubInfo {
	str := UrlRequest(method, url, nil)
	json_info := ResponseListSubInfo{}
	err := json.Unmarshal(str, &json_info)
	if err != nil {
		log.Println(err)
		log.Println(err)
	}
	//token:=json_info.Access_token
	//DepartmentInfo:=
	return json_info.Department
}

func GetSubDetailInfo(method string,DepID int) (ResponseDepDetailInter,error) {
	UrlDepDetail := "https://oapi.dingtalk.com/topapi/v2/department/get?access_token=" + Token + "&dept_id=" + strconv.Itoa(DepID)
	str := UrlRequest(method, UrlDepDetail, nil)
	json_info := ResponseDepDetailInter{}
	err := json.Unmarshal(str, &json_info)
	if err != nil {
		return json_info,err
	}
	if json_info.Errcode!=0{
		log.Println("getting department detail information is something wrong",DepID)
		return json_info,errors.New("getting department detail information is something wrong"+strconv.Itoa(DepID))
	}
	//token:=json_info.Access_token
	//DepartmentInfo:=
	return json_info,nil
}
