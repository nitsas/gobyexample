// Start a server on the given port (default 8090).
//
// Example usage:
// go run http-server.go
// go run http-server.go --port=8091
//
// Then in another terminal use curl to send requests to the available endpoints:
// curl localhost:8090/hello
// curl localhost:8090/headers

package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func randIndex(size int) int {
	randSource := rand.NewSource(time.Now().UnixNano())
	r := rand.New(randSource)
	return r.Intn(size)
}

func hello(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s", req.Method, req.RequestURI)
	msgs := []string{"hello\n", "hi\n"}
	fmt.Fprintf(w, msgs[randIndex(len(msgs))])
}

func headers(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s", req.Method, req.RequestURI)
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func getPort() uint {
	portPtr := flag.Uint("port", 8090, "the local port to start listening at.")
	flag.Parse()

	return *portPtr
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	port := getPort()
	addr := fmt.Sprintf("localhost:%d", port)

	log.Printf("Starting server at %s. Stop it with Ctrl-C.\n", addr)
	log.Printf("Available endpoints: GET /hello | GET /headers\n")

	http.ListenAndServe(addr, nil)
	// 2nd arg nil means "use the default router that we set up using calls to http.HandleFunc()"
}
