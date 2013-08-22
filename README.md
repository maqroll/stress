stress
======

HTTP load testing tool focused on simplicity.

Test your service capacity (Can handle 1k TPS?) and dumps the full data to further analysis (service time percentiles, etc..).

![Architecture](https://googledrive.com/host/0ByaAHLZo5u8GOTZCT2RsMGJoYlk/stress.png)

I know, it looks like a server architecture but in fact is the tool architecture.

## usage
```shell
$ stress -h
Usage of stress:
  -C=2: number of servers
  -k=10000: queue capacity
  -kind="D": arrival process (M|D)
  -lambda=10: request tps
  -requests=0: requests (0 means forever)
  -rq="targets.txt": requests file
```

stress populates the queue with requests from the file *rq* at rate *lambda* with arrival process *kind*. The servers (as many as set by *C*) pick up requests from the queue and process them.

Reports the following parameters:
```
    - served: request processed
    - the rate (transactions per second) for the last 5,50 and 500 requests processed.
    - queue: request waiting (in the queue, not the ones waiting in the server)
```

Dumps (out.txt) all the requests processed in the following format:

```
result_code,timestamp,duration
```  

Timestamp format as provided by time.Time.String().

Duration format as provided by time.Duration.String().