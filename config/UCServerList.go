package config

import (
	"io/ioutil"
	"os"

	"github.com/alecthomas/repr"
	"github.com/zxfonline/json"
)

//用户中心配置信息列表
var (
	C_UCServerList map[int16]UCServer
)

//用户中心信息
type UCServer struct {
	Name    string `json:"area_name"`     //大区名称
	AreaId  int16  `json:"area_id"`       //用户中心服务器id(大区id)
	WanAddr string `json:"area_wan_addr"` //用户中心对玩家开放地址
}

func init() {
	data, err := ioutil.ReadFile("./config/UCServerList.ini")
	if err != nil {
		println("server start error:UCServerList json file read failed,", err)
		os.Exit(1)
	}
	err = json.Unmarshal(data, &C_UCServerList)
	if err != nil {
		println("server start error:UCServerList json unmarshal failed:", err)
		os.Exit(1)
	}
	println(repr.String(C_UCServerList, repr.Indent(" ")))
	if len(C_UCServerList) == 0 {
		println("server start error:UCServerList is nil,")
		os.Exit(1)
	}
}
