package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("<h1>welcome to cloudnative</h1>"))\
	// 03.设置version
	os.Setenv("VERSION", "0.1")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	fmt.Printf("OS VERSION: %v\n", version)

	// 02.将requst中的header 设置到 reponse中
	for k, v := range r.Header {
		//fmt.Println(k, v)
		for _, vv := range v {
			fmt.Printf("Header key: %s, Header value: %s\n", k, v)
			//Set 方法 会覆盖之前已经存在的值
			w.Header().Set(k, vv)
			//w.Header().Add(k, vv)
		}
	}
	// 04.记录日志并输出
	currentip := getCurrentIP(r)
	println(r.RemoteAddr)
	//fmt.Printf("clientip: %v\n", clientip)
	log.Printf("Success! clientip: %s", currentip)

	clientip := ClientIP(r)
	println(r.RemoteAddr)
	//fmt.Printf("clientip: %v\n", clientip)
	log.Printf("Success! clientip: %s", clientip)

}

// func healthz(w http.ResponseWriter, r *http.Request) {

// 	fmt.Fprintf(w, "working")
// }

func getCurrentIP(r *http.Request) string {

	// 这里也可以通过X-Forwarded-For请求头的第一个值作为用户的ip
	// 但是要注意的是这两个请求头代表的ip都有可能是伪造的
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		// 当请求头不存在即不存在代理时直接获取ip
		ip = strings.Split(r.RemoteAddr, ":")[0]

	}
	return ip
}

// ClientIP 尽最大努力实现获取客户端 IP 的算法。
// 解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。

func ClientIP(r *http.Request) string {
	//解析 X-Forwarded-For
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	// 解析 X-Real-IP
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	// 判断解析RemoteAddr 并判断是否为空
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

// 05.健康检查的路由
func healthz(w http.ResponseWriter, r *http.Request) {
	//Fprintf：来格式化并输出到 io.Writers 而不是 os.Stdout。

	fmt.Fprintf(w, "OK,working")
}

func main() {
	// 构架http服务器
	mux := http.NewServeMux()
	// 06. debug
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	// 指定根路由
	mux.HandleFunc("/", index)
	// healthz 路由
	mux.HandleFunc("/healthz", healthz)
	// 判断httpserver 是否启动成功
	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatalf("Start httpserver failed, error: %v", err.Error())

	}
}
