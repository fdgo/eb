package common

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/etcd"
	"strconv"
)

type GlobalCfg struct {
	Mysql struct {
		Host     string `json:"host"`
		User     string `json:"user"`
		Pwd      string `json:"pwd"`
		Database string `json:"database"`
		Port     string `json:"port"`
	} `json:"mysql:"`
	GatewayToken string `json:"gateway_token"`
}

//设置配置中心
func GetEtcdConfig(host string, port int64, prefix string) (GlobalCfg, error) {
	//添加配置中心
	//配置中心使用consul key/value 模式
	etcdSource := etcd.NewSource(
		//设置配置中心地址
		etcd.WithAddress(host+":"+strconv.FormatInt(port, 10)),
		//设置前缀，不设置默认为 /micro/etcd
		etcd.WithPrefix(prefix),
		//是否移除前缀，这里设置为true 表示可以不带前缀直接获取对应配置
		etcd.StripPrefix(true),
	)
	err := config.Load(etcdSource)
	if err != nil {
		panic(err)
	}
	var conf GlobalCfg
	err = config.Scan(&conf)
	return conf, nil
}
