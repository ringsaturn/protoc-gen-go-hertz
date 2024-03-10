{{$svrType := .ServiceType}}
{{$svrName := .ServiceName}}

{{- range .MethodSets}}
const Operation{{$svrType}}{{.OriginalName}} = "/{{$svrName}}/{{.OriginalName}}"
{{- end}}

type {{.ServiceType}}HTTPServer interface {
{{- range .MethodSets}}
	{{- if ne .Comment ""}}
	{{.Comment}}
	{{- end}}
	{{.Name}}(context.Context, *{{.Request}}) (*{{.Reply}}, error)
{{- end}}
}

func Register{{.ServiceType}}HTTPServer(h *server.Hertz, srv {{.ServiceType}}HTTPServer) {
	group := h.Group("/")
	{{- range .Methods}}
	group.Handle("{{.Method}}", "{{.Path}}", _{{$svrType}}_{{.Name}}{{.Num}}_HTTP_Handler(srv))
	{{- end}}
}

{{range .Methods}}
func _{{$svrType}}_{{.Name}}{{.Num}}_HTTP_Handler(srv {{$svrType}}HTTPServer) func(c context.Context, ctx *app.RequestContext) {
	return func(c context.Context, ctx *app.RequestContext) {
		var in {{.Request}}
		if err := ctx.BindAndValidate(&in); err != nil {
			xhertz.HandleBadRequest(ctx, err)
			return
		}

		out, err := srv.{{.Name}}(c, &in)
		if err != nil {
			xhertz.HandleError(ctx, err)
			return
		}
		xhertz.Render(ctx, http.StatusOK, out)
	}
}
{{end}}
