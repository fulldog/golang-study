package main

import (
	"chitChat/controller"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", controller.Index)
	mux.HandleFunc("/send", controller.Send)
	serve := &http.Server{
		Addr:              ":8888",
		Handler:           mux,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	err := serve.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	w.WriteHeader(200)
	fmt.Println(r.RemoteAddr)
	_, _ = fmt.Fprint(w, []string{r.RemoteAddr, r.Host, r.Method})
}
