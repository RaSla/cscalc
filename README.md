# About

Несколько примеров Клиент-Серверной реализации калькулятора

## Server

### Python

```console
$ cd server python
## (once) - setup venv & deps
$ sudo apt install python3-venv
$ python3 -m venv venv

## Activate Venv
$ source venv/bin/activate
## (once)
$ pip3 install -r requirements.txt
$ python3 server.py
 * Running on all addresses (0.0.0.0)
 * Running on http://127.0.0.1:8080
 * Running on http://192.168.1.138:8080 (Press CTRL+C to quit)

## DeActivate Venv
$ deactivate
```

## Client

### cURL

```console
$ curl 'http://127.0.0.1:8080/plus?a=10&b=5'
a = 10, b = 5
a + b = 15

$ curl 'http://127.0.0.1:8080/minus?a=10&b=4'
a = 10, b = 4
a + b = 6

$ curl 'http://127.0.0.1:8080/multiply?a=10&b=5'
a = 10, b = 5
a * b = 50

$ curl 'http://127.0.0.1:8080/divide?a=10&b=2'
a = 10, b = 2
a / b = 5.0

$ curl 'http://127.0.0.1:8080/api/plus?a=10&b=5'
{"result": "15"}

$ curl 'http://127.0.0.1:8080/api/minus?a=10&b=4'
{"result": "6"}

$ curl 'http://127.0.0.1:8080/multiply?a=10&b=5'
{"result": "50"}

$ curl 'http://127.0.0.1:8080/divide?a=10&b=2'
{"result": "5.0"}
```

## Benchmarking

### Run test

```shell
## python 3.12.3
$ siege -f urls-server.txt -i -c10 -t30s
{	"transactions":			       68662,
	"availability":			      100.00,
	"elapsed_time":			       29.69,
	"data_transferred":		        1.35,
	"response_time":		        0.00,
	"transaction_rate":		     2312.63,
	"throughput":			        0.05,
	"concurrency":			        9.94,
	"successful_transactions":	   68662,
	"failed_transactions":		       0,
	"longest_transaction":		    0.02,
	"shortest_transaction":		    0.00
}

$ wrk -t2 -c50 -d30s 'http://127.0.0.1:8080/plus?a=10&b=5'
Running 30s test @ http://127.0.0.1:8080/plus?a=10&b=5
  2 threads and 50 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    20.44ms    1.38ms  38.23ms   85.45%
    Req/Sec     1.23k    46.62     1.31k    79.50%
  73252 requests in 30.01s, 13.90MB read
Requests/sec:   2441.03
Transfer/sec:    474.38KB

## openjdk 17.0.14
$ siege -i -c10 -t30s 'http://127.0.0.1:8080/plus?a=10&b=5'
{	"transactions":			      884136,
	"availability":			      100.00,
	"elapsed_time":			       29.09,
	"data_transferred":		        7.43,
	"response_time":		        0.00,
	"transaction_rate":		    30393.12,
	"throughput":			        0.26,
	"concurrency":			        9.24,
	"successful_transactions":	  884136,
	"failed_transactions":		       0,
	"longest_transaction":		    0.03,
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

## go version go1.23.3
$ siege -f urls-server.txt -i -c10 -t30s
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
