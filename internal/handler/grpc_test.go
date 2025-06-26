package handler

import (
	"context"
	blogpb "github.com/pradeepitm12/cb/bee/api/gen"
	"github.com/pradeepitm12/cb/bee/internal/repository/inmemory"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestBlogPostCRUDFlow(t *testing.T) {
	ctx := context.Background()
	memStore := inmemory.NewPostStore()
	svc := NewPostHandler(memStore)

	posts := []struct {
		title   string
		content string
		author  string
		tags    []string
	}{
		{"First Post", "This is the first post.", "Sohan", []string{"Apple", "Ball"}},
		{"Second Post", "This is the second posst.", "Rohan", []string{"grpc", "backend"}},
		{"Third Post", "This is the third post.", "Mohan", []string{"go", "concurrency"}},
		{"Fourth Post", "This is the fourth post", "Kumar", []string{"random"}},
	}

	createdIDs := make([]string, 0, len(posts))

	for _, p := range posts {
		req := &blogpb.CreateRequest{
			Title:   p.title,
			Content: p.content,
			Author:  p.author,
			Tags:    p.tags,
		}
		res, err := svc.CreatePost(ctx, req)
		if err != nil {
			t.Fatalf("error creating post: %v", err)
		}
		if res.GetPost().GetPostID() == "" {
			t.Fatalf("error creating post, post id is empty")
		}
		createdIDs = append(createdIDs, res.GetPost().GetPostID())
	}

	//	doing the read part
	for i, id := range createdIDs {
		res, err := svc.ReadPost(ctx, &blogpb.ReadRequest{PostID: id})
		assert.NoError(t, err)
		assert.Equal(t, posts[i].title, res.GetPost().GetTitle())
		assert.Equal(t, posts[i].author, res.GetPost().GetAuthor())
	}

	//	 test the update part for first post
	newTitle := posts[0].title + " (Updated)"
	req := &blogpb.UpdateRequest{
		Id:      createdIDs[0],
		Title:   newTitle,
		Content: posts[0].content + " [edited]",
		Author:  posts[0].author,
		Tags:    posts[0].tags,
	}
	res, err := svc.UpdatePost(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, newTitle, res.GetPost().GetTitle())

	deleteID := createdIDs[len(createdIDs)-1]
	delRes, err := svc.DeletePost(ctx, &blogpb.DeleteRequest{PostID: deleteID})
	if !strings.Contains(delRes.Message, "Deleted") {
		t.Fatalf("error deleting post")
	}

}
