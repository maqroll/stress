package stress

import (
    "container/ring"
    "fmt"
    vegeta "github.com/maqroll/vegeta/lib"
    "net/http"
    "text/tabwriter"
    "os"
)

func Display(results <-chan *vegeta.Result, queue chan *http.Request) {
    var lastResults *ring.Ring = nil 
    ticks := [...]byte{'-', '\\', '|', '/', '-'}
    count_ticks := 0
    var count_results uint64 = 0
    var lastN = make(map[int]*ring.Ring, 3)
    lastN[5] = nil
    lastN[50] = nil
    lastN[500] = nil
    var tpsLastN = make(map[int]float32, 3)


    f, err := os.OpenFile("out.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
    if err != nil {
        panic(err)
    }

    defer f.Close()

    w := new(tabwriter.Writer)
    w.Init(os.Stdout, 7, 0, 1, ' ', tabwriter.AlignRight)
    fmt.Fprintln(w,"\r \tserved\t-5\t-50\t-500\tqueue\t")
    w.Flush()

    for r := range results {
        count_results++
        if lastResults.Len() == 0 {
            lastResults = new(ring.Ring)
        } else if lastResults.Len() == 500 {
            lastResults = lastResults.Next()
        } else {
            lastResults.Link(new(ring.Ring))
            lastResults = lastResults.Next()
        }

        lastResults.Value = *r
        if _, err = f.WriteString(fmt.Sprintf("%d,%s,%s\n",r.Code,r.Timestamp.String(),r.Timing.String())); err != nil {
            panic(err)
        }

        for i, ring := range lastN {
            if ring != nil {
                lastN[i] = ring.Next()
            }

            if lastResults.Len() == i {
                lastN[i] = lastResults.Next()
            }

            if lastN[i] != nil {
                tpsLastN[i] = (float32(i) * 1.0e9) / float32(r.Timestamp.Add(r.Timing).UnixNano() - lastN[i].Value.(vegeta.Result).Timestamp.UnixNano())
            } else {
                tpsLastN[i] = 0.0
            }
        }

        fmt.Fprintf(w,"\r%s\t%6d\t%6.2f\t%6.2f\t%6.2f\t%5d\t                  ",string(ticks[count_ticks%len(ticks)]), count_results, tpsLastN[5], tpsLastN[50], tpsLastN[500], len(queue))
        w.Flush()
        count_ticks++
    }
}
