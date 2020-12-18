package controller

import (
	"chitChat/connectPool"
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	w.WriteHeader(200)
	fmt.Println(r.RemoteAddr)
	_, _ = fmt.Fprint(w, []string{r.RemoteAddr, r.Host, r.Method})
}

func Send(w http.ResponseWriter, r *http.Request) {
	connectPool.Init()
	w.Header().Set("content-type", "text/html")
	w.WriteHeader(200)
	_, _ = fmt.Fprint(w, []string{r.RemoteAddr, r.Host, r.Method})
}
