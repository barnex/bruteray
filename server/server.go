package server

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/barnex/bruteray"
)

var (
	port = flag.String("http", ":3700", "Port to serve HTTP")
)

var (
	env *bruteray.Env
)

const (
	DefaultWidth  = 800
	DefaultHeight = 600
)

func Serve(e *bruteray.Env) {

	env = e

	http.HandleFunc("/render", render)
	log.Println("listen", *port)
	log.Fatal(http.ListenAndServe(*port, nil))
}

func render(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	W := parseInt(q.Get("w"), DefaultWidth)
	H := parseInt(q.Get("h"), DefaultHeight)

}

func parseInt(s string, Default int) int {
	x, err := strconv.Atoi(s)
	Print(err)
	if x == 0 {
		return Default
	}
	return x
}

func Print(err error) {
	if err != nil {
		log.Println(err)
	}
}
