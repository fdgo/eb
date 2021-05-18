package main

import (
	"eb/cart/domain/repository"
	service2 "eb/cart/domain/service"
	"eb/cart/handler"
	proto "eb/cart/proto"
	"eb/common"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	etcd2 "github.com/micro/go-plugins/registry/etcd/v2"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
)

var QPS = 100

func main() {
	//配置中心
	globalcfg, err := common.GetEtcdConfig("120.27.239.127", 2379, "/micro/etcd")
	if err != nil {
		log.Error(err)
	}
	//注册中心
	etcd := etcd2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"120.27.239.127:2379",
		}
	})

	//链路追踪
	t, io, err := common.NewTracer("go.micro.service.cart", "localhost:6831")
	if err != nil {
		log.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//创建数据库连接
	db, err := gorm.Open("mysql", globalcfg.Mysql.User+":"+globalcfg.Mysql.Pwd+"@tcp("+globalcfg.Mysql.Host+")/"+globalcfg.Mysql.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err !=nil {
		log.Error(err)
	}
	defer db.Close()
	//禁止副表
	db.SingularTable(true)

	//第一次初始化
	//err = repository.NewCartRepository(db).InitTable()
	//if err !=nil {
	//	log.Error(err)
	//}


	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.cart"),
		micro.Version("latest"),
		//暴露的服务地址
		micro.Address("0.0.0.0:8077"),
		//注册中心
		micro.Registry(etcd),
		//链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		//添加限流
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
	)


	// Initialise service
	service.Init()

	cartDataService := service2.NewCartDataService(repository.NewCartRepository(db))

	// Register Handler
	proto.RegisterCartHandler(service.Server(), &handler.Cart{CartDataService: cartDataService})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
