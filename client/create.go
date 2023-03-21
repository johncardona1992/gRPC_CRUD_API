package main

import (
	"context"
	"log"
	pb "zelator/grpc/proto"
)

func CreateBlog(c pb.BlogServiceClient) string {
	log.Println("---createBlog was invoked---")
	blog := &pb.Blog{
		AuthorId: "Zelator",
		Title:    "My first blog.",
		Content:  "Content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %s\n", res.Id)

	return res.Id
}
