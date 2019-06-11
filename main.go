package main

import (
	"net"
	"log"
	"google.golang.org/grpc"

	pb "spiderProxy/interval/serve/grpc"
	"spiderProxy/interval/dao"
)

func main() {
	lis, err := net.Listen("tcp", ":2333")
	log.Printf("listen: 2333")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBookServer(s, &dao.Server{})
	if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
	}
}
