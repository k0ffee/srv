package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

type respHandler struct {
	mu sync.Mutex
	n  uint64
}

var (
	addr     string = os.Getenv("my_pod_ipaddress")
	vers     string = os.Getenv("app_version")
	overflow string
)

func (h *respHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.n == ^uint64(0) {
		overflow = "more than" + " "
	} else {
		h.n++
	}

	fmt.Fprintf(w,
		"address is \"%s\", %s%d requests served, version is \"%s\"\n",
		addr, overflow, h.n, vers)
}

func main() {
	var port string

	if len(os.Args) == 2 {
		port = os.Args[1]
	} else {
		port = "8080"
	}

	http.Handle("/", new(respHandler))
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
