package services

import (
	"hmdp/internal/domain/repository"
	"hmdp/pkg/serializer"
)

type IBlogService interface {
	GetBlog(id string) (serializer.Response, error)
	Hot(page, pageSize int) (serializer.Response, error)
}

type BlogConfiguration func(os *BlogService) error

func NewBlogService(cfgs ...BlogConfiguration) IBlogService {
	// Create the user service
	os := &BlogService{}
	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		// Pass the service into the configuration function
		err := cfg(os)
		if err != nil {
			return nil
		}
	}
	return os
}

func WithBlogRepo(userRepo repository.IBlogRepo) BlogConfiguration {
	return func(os *BlogService) error {
		os.blogRepo = userRepo
		return nil
	}
}

type BlogService struct {
	blogRepo repository.IBlogRepo
}

func (b BlogService) GetBlog(id string) (serializer.Response, error) {
	blog, err := b.blogRepo.GetBlog(id)
	if err != nil {
		return serializer.Response{}, err
	}
	return serializer.Response{Data: blog}, err
}

func (b BlogService) Hot(page, pageSize int) (serializer.Response, error) {
	blogs, err := b.blogRepo.GetBlogs(page, pageSize)
	if err != nil {
		return serializer.Response{}, err
	}
	return serializer.Response{Data: blogs}, nil
}
