package main

import (
	"net/http"
	"time"
)

func httpSlowHandler(rw http.ResponseWriter, req *http.Request) {
	reqContext := req.Context()
	select {
	case <-time.After(time.Second * 10):

	case <-reqContext.Done():
		println("request canceled!")
		return
	}
	_, _ = rw.Write([]byte("returned slowly..."))
}

func main() {
	http.Handle("/slow", http.HandlerFunc(httpSlowHandler))
	http.ListenAndServe("127.0.0.1:8006", nil)
}
