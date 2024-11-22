package errorx

import (
	"github.com/micro-services-roadmap/oneid-core/modelo"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func ErrorRespond(w http.ResponseWriter, r *http.Request, resp any, err error, code ...int) {
	if err == nil {
		httpx.OkJsonCtx(r.Context(), w, resp)
		return
	}

	var c int
	if len(code) > 0 {
		c = code[0]
	} else {
		c = 888_888
	}

	err = GrpcError(err)
	v, ok := err.(*modelo.CodeError)
	if !ok {
		httpx.WriteJsonCtx(r.Context(), w, http.StatusBadRequest, modelo.Res(c, nil, err.Error()))
		return
	}

	if v.Err == nil {
		httpx.WriteJsonCtx(r.Context(), w, http.StatusBadRequest, modelo.Res(v.Code, nil, v.Error()))
		return
	}

	err = GrpcError(v.Err)
	if v, ok := err.(*modelo.CodeError); ok {
		httpx.WriteJsonCtx(r.Context(), w, http.StatusBadRequest, modelo.Res(v.Code, nil, v.Error()))
	} else {
		httpx.WriteJsonCtx(r.Context(), w, http.StatusBadRequest, modelo.Res(c, nil, err.Error()))
	}
}
