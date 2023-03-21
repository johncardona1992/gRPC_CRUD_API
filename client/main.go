package main

import (
	"log"
	pb "zelator/grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	id := CreateBlog(c)

	ReadBlog(c, id)
	// ReadBlog(c, "aNonExistingId")
	UpdateBlog(c, id)
	listBlog(c)
	deleteBlog(c, id)
}
