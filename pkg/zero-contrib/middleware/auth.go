package middleware

import (
	"fmt"
	"go-zero-boilerplate/internal/config"
	"go-zero-boilerplate/pkg/telegramx"
	gwx "go-zero-boilerplate/pkg/zero-contrib/gatewayx"
	"go-zero-boilerplate/pkg/zero-contrib/jwtx"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func Auth(jwtManager *jwtx.JWTManager, botConfig []config.TelegramBotConfig, mapper *gwx.Router) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			logx.Debugf("header: %+v\n", r.Header)

			telegram := r.Header.Get("x-telegram")
			authorization := r.Header.Get("authorization")

			logx.Debug("telegram: ", telegram)
			logx.Debug("authorization: ", authorization)

			var jwtUid int64
			var tgUid int64
			var authDateStr, firstName, lastName, username string
			var secret string
			var expire int64

			botName := "zero" // NOTICE:
			for _, c := range botConfig {
				if c.Name == botName {
					secret = c.Secret
					expire = c.Expire
				}
			}

			if secret == "" {
				httpx.Error(w, errors.New("invalid bot config"))
				return
			}

			u, err := telegramx.GetUser(telegram, secret)
			if err != nil {
				logx.Error(err)
			} else {
				tgUid = u.Id
			}

			logx.Debug(u)

			verify := false
			if u != nil {
				authDate, err := strconv.Atoi(u.AuthDate)
				if err != nil {
					logx.Error(err)
					verify = false
				}
				now := time.Now().Unix()
				delta := now - int64(authDate)
				if delta > expire || delta < 0 {
					logx.Error("invalid timestamp")
					verify = false
				} else {
					verify = true
				}
			}

			if verify {
				authDateStr = u.AuthDate
				firstName = u.FirstName
				lastName = u.LastName
				username = u.Username
			} else {
				jwtUid = 0
				tgUid = 0
			}

			r.Header.Set("Grpc-Metadata-auth-date", authDateStr)
			r.Header.Set("Grpc-Metadata-tg-first-name", url.QueryEscape(firstName))
			r.Header.Set("Grpc-Metadata-tg-last-name", url.QueryEscape(lastName))
			r.Header.Set("Grpc-Metadata-tg-username", url.QueryEscape(username))
			r.Header.Set("Grpc-Metadata-tg-uid", fmt.Sprintf("%d", tgUid))

			if !mapper.IsRequireAuth(r.Method, r.RequestURI) {
				next(w, r)
				return
			}

			claims, err := jwtManager.Verify(authorization)
			if err == nil {
				jwtUid = claims.Payload.(jwtx.User).Uid // NOTICE: check convertion
			} else {
				logx.Error(err)
				httpx.Error(w, errors.New("invalid signature"))
				return
			}
			logx.Debug(claims)

			r.Header.Set("Grpc-Metadata-jwt-uid", fmt.Sprintf("%d", jwtUid))
			r.Header.Set("Grpc-Metadata-uid", fmt.Sprintf("%d", jwtUid))

			next(w, r)
		}
	}
}
