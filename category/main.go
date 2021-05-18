package main

import (
	"eb/category/domain/repository"
	service2 "eb/category/domain/service"
	"eb/category/handler"
	category "eb/category/proto"
	"eb/common"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	etcd2 "github.com/micro/go-plugins/registry/etcd/v2"
)

func main() {
	//配置中心
	etcdConfig, err := common.GetEtcdConfig("120.27.239.127", 2379, "/micro/etcd")
	if err != nil {
		log.Error(err)
	}
	//注册中心
	etcd := etcd2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"120.27.239.127:2379",
		}
	})

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.category"),
		micro.Version("latest"),
		//这里设置地址和需要暴露的端口
		micro.Address("127.0.0.1:8082"),
		//添加consul 作为注册中心
		micro.Registry(etcd),
	)

	//获取mysql配置,路径中不带前缀
	mysqlInfo := common.GetMysqlFromConsul(etcdConfig, "mysql")

	//连接数据库
	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@tcp("+mysqlInfo.Host+")/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Error(err)
	}
	defer db.Close()
	//禁止复表
	db.SingularTable(true)

	//rp := repository.NewCategoryRepository(db)
	//rp.InitTable()
	// Initialise service
	service.Init()

	categoryDataService := service2.NewCategoryDataService(repository.NewCategoryRepository(db))

	err = category.RegisterCategoryHandler(service.Server(), &handler.Category{CategoryDataService: categoryDataService})
	if err != nil {
		log.Error(err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
