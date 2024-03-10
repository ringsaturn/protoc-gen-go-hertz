// Code generated by protoc-gen-go-hertz. DO NOT EDIT.
// versions:
// - protoc-gen-go-hertz v0.3.0
// - protoc             (unknown)
// source: api/product/app/v1/v1.proto

package v1

import (
	context "context"
	app "github.com/cloudwego/hertz/pkg/app"
	server "github.com/cloudwego/hertz/pkg/app/server"
	xhertz "github.com/ringsaturn/protoc-gen-go-hertz/xhertz"
	http "net/http"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the Hertz package it is being compiled against.
var _ = new(context.Context)
var _ = new(server.Hertz)
var _ = new(app.Handler)
var _ = new(http.Handler)
var _ = new(xhertz.Error)

const OperationBlogServiceCreateArticle = "/api.product.app.v1.BlogService/CreateArticle"
const OperationBlogServiceGetArticles = "/api.product.app.v1.BlogService/GetArticles"

type BlogServiceHTTPServer interface {
	CreateArticle(context.Context, *CreateArticleRequest) (*CreateArticleResponse, error)
	GetArticles(context.Context, *GetArticlesRequest) (*GetArticlesResponse, error)
}

func RegisterBlogServiceHTTPServer(h *server.Hertz, srv BlogServiceHTTPServer) {
	group := h.Group("/")
	group.Handle("GET", "/v1/author/:author_id/articles", _BlogService_GetArticles0_HTTP_Handler(srv))
	group.Handle("POST", "/v1/author/:author_id/articles", _BlogService_CreateArticle0_HTTP_Handler(srv))
}

func _BlogService_GetArticles0_HTTP_Handler(srv BlogServiceHTTPServer) func(c context.Context, ctx *app.RequestContext) {
	return func(c context.Context, ctx *app.RequestContext) {
		var in GetArticlesRequest
		if err := ctx.BindAndValidate(&in); err != nil {
			xhertz.HandleBadRequest(ctx, err)
			return
		}

		out, err := srv.GetArticles(c, &in)
		if err != nil {
			xhertz.HandleError(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, out)
	}
}

func _BlogService_CreateArticle0_HTTP_Handler(srv BlogServiceHTTPServer) func(c context.Context, ctx *app.RequestContext) {
	return func(c context.Context, ctx *app.RequestContext) {
		var in CreateArticleRequest
		if err := ctx.BindAndValidate(&in); err != nil {
			xhertz.HandleBadRequest(ctx, err)
			return
		}

		out, err := srv.CreateArticle(c, &in)
		if err != nil {
			xhertz.HandleError(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, out)
	}
}
