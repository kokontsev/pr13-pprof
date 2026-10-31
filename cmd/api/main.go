package main

import (
	"example.com/pr13/internal/work"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"strconv"
	"time"
)

func main() {
	runtime.SetBlockProfileRate(1)
	runtime.SetMutexProfileFraction(1)
	m := http.NewServeMux()
	m.Handle("/debug/pprof/", http.DefaultServeMux)
	m.HandleFunc("GET /work", func(w http.ResponseWriter, r *http.Request) {
		n := 38
		if v := r.URL.Query().Get("n"); v != "" {
			var e error
			n, e = strconv.Atoi(v)
			if e != nil || n < 0 || n > 40 {
				http.Error(w, "n must be 0..40", 400)
				return
			}
		}
		defer work.TimeIt("Fib")()
		v := 0
		if r.URL.Query().Get("fast") == "1" {
			v = work.FibFast(n)
		} else {
			v = work.Fib(n)
		}
		fmt.Fprintln(w, v)
	})
	a := os.Getenv("HTTP_ADDR")
	if a == "" {
		a = "127.0.0.1:8080"
	}
	log.Fatal((&http.Server{Addr: a, Handler: m, ReadHeaderTimeout: 5 * time.Second}).ListenAndServe())
}
