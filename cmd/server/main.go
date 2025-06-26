package main

import (
	"log"
	"net"

	blogpb "github.com/pradeepitm12/cb/bee/api/gen"
	"github.com/pradeepitm12/cb/bee/internal/handler"
	"github.com/pradeepitm12/cb/bee/internal/repository/inmemory"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	inMemoryPostStore := inmemory.NewPostStore()
	hand := handler.NewPostHandler(inMemoryPostStore)
	grpcServer := grpc.NewServer()

	blogpb.RegisterBlogPostServiceServer(grpcServer, hand)

	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("gRPC server is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
