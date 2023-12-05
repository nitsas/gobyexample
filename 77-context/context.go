// Start a server on the given port (default 8090).
//
// Example usage:
// go run http-server.go
// go run http-server.go --port=8091
//
// Then in another terminal use curl to send requests to the available endpoints:
// curl localhost:8090/hello

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	endpoint := fmt.Sprintf("%s %s", req.Method, req.RequestURI)
	log.Println(endpoint)
	defer log.Printf("%s end\n", endpoint)

	ctx := req.Context()
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "delayed hello\n")
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("%s error: %#v\n", endpoint, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		// When the requester kills the request (e.g. via Ctrl-C in curl), then
		// err.Error() is "context canceled". http.StatusInternalServerError is always 500.
	}
}

func getPort() uint {
	portPtr := flag.Uint("port", 8090, "the local port to start listening at.")
	flag.Parse()

	return *portPtr
}

func main() {
	http.HandleFunc("/hello", hello)

	port := getPort()
	addr := fmt.Sprintf("localhost:%d", port)

	log.Printf("Starting server at %s. Stop it with Ctrl-C.\n", addr)
	log.Printf("Available endpoints: GET /hello\n")

	http.ListenAndServe(addr, nil)
	// 2nd arg nil means "use the default router that we set up using calls to http.HandleFunc()"
}
