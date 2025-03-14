# Rust

## Build

Change dir to the `server/rust`

### By Host-OS SoftWare

* `make build` - compile App by Host-OS software

```shell
$ time cargo build --release  && ls -al target/*/rust_server
...
   Compiling rust_server v0.1.0 (/home/rasla/git/_ra/cscalc/server/rust)
    Finished `release` profile [optimized] target(s) in 25.07s

real    0m25,097s
user    2m14,141s
sys     0m8,528s
-rwxrwxr-x 2 rasla rasla 71693240 мар 14 17:59 target/debug/rust_server
-rwxrwxr-x 2 rasla rasla  3452192 мар 14 18:21 target/release/rust_server

$ make build
time cargo build --release  && ls -al target/*/rust_server \
&& mv target/release/rust_server ./app.bin.local && ls -al ./app.bin.local
    Finished `release` profile [optimized] target(s) in 0.04s
0.03user 0.03system 0:00.06elapsed 101%CPU (0avgtext+0avgdata 42724maxresident)k
0inputs+0outputs (0major+7246minor)pagefaults 0swaps
-rwxrwxr-x 2 rasla rasla 71693240 мар 14 17:59 target/debug/rust_server
-rwxrwxr-x 2 rasla rasla  3452192 мар 14 18:21 target/release/rust_server
-rwxrwxr-x 2 rasla rasla 3452192 мар 14 18:21 ./app.bin.local
```

## RUN Application

Execute in SHELL: `./app.bin.local` or `./target/release/rust_server`

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
## rustc 1.85.0 (4d91de4e4 2025-02-17)
# HW: Core i5-1135G7 @ 2.40GHz / 32 Gb dual-DDR4
# OS: Linux Mint 22.1 / 6.11.0-19-generic x86_64

$ siege -i -c10 -t30s -f urls-server.txt
{	"transactions":			     1311428,
	"availability":			      100.00,
	"elapsed_time":			       29.29,
	"data_transferred":		       22.11,
	"response_time":		        0.00,
	"transaction_rate":		    44773.91,
	"throughput":			        0.76,
	"concurrency":			        8.67,
	"successful_transactions":   1311429,
	"failed_transactions":	           0,
	"longest_transaction":	        0.01,
	"shortest_transaction":	        0.00
}

$ wrk -t2 -c50 -d30 'http://127.0.0.1:8080/plus?a=10&b=5'
Running 30s test @ http://127.0.0.1:8080/plus?a=10&b=5
  2 threads and 50 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   108.30us   68.59us   0.99ms   88.10%
    Req/Sec   156.21k     4.67k  170.08k    69.44%
  9355965 requests in 30.10s, 1.23GB read
Requests/sec: 310831.25
Transfer/sec:   41.80MB

$ fortio load -c 10 -qps 0 -t 30s 'http://127.0.0.1:8080/plus?a=10&b=5'
...
Connection time histogram (s) : count 10 avg 0.0001462509 +/- 6.16e-05 min 5.7049e-05 max 0.000237084 sum 0.001462509
# range, mid point, percentile, count
>= 5.7049e-05 <= 0.000237084 , 0.000147067 , 100.00, 10
# target 50% 0.000137065
# target 75% 0.000187074
# target 90% 0.00021708
# target 99% 0.000235084
# target 99.9% 0.000236884
Sockets used: 10 (for perfect keepalive, would be 10)
Uniform: false, Jitter: false, Catchup allowed: true
IP addresses distribution:
127.0.0.1:8080: 10
Code 200 : 6312462 (100.0 %)
Response Header Sizes : count 6312462 avg 116 +/- 0 min 116 max 116 sum 732245592
Response Body/Total Sizes : count 6312462 avg 141 +/- 0 min 141 max 141 sum 890057142
All done 6312462 calls (plus 10 warmup) 0.047 ms avg, 210414.6 qps

$ fortio load -c 20 -qps 0 -t 30s 'http://127.0.0.1:8080/plus?a=10&b=5'
...
Connection time histogram (s) : count 20 avg 8.53211e-05 +/- 5.629e-05 min 4.2338e-05 max 0.000207296 sum 0.001706422
# range, mid point, percentile, count
>= 4.2338e-05 <= 0.000207296 , 0.000124817 , 100.00, 20
# target 50% 0.000120476
# target 75% 0.000163886
# target 90% 0.000189932
# target 99% 0.00020556
# target 99.9% 0.000207122
Sockets used: 20 (for perfect keepalive, would be 20)
Uniform: false, Jitter: false, Catchup allowed: true
IP addresses distribution:
127.0.0.1:8080: 20
Code 200 : 8131851 (100.0 %)
Response Header Sizes : count 8131851 avg 116 +/- 0 min 116 max 116 sum 943294716
Response Body/Total Sizes : count 8131851 avg 141 +/- 0 min 141 max 141 sum 1.14659099e+09
All done 8131851 calls (plus 20 warmup) 0.074 ms avg, 271060.5 qps

$ fortio load -c 100 -qps 0 -t 30s 'http://127.0.0.1:8080/plus?a=10&b=5'
...
Connection time histogram (s) : count 100 avg 0.00025279134 +/- 0.000252 min 4.1602e-05 max 0.001127875 sum 0.025279134
# range, mid point, percentile, count
>= 4.1602e-05 <= 0.001 , 0.000520801 , 97.00, 97
> 0.001 <= 0.00112787 , 0.00106394 , 100.00, 3
# target 50% 0.000530784
# target 75% 0.000780367
# target 90% 0.000930117
# target 99% 0.00108525
# target 99.9% 0.00112361
Sockets used: 100 (for perfect keepalive, would be 100)
Uniform: false, Jitter: false, Catchup allowed: true
IP addresses distribution:
127.0.0.1:8080: 100
Code 200 : 10964422 (100.0 %)
Response Header Sizes : count 10964422 avg 116 +/- 0 min 116 max 116 sum 1.27187295e+09
Response Body/Total Sizes : count 10964422 avg 141 +/- 1.668e-06 min 141 max 141 sum 1.5459835e+09
All done 10964422 calls (plus 100 warmup) 0.273 ms avg, 365475.0 qps
```

### Run App in Docker-image

```shell
## rustc 1.85.0 (4d91de4e4 2025-02-17)
# HW: Core i5-1135G7 @ 2.40GHz / 32 Gb dual-DDR4
# OS: Linux Mint 22.1 / 6.11.0-19-generic x86_64

$ make run-by-docker
docker run -it --rm -u `id -u`:`id -g` -v `pwd`:/mnt -w /mnt \
        -e HOME="/mnt" \
        -p 8080:8080 \
        debian:12 ./app.bin.local
Server is running on port 8080...

$ curl 'http://127.0.0.1:8080/minus?a=10&b=4'
curl: (56) Recv failure: Соединение разорвано другой стороной
```
