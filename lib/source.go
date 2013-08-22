package stress

import (
    "net/http"
    vegeta "github.com/maqroll/vegeta/lib"
)

func Source(requestLimit uint64, lambda float64, q chan *http.Request, requestCounter *int, targets vegeta.Targets, arrivalProcess string) {
    var arrivalStrategy Arrival
    
    switch arrivalProcess {
        case "M":
            arrivalStrategy = NewPoissonArrival(lambda)            
        case "D":    
            arrivalStrategy = NewDeterministicArrival(lambda)
    }


    for ; requestLimit == 0 || *requestCounter < int(requestLimit) ; *requestCounter++ {
        arrivalStrategy.Wait()
        q<-targets[*requestCounter%len(targets)] /* next arrival */
    }

    close(q)
}