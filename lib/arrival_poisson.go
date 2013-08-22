package stress

import (
    "time"
    "math/rand"
)

type PoissonArrival struct {
    lambda float64    
}

func NewPoissonArrival(lambda float64) *PoissonArrival {
    return &PoissonArrival{lambda: lambda}
}

func (a *PoissonArrival) Wait() {
    time.Sleep(time.Duration(int64(1000000000.0 * rand.ExpFloat64()/a.lambda)))
}