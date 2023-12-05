package {{.PkgName}}

import (
    "github.com/cat3306/demo/internal/config"
    "github.com/zeromicro/go-zero/core/logx"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	{{.ImportPackages}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		if err != nil {
			if config.AppConf.HandlerErrLog {
				logx.Errorf("{{.HandlerName}} err:%s,req:%+v", err.Error(), req)
			}
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			{{if .HasResp}}httpx.OkJsonCtx(r.Context(), w, resp){{else}}httpx.Ok(w){{end}}
		}
	}
}
