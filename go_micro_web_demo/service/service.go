package Services

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
)
/*
路由相关
*/
func Init() *gin.Engine{
	r := gin.Default()
	//普通handle
	r.GET("/a", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "ok",
		})
	})
	//内嵌rpc客户端
	r.GET("/b", func(c *gin.Context) {
		service := NewProdService("test1", micro.NewService().Client())
		res :=new(Request)
		res.Size=10
		prod, err := service.GetProd(context.Background(),res)
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(200,gin.H{
			"data": prod.Data,
		})
	})
	return r
}
