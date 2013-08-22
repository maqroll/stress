package stress

import (
    "time"
)

type DeterministicArrival struct {
    ticker *time.Ticker
}

func NewDeterministicArrival(lambda float64) *DeterministicArrival {
    return &DeterministicArrival{ticker: time.NewTicker(time.Duration(int64(1000000000.0 / lambda)))}
}

func (a *DeterministicArrival) Wait() {
    <- a.ticker.C
}