# cloudgo

In this project, I design a Golang-web-service program named *cloudgo*.

## Introduction

*cloudgo* is a simple http web server which receives URL from clients and returns corresponding response to clients according to the query parameters of the URL. Without using any external web frameworks, I use only `net/http` packet of Golang to implement the program.

## Build

For building the program, just download the service-computing repository into your directory `$GOPATH/src/`, and then execute build command on your terminal:

```
$ go build
```

## Usage

For the server:

```
$ ./cloudgo [-p port]

-p string
  the port for cloudgo server to listen.
  if not provided, the server will listen to your system environment variable DEFAULT_PORT or the default port 23333.
```

For the client:
there are only two kinds of URL the client can use to send http request to the server:

```
http://127.0.0.1:PORT/hello
http://127.0.0.1:PORT/hello?name=YOUR_NAME

PORT is the port the server is listening to (default 23333).
YOUR_NAME can be any string.
```

## Curl Test

URL without name parameter:

```
$ curl -v http://127.0.0.1:23333/hello
*   Trying 127.0.0.1...
* Connected to 127.0.0.1 (127.0.0.1) port 23333 (#0)
> GET /hello HTTP/1.1
> Host: 127.0.0.1:23333
> User-Agent: curl/7.47.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Date: Tue, 14 Nov 2017 11:24:59 GMT
< Content-Length: 18
< Content-Type: text/plain; charset=utf-8
<
Hello, my client!
* Connection #0 to host 127.0.0.1 left intact
```

URL with name parameter:

```
$ curl -v http://127.0.0.1:23333/hello?name=Jack
*   Trying 127.0.0.1...
* Connected to 127.0.0.1 (127.0.0.1) port 23333 (#0)
> GET /hello?name=Jack HTTP/1.1
> Host: 127.0.0.1:23333
> User-Agent: curl/7.47.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Date: Tue, 14 Nov 2017 11:26:28 GMT
< Content-Length: 13
< Content-Type: text/plain; charset=utf-8
<
Hello, Jack!
* Connection #0 to host 127.0.0.1 left intact
```

## ApacheBench Test

```
$ ab -n 4000 -c 1000 http://127.0.0.1:23333/hello
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 400 requests
Completed 800 requests
Completed 1200 requests
Completed 1600 requests
Completed 2000 requests
Completed 2400 requests
Completed 2800 requests
Completed 3200 requests
Completed 3600 requests
Completed 4000 requests
Finished 4000 requests


Server Software:        
Server Hostname:        127.0.0.1
Server Port:            23333

Document Path:          /hello
Document Length:        18 bytes

Concurrency Level:      1000
Time taken for tests:   0.457 seconds
Complete requests:      4000
Failed requests:        0
Total transferred:      540000 bytes
HTML transferred:       72000 bytes
Requests per second:    8762.07 [#/sec] (mean)
Time per request:       114.128 [ms] (mean)
Time per request:       0.114 [ms] (mean, across all concurrent requests)
Transfer rate:          1155.16 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    4  10.4      0      37
Processing:     0   17  11.1     15     214
Waiting:        0   17  11.1     15     214
Total:          0   22  16.5     16     215

Percentage of the requests served within a certain time (ms)
  50%     16
  66%     18
  75%     20
  80%     30
  90%     49
  95%     55
  98%     66
  99%     75
 100%    215 (longest request)
```

## Design details

```
+----------------------------+            +-----------------------+                  
| get port from command line |            | create a new router   |
|    if not provided, use    |            | use function sayHello |
|       default port         |            |      as handler       |
+----------------------------+            +-----------------------+
              |                                       |
              | port                                  | mux.GetMux()
              |                                       |
              -----------------------------------------
                                  |
                                  |
                                  V
                      +----------------------------+
                      | use port and the router to |
                      |  create a new http.Server  |
                      +----------------------------+
                                  |
                                  | server
                                  V
                      +----------------------------+
                      |   server.ListenAndServe()  |
                      +----------------------------+         
```
