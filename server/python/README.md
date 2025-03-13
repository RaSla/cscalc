# Python

## Build

Change dir to the `server/python`

### by Docker

* `make builder-shell` - run shell inside BUILDER docker-container
* `make build-docker` - make App docker-image

## RUN Application

* `make run` - run App in Host-OS
* `make run-by-docker` - run App in RUNNER docker-container
* `make run-docker` - run App docker-image (`make build-docker` must be run first)
* `make list-images` - list docker-images of this App

## CLEAN

* `make clean` - clean docker-stuff & build-artifacts

## Benchmarks

### Run App in Host-OS

```shell
## python 3.12.3
# HW: Core i5-1135G7 @ 2.40GHz / 32 Gb dual-DDR4
# OS: Linux Mint 22.1 / 6.11.0-19-generic x86_64

$ siege -i -c10 -t30s -f urls-server.txt
{	"transactions":			       71264,
	"availability":			      100.00,
	"elapsed_time":			       29.35,
	"data_transferred":		        1.40,
	"response_time":		        0.00,
	"transaction_rate":		     2428.07,
	"throughput":			        0.05,
	"concurrency":			        9.95,
	"successful_transactions":     71264,
	"failed_transactions":	           0,
	"longest_transaction":	        0.02,
	"shortest_transaction":	        0.00
}

$ wrk -t2 -c50 -d30 'http://127.0.0.1:8080/plus?a=10&b=5'
Running 30s test @ http://127.0.0.1:8080/plus?a=10&b=5
  2 threads and 50 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    18.87ms  795.05us  29.86ms   75.56%
    Req/Sec     1.33k    29.18     1.40k    73.17%
  79317 requests in 30.01s, 15.05MB read
Requests/sec:   2643.34
Transfer/sec:    513.71KB
```

### Run App in Docker-image

```shell
## python 3.12.3
# HW: Core i5-1135G7 @ 2.40GHz / 32 Gb dual-DDR4
# OS: Linux Mint 22.1 / 6.11.0-19-generic x86_64

$ siege -i -c10 -t30s 'http://127.0.0.1:8080/plus?a=10&b=5'
{	"transactions":			       28231,
	"availability":			       96.47,
	"elapsed_time":			       25.34,
	"data_transferred":		        0.67,
	"response_time":		        0.01,
	"transaction_rate":		     1114.09,
	"throughput":			        0.03,
	"concurrency":			        8.63,
	"successful_transactions":	       28231,
	"failed_transactions":		        1033,
	"longest_transaction":		        0.04,
	"shortest_transaction":		        0.00
}

$ wrk -t2 -c50 -d30 'http://127.0.0.1:8080/plus?a=10&b=5'
Running 30s test @ http://127.0.0.1:8080/plus?a=10&b=5
  2 threads and 50 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    36.80ms    8.93ms 141.97ms   55.48%
    Req/Sec   675.62    176.33     0.95k    62.62%
  28257 requests in 30.03s, 5.36MB read
  Socket errors: connect 0, read 3721, write 0, timeout 0
Requests/sec:    940.99
Transfer/sec:    182.87KB
```
