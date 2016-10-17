package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/rs/xhandler"
	"github.com/rs/xmux"

	"golang.org/x/net/context"
)

type ExampleLogger struct {
	next xhandler.HandlerC
}

func (h ExampleLogger) ServeHTTPC(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	log.Printf("%s", r.RequestURI)
	h.next.ServeHTTPC(ctx, w, r)
}

func getPort() int {
	if port, err := strconv.Atoi(os.Getenv("PORT")); err == nil {
		return port
	}
	return 8080
}

func main() {
	c := xhandler.Chain{}

	// Add close notifier handler so context is cancelled when the client closes
	// the connection
	c.UseC(xhandler.CloseHandler)

	// Add timeout handler
	c.UseC(xhandler.TimeoutHandler(2 * time.Second))

	c.UseC(func(next xhandler.HandlerC) xhandler.HandlerC {
		return ExampleLogger{next: next}
	})

	mux := xmux.New()

	// serve static assets
	serverHandler := http.FileServer(http.Dir("./"))
	mux.Handle("GET", "/*filepath", serverHandler)

	port := getPort()

	log.Printf("Running server on localhost:%d", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), c.Handler(mux)); err != nil {
		log.Fatal(err)
	}
}
