package xhertz

import (
	"github.com/cloudwego/hertz/pkg/app"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func buildRequestMetadata(ctx *app.RequestContext) map[string]string {
	return map[string]string{
		"request_ip":         ctx.ClientIP(),
		"request_user_agent": string(ctx.UserAgent()),
		"request_uri":        ctx.URI().String(),
		"request_method":     string(ctx.Method()),
		"request_host":       string(ctx.Host()),
		"request_path":       string(ctx.Path()),
		"request_query":      ctx.QueryArgs().String(),
	}
}

type Error struct {
	HTTPCode   int        `json:"code"`    // The HTTP status code that corresponds to `google.rpc.Status.code`.
	Message    string     `json:"message"` // This corresponds to `google.rpc.Status.message`.
	GRPCStatus codes.Code `json:"status"`  // This is the enum version for `google.rpc.Status.code`.
	Reason     string     `json:"reason"`
	Details    []any      `json:"details"` // This corresponds to `google.rpc.Status.details`.
}

type ErrorOption func(*Error)

func WithMessage(message string) ErrorOption {
	return func(err *Error) {
		err.Message = message
	}
}

func WithDetails(details ...any) ErrorOption {
	return func(err *Error) {
		err.Details = append(err.Details, details...)
	}
}

func NewError(code codes.Code, message string, opts ...ErrorOption) *Error {
	httpCode := FromGRPCCode(code)
	err := &Error{
		HTTPCode:   httpCode,
		Message:    message,
		GRPCStatus: code,
		Reason:     code.String(),
		Details:    []any{},
	}
	for _, opt := range opts {
		opt(err)
	}
	return err
}

func (e *Error) Error() string {
	return e.Message
}

// This message defines the error schema for Google's JSON HTTP APIs.
type ErrorResponse struct {
	// The actual error payload. The nested message structure is for backward
	// compatibility with [Google API Client
	// Libraries](https://developers.google.com/api-client-library). It also
	// makes the error more readable to developers.
	Status Error `json:"error"`
}

func HandleError(ctx *app.RequestContext, err error) {
	if xErr, ok := err.(*Error); ok {
		renderErr(ctx, xErr)
		return
	}
	status := NewError(codes.Unknown, err.Error(), WithDetails(buildRequestMetadata(ctx)))
	renderErr(ctx, status)
}

func HandleBadRequest(ctx *app.RequestContext, err error) {
	status := NewError(codes.InvalidArgument, err.Error(), WithDetails(buildRequestMetadata(ctx)))
	renderErr(ctx, status)
}

func renderErr(ctx *app.RequestContext, err *Error) {
	ctx.JSON(err.HTTPCode, &ErrorResponse{Status: *err})
}

func Render(ctx *app.RequestContext, status int, data any) {
	m, ok := data.(proto.Message)
	if !ok {
		ctx.JSON(status, data)
		return
	}
	b, err := protojson.Marshal(m)
	if err != nil {
		ctx.JSON(status, data)
		return
	}
	ctx.SetContentType("application/json")
	_, _ = ctx.Write(b)
}
