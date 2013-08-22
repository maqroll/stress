package main

import (
	"flag"
    "log"
    "net/http"
    vegeta "github.com/maqroll/vegeta/lib"
    stress "github.com/maqroll/stress/lib"
    "runtime"
)

var arrivalProcess = flag.String("kind","D", "arrival process (M|D)")
var lambda = flag.Float64("lambda", 10, "request tps")
var requests = flag.String("rq", "targets.txt", "requests file")
var k = flag.Int("k",10000, "queue capacity")
var C = flag.Int("C", 2, "number of servers")
var requestLimit *uint64 = flag.Uint64("requests", 0, "requests (0 means forever)")

var requestCounter int = 0
var queue chan *http.Request
var results chan *vegeta.Result
var end = make(chan int)

func main() {
	flag.Parse()
	queue = make(chan *http.Request, *k)
	results = make(chan *vegeta.Result, *k)

	runtime.GOMAXPROCS(runtime.NumCPU())

    targets, err := vegeta.NewTargetsFromFile(*requests)
    if err != nil {
        log.Fatal(err)
    }

	for c := 1; c <= *C; c = c + 1 {
		go stress.Service(queue, results, end)
	}

	go stress.Display(results,queue)
	go stress.Source(*requestLimit, *lambda, queue, &requestCounter, targets, *arrivalProcess)

	for c := 1; c <= *C; c = c + 1 {
		<-end
	}
}