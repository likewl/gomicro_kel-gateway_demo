package main

import (
	"flag"
	"fmt"
	"github.com/hashicorp/consul/api"
	proxy2 "kel-gateway/proxy"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 创建环境变量
	var (
		consulHost = flag.String("consul_host", "localhost", "consul server ip address")
		consulPort = flag.String("consul_port", "8500", "consul server port")
	)
	flag.Parse()

	//创建日志组件
	var logger log.Logger
	{
		logger = *log.New(os.Stderr, "kel", 1)

	}

	// 创建consul api客户端
	consulConfig := api.DefaultConfig()
	consulConfig.Address = "http://" + *consulHost + ":" + *consulPort
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		logger.Println("err", err)
		os.Exit(1)
	}

	//创建反向代理
	proxy := proxy2.Gateway(consulClient, logger)

	errc := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	//开始监听
	go func() {
		logger.Println("transport", "HTTP", "addr", "9090")
		err2 := http.ListenAndServe(":9090", proxy)
		fmt.Println("123")
		errc <- err2

	}()

	// 开始运行，等待结束
	logger.Println("exit", <-errc)

}
