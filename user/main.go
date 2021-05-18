package main

import (
	"eb/user/domain/repository"
	service2 "eb/user/domain/service"
	"eb/user/handler"
	"eb/user/proto"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
)

func main() {
	//服务参数设置
	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
	)
	//初始化服务
	srv.Init()

	//创建数据库连接
	db,err :=gorm.Open("mysql","root:000000@tcp(120.27.239.127)/gomicro?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.SingularTable(true)
	//只执行一次，数据表初始化
	//rp:=repository.NewUserRepository(db)
	//rp.InitTable()
	//创建服务实例
	userDataService := service2.NewUserDataService(repository.NewUserRepository(db))
	//注册Handler
	err = proto.RegisterUserHandler(srv.Server(),&handler.User{UserDataService: userDataService})
	if err != nil {
		fmt.Println(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
