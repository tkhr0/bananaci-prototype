package server

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"

	jobQueue "github.com/tkhr0/bananaci-prototype/job_queue"
	"github.com/tkhr0/bananaci-prototype/phase"
	"github.com/tkhr0/bananaci-prototype/runtime"
)

var dispatcher *jobQueue.Dispatcher

func CallHTTP(d *jobQueue.Dispatcher) {
	dispatcher = d

	httpHandler()

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		os.Exit(1)
	}
}

func httpHandler() {
	http.HandleFunc("/endpoint", pathEndpoint)
	http.HandleFunc("/new", pathNew)

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

func pathNew(w http.ResponseWriter, r *http.Request) {
	for k, params := range parseQuery(r) {
		for _, param := range params {
			switch k {
			case "opened":
				n, err := strconv.Atoi(param)
				if err != nil {
					continue
				}
				for i := 0; i < n; i++ {
					dispatcher.Add(*jobQueue.NewJob(*runtime.NewRuntime()))
				}
			case "labeled":
				n, err := strconv.Atoi(param)
				if err != nil {
					continue
				}
				for i := 0; i < n; i++ {
					r := *runtime.NewRuntime()
					r.Phase = phase.NewOpened().ToLabeled()
					dispatcher.Add(*jobQueue.NewJob(*runtime.NewRuntime()))
				}
			}
		}
	}
}
