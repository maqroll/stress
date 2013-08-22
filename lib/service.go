package stress

import (
    "net/http"
    vegeta "github.com/maqroll/vegeta/lib"
)

func Service(q <-chan *http.Request, results chan *vegeta.Result, end chan<-int) {
    for request:= range q {
        vegeta.Hit(request,results)
    }
    end<-1
}