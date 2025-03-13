# Golang

## Build

Change dir to the `server/golang`

### By Host-OS SoftWare

* `make build` - compile App by Host-OS software

```shell
$ go build -v -o ./server-go.bin.local ./src/main.go  && ls -al *.bin.local
-rwxrwxr-x 1 rasla rasla 7287283 мар 13 23:03 server-go.bin.local

$ make build
go build -v -o ./app.bin.local ./src/main.go \
&& chmod u+x ./app.* && ls -al app.*
-rwxrwxr-x 1 rasla rasla 7287283 мар 14 01:33 app.bin.local
```

### by Docker

* `make builder-shell` - run shell inside BUILDER docker-container
* `make build-by-docker` - make App by BUILDER docker-container
* `make build-docker` - make App docker-image

```shell
$ make build-by-docker
docker run -it --rm -u `id -u`:`id -g` -v `pwd`:/mnt -w /mnt \
        -e HOME="/mnt" \
        -e GOSUMDB="off" \
                -e GOOS=linux \
                -e GOARCH=amd64 \
                golang:1.23.3-alpine go build -v -o ./app.bin.local ./src/main.go \
&& chmod u+x ./app.* && ls -al app.*
...
-rwxr-xr-x 1 rasla rasla 7703703 мар 13 22:55 app.bin.local
```

## RUN Application

Execute in SHELL: `./app.bin.local` or `./server-go.bin.local`

* `make run` - run App in Host-OS (`make build` or `make build-by-docker` must be run first)
* `make run-by-docker` - run App in RUNNER docker-container (`make build` must be run first)
* `make run-docker` - run App docker-image (`make build-docker` must be run first)
* `make list-images` - list docker-images of this App

## CLEAN

* `make clean` - clean docker-stuff & build-artifacts

```shell
$ make clean 
docker container prune -f ; docker image prune -f ; docker volume prune -f
rm -rf .cache .config ./app.bin.local
Total reclaimed space: 0B
Total reclaimed space: 0B
Total reclaimed space: 0B
```

## Benchmarks

### Run App in Host-OS

