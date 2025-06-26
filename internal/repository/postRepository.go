package repository

import (
	"context"
	"time"

	"github.com/pradeepitm12/cb/bee/internal/model"
)

type PostRepository interface {
	Create(ctx *context.Context, post *model.Post) (*model.Post, error)
	Read(ctx *context.Context, id string) (*model.Post, error)
	Update(ctx *context.Context, postID, title, content, author string, tags []string, modTime time.Time) (*model.Post, error)
	Delete(ctx *context.Context, id string) string
}
