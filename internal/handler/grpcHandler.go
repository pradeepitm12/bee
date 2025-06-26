package handler

import (
	"context"
	"fmt"
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
	newPost := model.NewPost(req.Title, req.Content, req.Author, req.Tags)
	post, err := ph.repo.Create(&ctx, newPost)
	if err != nil && strings.Contains(err.Error(), errors.PostAlreadyExists) {
		return nil, status.Errorf(codes.AlreadyExists, "Post already exists")
	}
	resp := &blogpb.CreateResponse{
		Post:  transformToPost(post),
		Error: "",
	}
	return resp, nil
}

func (ph *PostHandler) ReadPost(ctx context.Context, req *blogpb.ReadRequest) (*blogpb.ReadResponse, error) {
	post, err := ph.repo.Read(&ctx, req.PostID)
	if err != nil {
		return nil, status.Error(codes.NotFound, "post not found")
	}
	resp := &blogpb.ReadResponse{
		Post:  transformToPost(post),
		Error: "",
	}
	return resp, nil
}

func (ph *PostHandler) UpdatePost(ctx context.Context, req *blogpb.UpdateRequest) (*blogpb.UpdateResponse, error) {
	post, err := ph.repo.Update(&ctx, req.Id, req.Title, req.Content, req.Author, req.Tags, time.Now())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "post not found")
	}
	resp := &blogpb.UpdateResponse{Post: transformToPost(post), Error: ""}
	return resp, nil
}

func (ph *PostHandler) DeletePost(ctx context.Context, req *blogpb.DeleteRequest) (*blogpb.DeleteResponse, error) {
	msg := ph.repo.Delete(&ctx, req.PostID)
	if msg != "" && msg == errors.PostNotFound {
		return nil, status.Error(codes.NotFound, fmt.Sprintf(errors.PostNotFound+": %s", req.PostID))
	}
	resp := &blogpb.DeleteResponse{Message: fmt.Sprintf("Post %s Deleted", req.PostID)}
	return resp, nil
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
