package main

import (
	"context"
	"net/http"
)

var userInfoKey = struct{}{}

type userInfo struct {
	Name string
}

func userFromCtx(ctx context.Context) (user *userInfo, haveUserInfo bool) {
	user, haveUserInfo = ctx.Value(userInfoKey).(*userInfo)
	return
}

func ctxWithUser(parent context.Context, user *userInfo) context.Context {
	return context.WithValue(parent, userInfoKey, user)
}

func httpHandler(rw http.ResponseWriter, req *http.Request) {
	user, haveUser := userFromCtx(req.Context())
	if !haveUser {
		rw.WriteHeader(http.StatusForbidden)
		_, _ = rw.Write([]byte(http.StatusText(http.StatusForbidden)))
		return
	}

	rw.WriteHeader(http.StatusOK)
	_, _ = rw.Write([]byte("OK: " + user.Name))
}

func authMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		userName := req.Header.Get("X-Username")
		if userName != "" {
			req = req.WithContext(ctxWithUser(req.Context(), &userInfo{
				Name: userName,
			}))
		}

		next.ServeHTTP(rw, req)
	})
}

func main() {
	http.Handle("/secure", authMiddleWare(http.HandlerFunc(httpHandler)))
	http.ListenAndServe("127.0.0.1:8005", nil)
}
