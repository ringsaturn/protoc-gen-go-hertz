type {{ $.InterfaceName }} interface {
{{range .MethodSet}}
	{{.Name}}(context.Context, *{{.Request}}) (*{{.Reply}}, error)
{{end}}
}
func Register{{ $.InterfaceName }}(h *server.Hertz, srv {{ $.InterfaceName }}) {
	s := {{.Name}}{
		server: srv,
		h:     h,
	}
	s.RegisterService()
}

type {{$.Name}} struct{
	server {{ $.InterfaceName }}
	h *server.Hertz
}

{{range .Methods}}
func (s *{{$.Name}}) {{ .HandlerName }} (c context.Context, ctx *app.RequestContext) {
	var in {{.Request}}

	if err := ctx.BindAndValidate(&in); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.H{"error": err.Error()})
		return
	}
	out, err := s.server.({{ $.InterfaceName }}).{{.Name}}(c, &in)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, out)
}
{{end}}

func (s *{{$.Name}}) RegisterService() {
{{range .Methods}}
		s.h.Handle("{{.Method}}", "{{.Path}}", s.{{ .HandlerName }})
{{end}}
}