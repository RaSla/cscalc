# Java

## Build

Change dir to the `server/java`

### By Host-OS SoftWare

* `make build` - compile App by Host-OS software

### by Docker

* `make builder-shell` - run shell inside BUILDER docker-container
* `make build-by-docker` - make App by BUILDER docker-container
* `make build-docker` - make App docker-image

## RUN Application

* `make run` - run App in Host-OS (`make build` or `make build-by-docker` must be run first)
* `make run-by-docker` - run App in RUNNER docker-container (`make build` must be run first)
* `make run-docker` - run App docker-image (`make build-docker` must be run first)
* `make list-images` - list docker-images of this App

## CLEAN

* `make clean` - clean docker-stuff & build-artifacts

## Benchmarks

### Run App in Host-OS

```shell
## openjdk 17.0.14
# HW: Core i5-1135G7 @ 2.40GHz / 32 Gb dual-DDR4
# OS: Linux Mint 22.1 / 6.11.0-19-generic x86_64

$ siege -i -c10 -t30s -f urls-server.txt
{	"transactions":			      889777,
	"availability":			      100.00,
	"elapsed_time":			       29.32,
	"data_transferred":		        7.48,
	"response_time":		        0.00,
	"transaction_rate":		    30347.10,
	"throughput":			        0.25,
	"concurrency":			        9.25,
	"successful_transactions":	  889777,
	"failed_transactions":		       0,
	"longest_transaction":          0.02,
	"shortest_transaction":		    0.00
}


$ wrk -t2 -c50 -d30 'http://127.0.0.1:8080/plus?a=10&b=5'
Running 30s test @ http://127.0.0.1:8080/plus?a=10&b=5
  2 threads and 50 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.06ms  455.16us  21.33ms   84.23%
    Req/Sec    21.96k     2.71k   27.07k    79.83%
  1310575 requests in 30.00s, 132.49MB read
Requests/sec:  43684.87
Transfer/sec:      4.42MB
```

### Run App in Docker-image

```shell
## openjdk 17.0.14
# HW: Core i5-1135G7 @ 2.40GHz / 32 Gb dual-DDR4
# OS: Linux Mint 22.1 / 6.11.0-19-generic x86_64

$ siege -i -c10 -t30s 'http://127.0.0.1:8080/plus?a=10&b=5'
{	"transactions":			      450406,
	"availability":			      100.00,
	"elapsed_time":			       29.31,
	"data_transferred":		        0.86,
	"response_time":		        0.00,
	"transaction_rate":		    15366.97,
	"throughput":			        0.03,
	"concurrency":			        9.37,
	"successful_transactions":    450406,
	"failed_transactions":	           0,
	"longest_transaction":	        0.02,
	"shortest_transaction":         0.00
}

$ wrk -t2 -c50 -d30 'http://127.0.0.1:8080/plus?a=10&b=5'
Running 30s test @ http://127.0.0.1:8080/plus?a=10&b=5
  2 threads and 50 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.70ms  498.94us  19.59ms   79.95%
    Req/Sec    13.47k   413.16    14.39k    77.17%
  804299 requests in 30.01s, 81.31MB read
Requests/sec:  26804.37
Transfer/sec:      2.71MB
```
