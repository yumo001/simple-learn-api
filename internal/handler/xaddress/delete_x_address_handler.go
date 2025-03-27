package xaddress

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/yumo001/simple-learn-api/internal/logic/xaddress"
	"github.com/yumo001/simple-learn-api/internal/svc"
	"github.com/yumo001/simple-learn-api/internal/types"
)

// swagger:route post /x_address/delete xaddress DeleteXAddress
//
// Delete x address information | 删除XAddress信息
//
// Delete x address information | 删除XAddress信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDsReq
//
// Responses:
//  200: BaseMsgResp

func DeleteXAddressHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDsReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := xaddress.NewDeleteXAddressLogic(r.Context(), svcCtx)
		resp, err := l.DeleteXAddress(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
