package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app/server"
	v1 "github.com/ringsaturn/protoc-gen-go-hertz/_example/hertz-blog-service/api/product/app/v1"
)

type Service struct{}

func (s Service) CreateArticle(ctx context.Context, article *v1.CreateArticleRequest) (*v1.CreateArticleResponse, error) {
	return &v1.CreateArticleResponse{
		Article: &v1.Article{
			Title:    article.Title,
			Content:  article.Content,
			AuthorId: article.AuthorId,
		},
	}, nil
}

func (s Service) GetArticles(ctx context.Context, req *v1.GetArticlesRequest) (*v1.GetArticlesResponse, error) {
	return &v1.GetArticlesResponse{
		Total: 1,
		Articles: []*v1.Article{
			{
				Title:    "test article: " + "Hello",
				Content:  "test",
				AuthorId: 1,
			},
		},
	}, nil
}

func main() {
	h := server.Default()
	v1.RegisterBlogServiceHTTPServer(h, &Service{})
	h.Run()
}