```shell
## go version go1.23.3
# HW: Core i5-1135G7 @ 2.40GHz / 32 Gb dual-DDR4
# OS: Linux Mint 22.1 / 6.11.0-19-generic x86_64

$ siege -i -c10 -t30s -f urls-server.txt
{	"transactions":			     1174558,
	"availability":			      100.00,
	"elapsed_time":			       29.46,
	"data_transferred":		       19.81,
	"response_time":		        0.00,
	"transaction_rate":		    39869.59,
	"throughput":			        0.67,
	"concurrency":			        8.61,
	"successful_transactions":	 1174561,
	"failed_transactions":		       0,
	"longest_transaction":		    0.01,
	"shortest_transaction":		    0.00
}

$ wrk -t2 -c50 -d30s 'http://127.0.0.1:8080/plus?a=10&b=5'
Running 30s test @ http://127.0.0.1:8080/plus?a=10&b=5
  2 threads and 50 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   156.64us  182.69us   5.22ms   90.13%
    Req/Sec   146.55k     3.62k  154.13k    78.67%
  8747234 requests in 30.00s, 1.16GB read
Requests/sec: 291557.27
Transfer/sec:     39.48MB

$ fortio load -c 10 -qps 0 -t 30s 'http://127.0.0.1:8080/plus?a=10&b=5'
...
Connection time histogram (s) : count 10 avg 0.0001881338 +/- 6.766e-05 min 5.3461e-05 max 0.000271971 sum 0.001881338
# range, mid point, percentile, count
>= 5.3461e-05 <= 0.000271971 , 0.000162716 , 100.00, 10
# target 50% 0.000150577
# target 75% 0.000211274
# target 90% 0.000247692
# target 99% 0.000269543
# target 99.9% 0.000271728
Sockets used: 10 (for perfect keepalive, would be 10)
Uniform: false, Jitter: false, Catchup allowed: true
IP addresses distribution:
127.0.0.1:8080: 10
Code 200 : 5045537 (100.0 %)
Response Header Sizes : count 5045537 avg 117 +/- 0 min 117 max 117 sum 590327829
Response Body/Total Sizes : count 5045537 avg 142 +/- 0 min 142 max 142 sum 716466254
All done 5045537 calls (plus 10 warmup) 0.059 ms avg, 168184.0 qps

$ fortio load -c 20 -qps 0 -t 30s 'http://127.0.0.1:8080/plus?a=10&b=5'
...
Connection time histogram (s) : count 20 avg 0.00023447655 +/- 0.0001373 min 5.6856e-05 max 0.000536046 sum 0.004689531
# range, mid point, percentile, count
>= 5.6856e-05 <= 0.000536046 , 0.000296451 , 100.00, 20
# target 50% 0.000283841
# target 75% 0.000409943
# target 90% 0.000485605
# target 99% 0.000531002
# target 99.9% 0.000535542
Sockets used: 20 (for perfect keepalive, would be 20)
Uniform: false, Jitter: false, Catchup allowed: true
IP addresses distribution:
127.0.0.1:8080: 20
Code 200 : 6624909 (100.0 %)
Response Header Sizes : count 6624909 avg 117 +/- 0 min 117 max 117 sum 775114353
Response Body/Total Sizes : count 6624909 avg 142 +/- 0 min 142 max 142 sum 940737078
All done 6624909 calls (plus 20 warmup) 0.090 ms avg, 220829.3 qps

$ fortio load -c 100 -qps 0 -t 30s 'http://127.0.0.1:8080/plus?a=10&b=5'
...
Connection time histogram (s) : count 100 avg 0.00015840906 +/- 0.0001927 min 3.8836e-05 max 0.001849624 sum 0.015840906
# range, mid point, percentile, count
>= 3.8836e-05 <= 0.001 , 0.000519418 , 99.00, 99
> 0.001 <= 0.00184962 , 0.00142481 , 100.00, 1
# target 50% 0.000519418
# target 75% 0.000764613
# target 90% 0.00091173
# target 99% 0.001
# target 99.9% 0.00176466
Sockets used: 100 (for perfect keepalive, would be 100)
Uniform: false, Jitter: false, Catchup allowed: true
IP addresses distribution:
127.0.0.1:8080: 100
Code 200 : 7275782 (100.0 %)
Response Header Sizes : count 7275782 avg 117 +/- 0 min 117 max 117 sum 851266494
Response Body/Total Sizes : count 7275782 avg 142 +/- 0 min 142 max 142 sum 1.03316104e+09
All done 7275782 calls (plus 100 warmup) 0.412 ms avg, 242519.2 qps
```

### Run App in Docker-image

```shell
## go version go1.23.3 / runner-image: debian:12
# HW: Core i5-1135G7 @ 2.40GHz / 32 Gb dual-DDR4
# OS: Linux Mint 22.1 / 6.11.0-19-generic x86_64

$ siege -i -c10 -t30s 'http://127.0.0.1:8080/plus?a=10&b=5'
{	"transactions":			      631429,
	"availability":			      100.00,
	"elapsed_time":			       29.82,
	"data_transferred":		       15.05,
	"response_time":		        0.00,
	"transaction_rate":		    21174.68,
	"throughput":			        0.50,
	"concurrency":			        9.33,
	"successful_transactions":    631429,
	"failed_transactions":	           0,
	"longest_transaction":	        0.01,
	"shortest_transaction":         0.00
}

$ wrk -t2 -c50 -d30 'http://127.0.0.1:8080/plus?a=10&b=5'
Running 30s test @ http://127.0.0.1:8080/plus?a=10&b=5
  2 threads and 50 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   322.78us  379.62us  29.11ms   94.25%
    Req/Sec    78.93k     2.30k   84.53k    75.67%
  4712024 requests in 30.00s, 638.11MB read
Requests/sec: 157061.27
Transfer/sec:     21.27MB
```
