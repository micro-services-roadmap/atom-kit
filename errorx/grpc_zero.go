package errorx

import (
	"github.com/micro-services-roadmap/oneid-core/model"
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
	v, ok := err.(*model.CodeError)
	if !ok {
		httpx.WriteJsonCtx(r.Context(), w, http.StatusBadRequest, model.Res(c, nil, err.Error()))
		return
	}

	if v.Err == nil {
		httpx.WriteJsonCtx(r.Context(), w, http.StatusBadRequest, model.Res(v.Code, nil, v.Error()))
		return
	}

	err = GrpcError(v.Err)
	if v, ok := err.(*model.CodeError); ok {
		httpx.WriteJsonCtx(r.Context(), w, http.StatusBadRequest, model.Res(v.Code, nil, v.Error()))
	} else {
		httpx.WriteJsonCtx(r.Context(), w, http.StatusBadRequest, model.Res(c, nil, err.Error()))
	}
}
