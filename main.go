package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io/ioutil"
	"net"
	"net/http"
	pb "proxy/service/proxy"
)

type server struct{

}

const (
	address = ":9999"
)

func(s *server) Get(ctx context.Context, in *pb.Url) (*pb.Answer, error) {
	fmt.Printf("Recive request: %s\n", in.Value)

	resp, err := http.Get(in.Value)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &pb.Answer{Value: string(body)}, nil
}

func main()  {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println(err)
	}

	s := grpc.NewServer()

	pb.RegisterProxyServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		fmt.Println(err)
	}
}
