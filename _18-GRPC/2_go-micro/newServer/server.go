package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-plugins/registry/consul"
	"io"
	greetpb "jonson/grpc-study/diffproto"
	"log"
)



type GreetService struct {}

func (g *GreetService) LongGreet(ctx context.Context, stream greetpb.GreetService_LongGreetStream) error {
	fmt.Printf("LongGreet function was invoked with a streaming request\n")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("1123")
			// we have finished reading the client stream
			 err = stream.SendMsg(&greetpb.LongGreetResponse{
				Result: result,
			})
			stream.Close()
			return err

		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		firstName := req.GetGreeting().GetFirstName()
		result += "Hello " + firstName + "! "
	}
}

func (g *GreetService) GreetEveryone(ctx context.Context, stream greetpb.GreetService_GreetEveryoneStream) error {
	fmt.Printf("GreetEveryone function was invoked with a streaming request\n")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}
		firstName := req.GetGreeting().GetFirstName()
		result := "Hello " + firstName + "! "

		sendErr := stream.Send(&greetpb.GreetEveryoneResponse{
			Result: result,
		})
		if sendErr != nil {
			log.Fatalf("Error while sending data to client: %v", sendErr)
			return sendErr
		}
	}

}

func (g *GreetService) Greet( ctx context.Context, req*greetpb.GreetRequest, res*greetpb.GreetResponse) error {
		fmt.Printf("Greet function was invoked with %v\n", req)
		firstName := req.GetGreeting().GetFirstName()
		result := "Hello " + firstName
		res.Result = result
		return nil
}

func (g *GreetService) GreetManyTimes(ctx context.Context, req*greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesStream) error {
		fmt.Printf("LongGreet function was invoked with a streaming request\n")
	 log.Printf("Got msg %v", req)
	for i := 0; i < int(10); i++ {
		if err := stream.Send(&greetpb.GreetManytimesResponse{Result: "kk"}); err != nil {
			return err
		}
	}
	return nil
}

func main(){
	reg:= consul.NewRegistry(
		func(options *registry.Options) {
			options.Addrs = []string{"127.0.0.1:8500"}
		})
	service := grpc.NewService(
		micro.Registry(reg),
		micro.Name("hello-51cto"),
	)

	grpcserver := service.Server()
	myservice := GreetService{}
	// 只能用这种方式
	err:= greetpb.RegisterGreetServiceHandler(grpcserver,&myservice)


	if err!=nil{
		fmt.Println(err)
	}
	err = service.Run()
	if err!=nil{
		fmt.Println(err)
	}
}

