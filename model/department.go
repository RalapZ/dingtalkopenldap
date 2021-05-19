package model

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"time"
)

var (
	DepListId          = make(map[int][]int)         //department 子列表信息
	DepListDetailInfo  = make(map[int]DepDetailInfo) //department 详细信息
	ChanageDepDetailch chan *ChanageDepDetailInfo
	StackDepmentinfo   []string
)

type ResponseListSubInfo struct {
	Errcode    int       `json:"errcode"`
	Errmsg     string    `json:"errmsg"`
	Department []SubInfo `json:"department"`
}

type SubInfo struct {
	CreateDeptGroup bool   `json:"createDeptGroup"` //是否同步创建一个关联此部门的企业群： true：创建false：不创建
	Name            string `json:"name"`            //部门名称
	Id              int    `json:"id"`              //部门ID
	AutoAddUser     bool   `json:"autoAddUser"`     //部门群已经创建后，有新人加入部门是否会自动加入该群： true：会自动入群false：不会
	Parentid        int    `json:"parentid"`        //父部门ID。
}

//department 详细信息
type DepDetailInfo struct {
	AutoAddUser         bool          `json:"auto_add_user"`          //当部门群已经创建后，是否有新人加入部门会自动加入该群： true：自动加入群 false：不会自动加入群
	CreateDeptGroup     bool          `json:"create_dept_group"`      //是否同步创建一个关联此部门的企业群：  true：创建 false：不创建
	DeptGroupChatId     string        `json:"dept_group_chat_id"`     //部门群ID
	DeptId              int           `json:"dept_id"`                //部门ID
	DeptPermits         []interface{} `json:"dept_permits"`           //当隐藏本部门时（即hide_dept为true时），配置的允许在通讯录中查看本部门的部门列表。
	GroupContainSubDept bool          `json:"group_contain_sub_dept"` //部门群是否包含子部门： true：包含false：不包含
	HideDept            bool          `json:"hide_dept"`              //是否隐藏本部门： true：隐藏部门，隐藏后本部门将不会显示在公司通讯录中 false：显示部门
	Name                string        `json:"name"`                   //部门名称
	Order               int           `json:"order"`                  //在父部门中的次序值
	OrgDeptOwner        string        `json:"org_dept_owner"`         //企业群群主ID。
	OuterDept           bool          `json:"outer_dept"`             // 是否限制本部门成员查看通讯录：true：开启限制。开启后本部门成员只能看到限定范围内的通讯录false：不限制
	OuterPermitDepts    []interface{} `json:"outer_permit_depts"`     // 当限制部门成员的通讯录查看范围时（即outer_dept为true时），配置的部门员工可见部门列表。
	OuterPermitUsers    []interface{} `json:"outer_permit_users"`     //当限制部门成员的通讯录查看范围时（即outer_dept为true时），配置的部门员工可见员工列表。
	ParentId            int           `json:"parent_id"`              //父部门ID
	UserPermits         []interface{} `json:"user_permits"`           //当隐藏本部门时（即hide_dept为true时），配置的允许在通讯录中查看本部门的员工列表。
	LdapDepPath         []string      //记录部门的group path
}

type ChanageDepDetailInfo struct {
	Action     string
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
func InitListSubId(method string, DepID int, url string) {
	body := make(map[string]interface{})
	body["dept_id"] = DepID
	str := UrlRequest(method, url, &body)
	json_info := ResponseDepListSubId{}
	err := json.Unmarshal(str, &json_info)
	if err != nil {
		log.Println(err)
	}
	DepListId[DepID] = json_info.Result.DeptIdList //更新部门子部门map信息
	info, err := GetSubDetailInfo("POST", DepID)
	if err != nil {
		log.Println(err)
	}
	if _, ok := DepListDetailInfo[DepID]; !ok {
		StackDepmentinfo = append(StackDepmentinfo, info.Result.Name)
		info.Result.LdapDepPath = StackDepmentinfo
		DepListDetailInfo[DepID] = info.Result
		LDAPservice.AddGroupinfo(DepListDetailInfo[DepID])
	}
	//更新部门部门详细信息到map
	for _, v := range json_info.Result.DeptIdList {
		InitListSubId(method, v, url)
		StackDepmentinfo = StackDepmentinfo[:len(StackDepmentinfo)-1]
	}
}

//获取列表信息
func GetListSubId(method string, DepID int, url string) {
	body := make(map[string]interface{})
	body["dept_id"] = DepID
	str := UrlRequest(method, url, &body)
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
	}
	return json_info.Department
}

func GetSubDetailInfo(method string, DepID int) (ResponseDepDetailInter, error) {
	UrlDepDetail := GetDepDetailUrl + "?access_token=" + Token + "&dept_id=" + strconv.Itoa(DepID)
	str := UrlRequest(method, UrlDepDetail, nil)
	json_info := ResponseDepDetailInter{}
	err := json.Unmarshal(str, &json_info)
	if err != nil {
		return json_info, err
	}
	if json_info.Errcode != 0 {
		log.Println("getting department detail information is something wrong", DepID)
		return json_info, errors.New("getting department detail information is something wrong" + strconv.Itoa(DepID))
	}
	return json_info, nil
}

//对比group信息的parent路径信息
func GroupCompareInfo(srcinfo DepDetailInfo, dstinfo DepDetailInfo) bool {
	if len(srcinfo.LdapDepPath) != len(dstinfo.LdapDepPath) {
		return false
	}
	for i := 0; i < len(srcinfo.LdapDepPath); i++ {
		if srcinfo.LdapDepPath[i] != dstinfo.LdapDepPath[i] {
			return false
		}
	}
	return true
}

func UpdataDepListIdAndDepListDetailInfo(){
}

func ScheduleUpdateSub(){
	for{
		select {
		case <-time.After(time.Duration(Defaultconfig.SchedulerTime)):
			UpdataDepListIdAndDepListDetailInfo()
		}
	}
}
