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

## Nuitka

Compiler for Python-apps - create single binary file

### Install

```shell
$ sudo apt install python3-dev patchelf ccache
$ source venv/bin/activate
$ (venv) pip3 install nuitka==2.6.8
$ (venv) python -m nuitka --version
2.6.8
Commercial: None
Python: 3.12.3 (main, Feb  4 2025, 14:48:35) [GCC 13.3.0]
Flavor: Debian Python
Executable: ~/git/_ra/cscalc/server/python/venv/bin/python
OS: Linux
Arch: x86_64
Distribution: Linuxmint (based on Ubuntu) 22.1
Version C compiler: /usr/bin/gcc (gcc 13).
```

### Usage

```shell
$ (venv) time nuitka --standalone --onefile --output-dir=bin server.py
Nuitka-Options: Used command line options: --standalone --onefile --output-dir=bin server.py
Nuitka: Starting Python compilation with Nuitka '2.6.8' on Python (flavor Debian Python), '3.12' commercial grade 'not installed'.
Nuitka-Plugins:anti-bloat: Not including '_json' automatically in order to avoid bloat, but this may cause: may slow down by using fallback
Nuitka-Plugins:anti-bloat: implementation.
Nuitka: Completed Python level compilation and optimization.
Nuitka: Generating source code for C backend compiler.
Nuitka: Running data composer tool for optimal constant value handling.                            
Nuitka: Running C compilation via Scons.
Nuitka-Scons: Backend C compiler: gcc (gcc 13).
Nuitka-Scons: Backend C linking with 128 files (no progress information available for this stage).
Nuitka-Scons: Compiled 127 C files using ccache.
Nuitka-Scons: Cached C files (using ccache) with result 'cache hit': 127
Nuitka-Postprocessing: Creating single file from dist folder, this may take a while.
Nuitka-Onefile: Running bootstrap binary compilation via Scons.
Nuitka-Scons: Onefile C compiler: gcc (gcc 13).
Nuitka-Scons: Onefile C linking.              
Nuitka-Scons: Compiled 1 C files using ccache.
Nuitka-Scons: Cached C files (using ccache) with result 'cache hit': 1
Nuitka-Onefile: Using compression for onefile payload.
Nuitka-Onefile: Onefile payload compression ratio (27.49%) size 39298889 to 10801829.
Nuitka-Onefile: Keeping onefile build directory 'bin/server.onefile-build'.      
Nuitka: Keeping dist folder 'bin/server.dist' for inspection, no need to use it.
Nuitka: Keeping build directory 'bin/server.build'.
Nuitka: Successfully created 'bin/server.bin'.
real    2m28,921s
user    10m3,371s
sys     0m16,175s

$ (venv) ls -al bin/
-rwxrwxr-x 1 rasla rasla 10980864 мар 15 05:14 server.bin
drwxrwxr-x 3 rasla rasla    20480 мар 15 05:14 server.build
drwxrwxr-x 3 rasla rasla     4096 мар 15 05:14 server.dist
drwxrwxr-x 3 rasla rasla     4096 мар 15 05:14 server.onefile-build

$ ./bin/server.bin
```

### Benchmarks

```shell
### Native python3 server.py
$ wrk -t2 -c50 -d30 'http://127.0.0.1:8080/plus?a=10&b=5'
Running 30s test @ http://127.0.0.1:8080/plus?a=10&b=5
  2 threads and 50 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    23.94ms    0.98ms  38.31ms   75.93%
    Req/Sec     1.05k    23.17     1.12k    78.17%
  62532 requests in 30.01s, 11.87MB read
Requests/sec:   2083.83
Transfer/sec:    404.97KB

### Nuitka 2.6.8 - server.bin
$ wrk -t2 -c50 -d30 'http://127.0.0.1:8080/plus?a=10&b=5'
Running 30s test @ http://127.0.0.1:8080/plus?a=10&b=5
  2 threads and 50 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    29.72ms    2.07ms  60.60ms   89.95%
    Req/Sec   843.49     45.44     0.93k    86.33%
  50376 requests in 30.01s, 9.56MB read
Requests/sec:   1678.85
Transfer/sec:    326.26KB

### PyPy 7.3.15 - Python 3.9
$ wrk -t2 -c50 -d30 'http://127.0.0.1:8080/plus?a=10&b=5'
Running 30s test @ http://127.0.0.1:8080/plus?a=10&b=5
  2 threads and 50 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   478.43ms   64.72ms 677.19ms   74.78%
    Req/Sec    51.93     16.96   111.00     62.17%
  3116 requests in 30.03s, 605.55KB read
Requests/sec:    103.75
Transfer/sec:     20.16KB
```

## PyPy3

### Install

```shell
$ sudo apt install pypy3 pypy3-venv
$ pypy3 -m venv venv-pypy3
$ source venv-pypy3/bin/activate
$ (venv-pypy3) python3 --version
Python 3.9.18 (7.3.15+dfsg-1build3, Apr 01 2024, 03:12:48)
[PyPy 7.3.15 with GCC 13.2.0]
$ (venv-pypy3) pip3 install -r requirements.txt
...

$ (venv-pypy3) pypy3 server.py
```
