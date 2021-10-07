package main

import (
	"fmt"
	"google.golang.org/grpc"
	"io/ioutil"
	"net"
	"net/http"
	pb "proxy/service/proxy/proxy/proxy"

	"context"
)

type server struct{

}

func(s *server) Get(ctx context.Context, in *pb.Url) (*pb.Answer, error) {
	fmt.Println(in)

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
	lis, err := net.Listen("tcp", ":9999")

	if err != nil {
		fmt.Println(err)
	}

	s := grpc.NewServer()
	pb.RegisterProxyServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		fmt.Println(err)
	}
}
