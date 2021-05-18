package main

import (
	"eb/common"
	"eb/payment/domain/repository"
	service2 "eb/payment/domain/service"
	"eb/payment/handler"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	etcd2 "github.com/micro/go-plugins/registry/etcd/v2"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	payment "eb/payment/proto"
)

func main() {
	// 配置中心
	etcdConfig,err := common.GetEtcdConfig("localhost",2379,"/micro/etcd")
	if err !=nil {
		common.Error(err)
	}
	//注册中心
	etcd := etcd2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"120.27.239.127:2379",
		}
	})
	//jaeger 链路追踪
	t,io,err:=common.NewTracer("go.micro.service.payment","localhost:6831")
	if err!=nil{
		common.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//mysql 设置
	mysqlInfo := common.GetMysqlFromConsul(etcdConfig,"mysql")
	//初始化数据库
	db,err := gorm.Open("mysql" , mysqlInfo.User+":"+mysqlInfo.Pwd+"@/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err !=nil {
		common.Error(err)
	}
	defer db.Close()
	//禁止复数表
	db.SingularTable(true)

	//创建表
	tableInit := repository.NewPaymentRepository(db)
	tableInit.InitTable()

	//监控
	common.PrometheusBoot(9089)


	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.payment"),
		micro.Version("latest"),
		micro.Address("0.0.0.0:8089"),
		//添加注册中心
		micro.Registry(etcd),
		//添加链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		//加载限流
		micro.WrapHandler(ratelimit.NewHandlerWrapper(1000)),
		//加载监控
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
	)

	// Initialise service
	service.Init()

	paymentDataService := service2.NewPaymentDataService(repository.NewPaymentRepository(db))

	// Register Handler
	payment.RegisterPaymentHandler(service.Server(), &handler.Payment{PaymentDataService:paymentDataService})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
