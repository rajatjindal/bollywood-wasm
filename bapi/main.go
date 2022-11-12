package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"

	spinhttp "github.com/fermyon/spin/sdk/go/http"
	"github.com/rajatjindal/bollywood-wasm/bapi/pkg/movie"
)

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("inside handle")
		if r.Method == http.MethodOptions {
			w.Header().Set("access-control-allow-origin", "*")
			w.Header().Set("access-control-allow-methods", "OPTIONS, GET, POST")
			w.Header().Set("access-control-allow-headers", "authorization, content-type")
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method == http.MethodPost && r.URL.Path == "/bapi/new" {
			fmt.Println("inside this method")
			movie.NewGame(w, r)
			return
		}

		if r.Method == http.MethodPost && strings.Contains(r.URL.Path, "/bapi/guess") {
			movie.MakeAGuess(w, r)
			return
		}

		if r.Method == http.MethodGet && strings.Contains(r.URL.Path, "/bapi/game") {
			movie.GetGameState(w, r)
			return
		}

		d, _ := httputil.DumpRequest(r, true)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(d))
	})
}

func main() {}
