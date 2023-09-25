package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "[v2] Hello, Kubernetes!")
}

func main() {
	started := time.Now()
	// /healthz 接口会在启动成功的 15s 内正常返回 200 状态码，在 15s 后，会一直返回 500 的状态码。
	// 这样就能触发k8s的存活探针 (livenessProb)
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		duration := time.Since(started)
		if duration.Seconds() > 15 {
			w.WriteHeader(500)
			w.Write([]byte(fmt.Sprintf("error: %v", duration.Seconds())))
		} else {
			w.WriteHeader(200)
			w.Write([]byte("ok"))
		}
	})

	http.HandleFunc("/", hello)
	http.ListenAndServe(":3000", nil)
}
