package main

import (
	"flag"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	Services "rpctest/services"
)

func main() {
	//设置ctl指令
	var (
		consulHost = flag.String("consul_host", "localhost", "consul server ip address")
		consulPort = flag.String("consul_port", "8500", "consul server port")
		addr = *consulHost +":"+*consulPort
	)
	flag.Parse()

	//初始化要注册的consul
	newRegistry := consul.NewRegistry(
		registry.Addrs(addr),
	)
	//设置micro的rpc服务参数
	service := micro.NewService(
		micro.Name("test1"),
		micro.Address(":8002"),
		micro.Registry(newRegistry),
	)
	//注册相应的rpc路由到service中
	Services.RegisterProdServiceHandler(service.Server(), imp{})
	//初始化service
	service.Init()
	//运行service服务
	service.Run()

}
