package inmemory

import (
	"context"
	"fmt"
	"slices"
	"sync"
	"time"

	"github.com/pradeepitm12/cb/bee/internal/errors"
	"github.com/pradeepitm12/cb/bee/internal/model"
)

type PostStore struct {
	rwMutex sync.RWMutex
	store   map[string]*model.Post
}

func NewPostStore() *PostStore {
	return &PostStore{
		store: make(map[string]*model.Post),
	}
}

func (s *PostStore) Create(ctx *context.Context, post *model.Post) (*model.Post, error) {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()
	if _, ok := s.store[post.ID]; ok {
		return nil, fmt.Errorf(errors.PostAlreadyExists)
	}
	s.store[post.ID] = post
	return post, nil
}

func (s *PostStore) Read(ctx *context.Context, id string) (*model.Post, error) {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	post, ok := s.store[id]
	if !ok {
		return nil, fmt.Errorf(errors.PostNotFound)
	}
	return post, nil
}

func (s *PostStore) Update(ctx *context.Context, postID, title, content, author string, tags []string, modTime time.Time) (*model.Post, error) {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()
	post, ok := s.store[postID]
	if !ok {
		return nil, fmt.Errorf(errors.PostNotFound)
	}
	if title != "" {
		post.Title = title
	}
	if content != "" {
		post.Content = content
	}
	if author != "" {
		post.Author = author
	}
	if tags != nil {
		post.Tags = append(post.Tags, tags...)
	}
	post.LastModified = modTime
	s.store[post.ID] = post
	return post, nil
}

func (s *PostStore) Delete(ctx *context.Context, id string) string {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()
	//wg := sync.WaitGroup{}
	//for i := 0; i < 5; i++ {
	//	wg.Add(1)
	//	go func(wg *sync.WaitGroup) {
	//		defer wg.Done()
	//		// doing something
	//	}(wg)
	//}
	//
	//wg.Wait()
	//
	//val := read(post)

	_, ok := s.store[id]
	if !ok {
		return errors.PostNotFound
	}
	delete(s.store, id)
	return fmt.Sprintf("post deleted: %s", id)
}

func (s *PostStore) List(ctx *context.Context) []*model.Post {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	posts := make([]*model.Post, 0, len(s.store))
	for _, post := range s.store {
		posts = append(posts, post)
	}

	slices.SortFunc(posts, func(posta, postb *model.Post) int {
		if posta.Title < postb.Title {
			return -1
		} else {
			return 1
		}
		return 0
	})
	return posts
}
