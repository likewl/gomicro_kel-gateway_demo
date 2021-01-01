package main

import (
	"flag"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	Services "microtest/service"
)

func main() {
	//设置ctl指令
	var (
		consulHost = flag.String("consul_host", "localhost", "consul server ip address")
		consulPort = flag.String("consul_port", "8500", "consul server port")
		addr       = *consulHost + ":" + *consulPort
	)
	flag.Parse()

	//初始化要注册的consul
	newRegistry := consul.NewRegistry(
		registry.Addrs(addr),
	)
	//相关业务路由初始化
	router := Services.Init()
	//设置micro的rpc服务参数
	service := web.NewService(
		web.Name("test"),
		web.Address(":8001"),
		web.Handler(router),
		web.Registry(newRegistry),
	)
	//初始化service
	service.Init()
	//运行service服务
	service.Run()
}
