package main

import (
	"context"
	"fmt"
	blogpb "github.com/pradeepitm12/cb/bee/api/gen"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := blogpb.NewBlogPostServiceClient(conn)
	ctx := context.Background()

	// For Create post
	createReq := &blogpb.CreateRequest{
		Title:   "My First Blog",
		Content: "Some content",
		Author:  "ABC",
		Tags:    []string{"golang", "grpc"},
	}
	createResp, err := client.CreatePost(ctx, createReq)
	if err != nil {
		log.Fatalf("CreatePost failed: %v", err)
	}
	postID := createResp.Post.PostID
	fmt.Printf("Created Post: %+v\n", createResp.Post)

	// To Read a post
	readResp, err := client.ReadPost(ctx, &blogpb.ReadRequest{PostID: postID})
	if err != nil {
		log.Fatalf("ReadPost failed: %v", err)
	}
	fmt.Printf("Read Post: %+v\n", readResp.Post)

	// To update a post
	updateReq := &blogpb.UpdateRequest{
		Id:      postID,
		Title:   "My First Blog (Updatedd)",
		Content: "Updated content",
		Author:  "ABC",
		Tags:    []string{"golang", "grpc", "updated"},
	}
	updateResp, err := client.UpdatePost(ctx, updateReq)
	if err != nil {
		log.Fatalf("UpdatePost failed: %v", err)
	}
	fmt.Printf("Updated Post: %+v\n", updateResp.Post)

	// To delete a post
	_, err = client.DeletePost(ctx, &blogpb.DeleteRequest{PostID: postID})
	if err != nil {
		log.Fatalf("DeletePost failed: %v", err)
	}
	fmt.Println("Post deleted successfully")

	// Try to Read Deleted Post
	_, err = client.ReadPost(ctx, &blogpb.ReadRequest{PostID: postID})
	if err != nil {
		fmt.Printf("As expected, post not found after delete: %v\n", err)
	}
}
