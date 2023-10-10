package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/chfern/golang-simple-vanity-server/internal"
)

var (
	aliases = map[string]internal.ActualLocation{
		"hehe/haha": {Protocol: "git", Uri: "<actualurl>"}, // TODO CHANGEME
		// Example:
		// "custom/modpath": {Protocol: "git", Uri: "ssh://git@github.com/username/yourmodule"},

		// If you want to add more aliases, do so here
	}
)

const (
	port             = 8000
	vanityServerHost = "<vanityhost>" // TODO CHANGEME. Example: "vanityserver.com"
)

// handler will read the request path, then try to see if it matches any aliases declared above
//
// Example: given following request localhost:8000/custom/path/to/mod
// It will search aliases using the following keys: custom/path/to/mod, custom/path/to, custom/path, custom
func handler(w http.ResponseWriter, r *http.Request) {
	pathSplitted := strings.Split(strings.TrimPrefix(r.URL.Path, "/"), "/")

	for i := 0; i < len(pathSplitted); i++ {
		aliasPath := strings.Join(pathSplitted[:len(pathSplitted)-i], "/")
		fmt.Printf("Searching alias with path: %s\n", aliasPath)

		if actualLocation, ok := aliases[aliasPath]; ok {
			resp := internal.CreateVanityServerResponse(vanityServerHost, aliasPath, actualLocation)
			w.Write([]byte(resp))
			return
		}
	}
	http.Error(w, "Alias not found", 404)
}

func main() {
	log.Println("Starting")

	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
