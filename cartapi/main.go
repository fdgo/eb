package main

import (
	"context"
	go_micro_service_cart "eb/cart/proto"
	"eb/cartapi/handler"
	"eb/common"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	etcd2 "github.com/micro/go-plugins/registry/etcd/v2"
	"github.com/micro/go-plugins/wrapper/select/roundrobin/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"net"
	"net/http"

	cartApi "eb/cartapi/proto"
)
var (
	ETCD_ADDRESS       = "120.27.239.127"
	ETCD_PORT          = "2379"
	ETCD_CONFIG_PREFIX = "/micro/etcd"
	ETCD_GATEWAY_TOKEN = "gateway_token"
)
func main() {
	//注册中心
	etcd := etcd2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"120.27.239.127:2379",
		}
	})
	t,io,err := common.NewTracer("go.micro.api.cartApi","localhost:6831")
	if err != nil {
		log.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//熔断器
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	//启动端口
	go func() {
		err = http.ListenAndServe(net.JoinHostPort("192.168.204.133","8989"),hystrixStreamHandler)
		if err !=nil {
			log.Error(err)
		}
	}()


	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.cartApi"),
		micro.Version("latest"),
		micro.Address("0.0.0.0:8086"),
		//添加 consul 注册中心
		micro.Registry(etcd),
		//添加链路追踪
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
		//添加熔断
		micro.WrapClient(NewClientHystrixWrapper()),
		//添加负载均衡
		micro.WrapClient(roundrobin.NewClientWrapper()),
	)

	// Initialise service
	service.Init()

	cartService:=go_micro_service_cart.NewCartService("go.micro.service.cart",service.Client())

	cartService.AddCart(context.TODO(),&go_micro_service_cart.CartInfo{

		UserId:    3,
		ProductId: 4,
		SizeId:    5,
		Num:       5,
	})

	// Register Handler
	if err := cartApi.RegisterCartApiHandler(service.Server(), &handler.CartApi{CartService:cartService});err !=nil {
		log.Error(err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

type clientWrapper struct {
	client.Client
}

func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	return hystrix.Do(req.Service()+"."+req.Endpoint(), func() error {
		//run 正常执行
		fmt.Println(req.Service()+"."+req.Endpoint())
		return c.Client.Call(ctx,req,rsp,opts...)
	}, func(err error) error {
		fmt.Println(err)
		return err
	})
}

func NewClientHystrixWrapper() client.Wrapper  {
	return func(i client.Client) client.Client {
		return &clientWrapper{i}
	}
}



