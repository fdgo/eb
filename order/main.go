package main

import (
	"eb/common"
	"eb/order/domain/repository"
	service2 "eb/order/domain/service"
	"eb/order/handler"
	order "eb/order/proto"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	etcd2 "github.com/micro/go-plugins/registry/etcd/v2"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
)

var (
	//qps = os.Getenv("QPS")
	QPS = 1000
)

func main() {
	// 1.配置中心
	etcdConfig,err:=common.GetEtcdConfig("localhost",2379,"/micro/etcd")
	if err !=nil {
		log.Error(err)
	}
	// 2.注册中心
	etcd :=etcd2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"120.27.239.127:2379",
		}
	})
	// 3.jaeger 链路追踪
	t,io,err := common.NewTracer("go.micro.service.order","localhost:6831")
	if err !=nil {
		log.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)
	// 4.初始化数据库
	mysqlInfo := common.GetMysqlFromConsul(etcdConfig,"mysql")
	db,err:= gorm.Open("mysql",mysqlInfo.User+":"+mysqlInfo.Pwd+"@/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err !=nil {
		log.Error(err)
	}
	defer db.Close()
	//禁止副表
	db.SingularTable(true)

	//第一次运行的时候创建表
	//tableInit := repository.NewOrderRepository(db)
	//tableInit.InitTable()

	//创建实例
	orderDataService := service2.NewOrderDataService(repository.NewOrderRepository(db))

	// 5.暴露监控地址
	common.PrometheusBoot(9092)

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.order"),
		micro.Version("latest"),
		//暴露的服务地址
		micro.Address(":9085"),
		//添加consul 注册中心
		micro.Registry(etcd),
		//添加链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		//添加限流
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
		//添加监控
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
	)

	// Initialise service
	service.Init()

	// Register Handler
	order.RegisterOrderHandler(service.Server(), &handler.Order{OrderDataService:orderDataService})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
