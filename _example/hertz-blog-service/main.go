package main

import (
	"context"
	_ "embed"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/swagger"
	v1 "github.com/ringsaturn/protoc-gen-go-hertz/_example/hertz-blog-service/api/product/app/v1"
	swaggerFiles "github.com/swaggo/files"
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

//go:embed gen/openapi.yaml
var openapiYAML []byte

func bindSwagger(h *server.Hertz) {
	h.GET("/swagger/*any", swagger.WrapHandler(
		swaggerFiles.Handler,
		swagger.URL("/openapi.yaml"),
	))

	h.GET("/openapi.yaml", func(c context.Context, ctx *app.RequestContext) {
		ctx.Header("Content-Type", "application/x-yaml")
		ctx.Write(openapiYAML)
	})
}

func main() {
	h := server.Default()

	bindSwagger(h)

	v1.RegisterBlogServiceHTTPServer(h, &Service{})
	h.Run()
}
