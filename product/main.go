package main

import (
	"eb/common"
	"eb/product/domain/repository"
	service2 "eb/product/domain/service"
	"eb/product/handler"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	etcd2 "github.com/micro/go-plugins/registry/etcd/v2"
	"github.com/opentracing/opentracing-go"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	product "eb/product/proto"
)

func main() {
	//配置中心
	etcdConfig,err := common.GetEtcdConfig("120.27.239.127",2379,"/micro/etcd")
	if err !=nil {
		log.Error(err)
	}
	//注册中心
	etcd := etcd2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"120.27.239.127:2379",
		}
	})

	//链路追踪
	t,io,err:=common.NewTracer("go.micro.service.product","localhost:6831")
	if err !=nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//数据库设置
	mysqlInfo := common.GetMysqlFromConsul(etcdConfig,"mysql")
	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@tcp("+mysqlInfo.Host+")/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err !=nil {
		log.Error(err)
	}
	defer db.Close()
	//禁止副表
	db.SingularTable(true)

	//初始化
	//repository.NewProductRepository(db).InitTable()

	productDataService := service2.NewProductDataService(repository.NewProductRepository(db))

	// 设置服务
	service := micro.NewService(
		micro.Name("go.micro.service.product"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8083"),
		//添加注册中心
		micro.Registry(etcd),
		//绑定链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
	)

	// Initialise service
	service.Init()

	// Register Handler
	product.RegisterProductHandler(service.Server(), &handler.Product{ProductDataService:productDataService})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
