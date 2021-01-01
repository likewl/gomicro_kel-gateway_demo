package main

import (
	context "context"
	"fmt"
	Services "rpctest/services"
)
/**
service层
*/
type imp struct{}

//rpc服务器
func (i imp) GetProd(ctx context.Context, request *Services.Request, response *Services.Response) error {
	fmt.Println(request.Size)
	response.Data = &Services.Test{
		Code:    200,
		Message: "ok!",
	}
	return nil
}
