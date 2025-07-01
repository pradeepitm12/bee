package handler

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	blogpb "github.com/pradeepitm12/cb/bee/api/gen"
	"github.com/pradeepitm12/cb/bee/internal/errors"
	"github.com/pradeepitm12/cb/bee/internal/model"
	"github.com/pradeepitm12/cb/bee/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PostHandler struct {
	repo repository.PostRepository
	blogpb.UnimplementedBlogPostServiceServer
}

func NewPostHandler(repo repository.PostRepository) *PostHandler {
	return &PostHandler{repo: repo}
}

func (ph *PostHandler) CreatePost(ctx context.Context, req *blogpb.CreateRequest) (*blogpb.CreateResponse, error) {
	log.Printf("Received CreatePost request - Title: %s, Author: %s", req.Title, req.Author)
	newPost := model.NewPost(req.Title, req.Content, req.Author, req.Tags)
	post, err := ph.repo.Create(&ctx, newPost)
	if err != nil && strings.Contains(err.Error(), errors.PostAlreadyExists) {
		log.Printf("Error: Post already exists - Title: %s, Author: %s", newPost.Title, newPost.Author)
		return nil, status.Errorf(codes.AlreadyExists, "Post already exists")
	}
	resp := &blogpb.CreateResponse{
		Post:  transformToPost(post),
		Error: "",
	}
	log.Printf("Created new post with ID: %s", post.ID)
	return resp, nil
}

func (ph *PostHandler) ReadPost(ctx context.Context, req *blogpb.ReadRequest) (*blogpb.ReadResponse, error) {
	log.Printf("Received ReadPost request for ID: %s", req.PostID)
	post, err := ph.repo.Read(&ctx, req.PostID)
	if err != nil {
		log.Printf("Error: Post not found: %s", req.PostID)
		return nil, status.Error(codes.NotFound, "post not found")
	}
	resp := &blogpb.ReadResponse{
		Post:  transformToPost(post),
		Error: "",
	}
	log.Printf("The post: %+v", post)
	return resp, nil
}

func (ph *PostHandler) UpdatePost(ctx context.Context, req *blogpb.UpdateRequest) (*blogpb.UpdateResponse, error) {
	log.Printf("Received UpdatePost request for ID: %s", req.Id)
	post, err := ph.repo.Update(&ctx, req.Id, req.Title, req.Content, req.Author, req.Tags, time.Now())
	if err != nil {
		log.Printf("Error: Post not found: %s", req.Id)
		return nil, status.Errorf(codes.NotFound, "post not found")
	}
	resp := &blogpb.UpdateResponse{Post: transformToPost(post), Error: ""}
	log.Printf("Updated post with ID: %s", post.ID)
	return resp, nil
}

func (ph *PostHandler) DeletePost(ctx context.Context, req *blogpb.DeleteRequest) (*blogpb.DeleteResponse, error) {
	log.Printf("Received DeletePost request for ID: %s", req.PostID)
	msg := ph.repo.Delete(&ctx, req.PostID)
	if msg != "" && msg == errors.PostNotFound {
		log.Printf("Error: Post not found: %s", req.PostID)
		return nil, status.Error(codes.NotFound, fmt.Sprintf(errors.PostNotFound+": %s", req.PostID))
	}
	resp := &blogpb.DeleteResponse{Message: fmt.Sprintf("Post %s Deleted", req.PostID)}
	log.Printf("Deleted post with ID: %s", req.PostID)
	return resp, nil
}

func (ph *PostHandler) ListPost(ctx context.Context, req *blogpb.ListRequest) (*blogpb.ListResponse, error) {
	log.Printf("Received ListPost Request")
	posts := ph.repo.List(&ctx)
	resp := &blogpb.ListResponse{
		Post: transformToPosts(posts),
	}
	return resp, nil
}

func transformToPosts(posts []*model.Post) []*blogpb.Post {
	transformedPosts := []*blogpb.Post{}

	for _, post := range posts {
		tposts := transformToPost(post)
		transformedPosts = append(transformedPosts, tposts)
	}
	return transformedPosts
}
func transformToPost(p *model.Post) *blogpb.Post {
	return &blogpb.Post{
		PostID:          p.ID,
		Title:           p.Title,
		Content:         p.Content,
		Author:          p.Author,
		PublicationDate: p.PublicationDate.Format(time.RFC3339),
		Tags:            nil,
	}
}
