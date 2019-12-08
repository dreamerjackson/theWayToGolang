package main
//
//import (
//	"context"
//	"fmt"
//	"github.com/micro/go-micro"
//	"github.com/micro/go-micro/registry"
//	"github.com/micro/go-micro/registry/consul"
//	"github.com/micro/go-micro/service/grpc"
//	go_micro_srv_greeter "jonson/grpc-study/myproto"
//)
//
//
//
//type Say struct {
//
//}
//
//func (h Say) Hello(ctx context.Context, in *go_micro_srv_greeter.Request, out*go_micro_srv_greeter.Response) error {
//	out.Msg = "olaya"
//	return nil
//}
//
////func middleware(fn server.HandlerFunc) server.HandlerFunc{
////
////return func(ctx context.Context, req server.Request, rsp interface{}) error{
////	fmt.Println("hello this is middleware")
////	return fn(ctx,req,rsp)
////}}
//
//func main(){
//	reg:= consul.NewRegistry(
//		func(options *registry.Options) {
//					options.Addrs = []string{"127.0.0.1:8500"}
//		})
//	service := grpc.NewService(
//		micro.Registry(reg),
//		micro.Name("hello-51cto"),
//		//micro.WrapHandler(middleware),
//		)
//
//	 grpcserver := service.Server()
//	  myservice := Say{}
//	err:=  grpcserver.Handle(grpcserver.NewHandler(&myservice))
//	if err!=nil{
//		fmt.Println(err)
//	}
//	err = service.Run()
//	if err!=nil{
//		fmt.Println(err)
//	}
//}