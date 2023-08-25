package result

import (
	"app/common/xerr"
	"fmt"
	"google.golang.org/grpc/status"
	"net/http"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func HTTPResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {

	if err == nil {
		// 成功返回
		httpx.WriteJson(w, http.StatusOK, Success(resp))
	} else {
		// 错误返回
		errCode := xerr.ServerCommonError
		errMsg := "服务器开小差啦，稍后再来试一试"

		causeErr := errors.Cause(err) // err类型
		var e *xerr.CodeError
		if errors.As(causeErr, &e) { // 自定义错误类型
			// 自定义CodeError
			errCode = e.Code
			errMsg = e.Msg
		} else {
			// rpc 错误
			s, ok := status.FromError(err)
			if ok {
				errCode = uint32(s.Code())
				errMsg = s.Message()
				logx.WithContext(r.Context()).Errorf("【RPC-ERR】 : %+v ", err)
				httpx.WriteJson(w, 494, Error(errCode, errMsg))
			} else {
				logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)
				httpx.WriteJson(w, http.StatusBadRequest, Error(errCode, errMsg))
			}

		}

		logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)

		httpx.WriteJson(w, http.StatusBadRequest, Error(errCode, errMsg))
	}
}

func ParamErrorResult(r *http.Request, w http.ResponseWriter, err error) {
	errMsg := fmt.Sprintf("%s ,%s", xerr.ErrMsg(xerr.RequestParamError), err.Error())
	httpx.WriteJson(w, http.StatusBadRequest, Error(xerr.RequestParamError, errMsg))
}
