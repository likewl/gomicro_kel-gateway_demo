package proxy

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"strings"
)

func Gateway(client *api.Client, logger log.Logger) *httputil.ReverseProxy{
	director :=func(req *http.Request){
		//查询原始请求路径，如：/arithmetic/calculate/10/5
		reqPath := req.URL.Path
		if reqPath == "" {
			return
		}
		//按照分隔符'/'对路径进行分解，获取服务名称serviceName
		pathArray := strings.Split(reqPath, "/")
		serviceName := pathArray[1]

		//调用consul api查询serviceName的服务实例列表
		result, _, err := client.Catalog().Service(serviceName, "", nil)
		if err != nil {
			logger.Println("ReverseProxy failed", "query service instace error", err.Error())

			return
		}

		if len(result) == 0 {
			logger.Println("ReverseProxy failed", "no such service instance", serviceName)
			return
		}

		//重新组织请求路径，去掉服务名称部分
		destPath := strings.Join(pathArray[2:], "/")

		//随机选择一个服务实例
		tgt := result[rand.Int()%len(result)]
		logger.Println("service id", tgt.ServiceID)

		//设置代理服务地址信息
		req.URL.Scheme = "http"
		req.URL.Host = fmt.Sprintf("%s:%d", tgt.ServiceAddress, tgt.ServicePort)
		req.URL.Path = "/" + destPath
	}
	return &httputil.ReverseProxy{
		Director: director,
	}
}