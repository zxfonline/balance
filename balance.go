package balance

import (
	"github.com/zxfonline/balance/config"

	"strconv"

	"github.com/zxfonline/strutil"

	"github.com/stathat/consistent"
)

var (
	c         = consistent.New()
	ucservers map[int16]config.UCServer
)

func init() {
	ucservers = config.C_UCServerList
	c.NumberOfReplicas = len(ucservers)
	for _, server := range ucservers {
		c.Add(strconv.Itoa(int(server.AreaId)))
	}
}

//通过玩家唯一id获取进入的用户中心配置信息
func GetUCServerById(userid int64) config.UCServer {
	server, _ := c.Get(strconv.Itoa(int(userid)))
	return ucservers[strutil.Stoi16(server)]
}
