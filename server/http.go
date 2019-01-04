package server

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func CallHTTP() {

	httpHandler()

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		os.Exit(1)
	}
}

func httpHandler() {
	http.HandleFunc("/endpoint", pathEndpoint)

}

func parseQuery(r *http.Request) url.Values {
	return r.URL.Query()
}

func pathEndpoint(w http.ResponseWriter, r *http.Request) {
	params := parseQuery(r)

	fmt.Println(len(params))
	for k, param := range params {
		fmt.Printf("%s: %s\n", k, param)
	}
}
