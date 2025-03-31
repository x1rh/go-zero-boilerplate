package middleware

import (
	"fmt"
	gwx "go-zero-boilerplate/pkg/zero-contrib/gatewayx"
	"go-zero-boilerplate/pkg/zero-contrib/jwtx"
	"net/http"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func Auth(jwtManager *jwtx.JWTManager, mapper *gwx.Router) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			logx.Debugf("header: %+v\n", r.Header)

			authorization := r.Header.Get("authorization")
			logx.Debug("authorization: ", authorization)

			// r.Header.Set("Grpc-Metadata-auth-date", authDateStr)
			// r.Header.Set("Grpc-Metadata-tg-first-name", url.QueryEscape(firstName))

			if !mapper.IsRequireAuth(r.Method, r.RequestURI) {
				next(w, r)
				return
			}

			var uid int64
			claims, err := jwtManager.Verify(authorization)
			if err == nil {
				uid = claims.Payload.(jwtx.User).Uid // NOTICE: check convertion
			} else {
				logx.Error(err)
				httpx.Error(w, errors.New("invalid signature"))
				return
			}
			logx.Debug(claims)
			r.Header.Set("Grpc-Metadata-uid", fmt.Sprintf("%d", uid))

			next(w, r)
		}
	}
}
