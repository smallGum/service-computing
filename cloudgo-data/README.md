# cloudgo-data

## Introduction

*cloudgo-data* implements three database service programs *cloudgo-data-dbsql*, *cloudgo-data-orm* and *cloudgo-data-template* by Golang. The service provided by three programs are almost the same:

```sql
INSERT userinfo SET username=?,departname=?,created=?
```

```sql
SELECT * FROM userinfo
```

```sql
SELECT * FROM userinfo where uid = ?
```

while *cloudgo-data-template* implements an extra service:

```sql
SELECT count(*) FROM userinfo
```

The only difference is that *cloudgo-data-dbsql* uses Go's native library `database/sql` to implement those service, while *cloudgo-data-orm* uses `xorm`, an ORM framework for Go and *cloudgo-data-template* implements a self-define template `mysqlt` for operations of the database service and uses it to implements those service.
(**Note:** The *cloudgo-data-dbsql* program are cloned from [pml's repository](https://github.com/pmlpml/golang-learning/tree/master/web/cloudgo-data))

We use *Apache Bench* to test the performances of the three programs and compare them with each other.

## Service Implementation

Select all users:

```shell
$ curl http://localhost:8080/service/userinfo?userid=

[
  {
    "UID": 1,
    "UserName": "ooo",
    "DepartName": "1",
    "CreateAt": "2017-11-27T00:00:00Z"
  },
  {
    "UID": 2,
    "UserName": "oop",
    "DepartName": "2",
    "CreateAt": "2017-11-28T00:00:00Z"
  },
  {
    "UID": 3,
    "UserName": "Jack",
    "DepartName": "16",
    "CreateAt": "2017-11-29T00:00:00Z"
  }
]
```

Select one user by id:

```shell
$ curl http://localhost:8080/service/userinfo?userid=2

{
  "UID": 2,
  "UserName": "oop",
  "DepartName": "2",
  "CreateAt": "2017-11-28T00:00:00Z"
}
```

Insert a new user:

```shell
$ curl -d "username=Tracy&departname=5" http://localhost:8080/service/userinfo

{
  "UID": 4,
  "UserName": "Tracy",
  "DepartName": "5",
  "CreateAt": "2017-11-29T23:17:58.31375+08:00"
}
```

Count the number of users (this is only provided by *cloudgo-data-template*):

```shell
$ curl http://localhost:8080/service/usercount

4
```

## Apache Bench Test

For *cloudgo-data-dbsql*:

```
$ ab -n 10000 http://127.0.0.1:8080/service/userinfo

This is ApacheBench, Version 2.3 <$Revision: 1796539 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        127.0.0.1
Server Port:            8080

Document Path:          /service/userinfo
Document Length:        3748 bytes

Concurrency Level:      1
Time taken for tests:   4.539 seconds
Complete requests:      10000
Failed requests:        10
   (Connect: 0, Receive: 0, Length: 10, Exceptions: 0)
Non-2xx responses:      10000
Total transferred:      38640010 bytes
HTML transferred:       37480010 bytes
Requests per second:    2203.29 [#/sec] (mean)
Time per request:       0.454 [ms] (mean)
Time per request:       0.454 [ms] (mean, across all concurrent requests)
Transfer rate:          8313.99 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       4
Processing:     0    0   0.4      0      18
Waiting:        0    0   0.4      0      18
Total:          0    0   0.5      0      18

Percentage of the requests served within a certain time (ms)
  50%      0
  66%      0
  75%      0
  80%      0
  90%      1
  95%      1
  98%      1
  99%      1
 100%     18 (longest request)
```

For *cloudgo-data-orm*:

```
$ ab -n 10000 -c 100 http://127.0.0.1:8080/service/userinfo

This is ApacheBench, Version 2.3 <$Revision: 1796539 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        127.0.0.1
Server Port:            8080

Document Path:          /service/userinfo
Document Length:        3744 bytes

Concurrency Level:      1
Time taken for tests:   4.614 seconds
Complete requests:      10000
Failed requests:        0
Non-2xx responses:      10000
Total transferred:      38600000 bytes
HTML transferred:       37440000 bytes
Requests per second:    2167.20 [#/sec] (mean)
Time per request:       0.461 [ms] (mean)
Time per request:       0.461 [ms] (mean, across all concurrent requests)
Transfer rate:          8169.32 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0      11
Processing:     0    0   0.8      0      71
Waiting:        0    0   0.8      0      71
Total:          0    0   0.9      0      71

Percentage of the requests served within a certain time (ms)
  50%      0
  66%      0
  75%      0
  80%      0
  90%      1
  95%      1
  98%      1
  99%      1
 100%     71 (longest request)
```

For *cloudgo-data-template*:

```
$ ab -n 10000 -c 100 http://127.0.0.1:8080/service/userinfo

This is ApacheBench, Version 2.3 <$Revision: 1796539 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        127.0.0.1
Server Port:            8080

Document Path:          /service/userinfo
Document Length:        3754 bytes

Concurrency Level:      1
Time taken for tests:   4.743 seconds
Complete requests:      10000
Failed requests:        0
Non-2xx responses:      10000
Total transferred:      38700000 bytes
HTML transferred:       37540000 bytes
Requests per second:    2108.40 [#/sec] (mean)
Time per request:       0.474 [ms] (mean)
Time per request:       0.474 [ms] (mean, across all concurrent requests)
Transfer rate:          7968.28 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0      12
Processing:     0    0   0.5      0      16
Waiting:        0    0   0.4      0      16
Total:          0    0   0.5      0      16

Percentage of the requests served within a certain time (ms)
  50%      0
  66%      0
  75%      0
  80%      0
  90%      1
  95%      1
  98%      1
  99%      2
 100%     16 (longest request)
```

## Compare

From the result of Apache Bench test, focusing on the value of **Transfer rate**, **Requests per second** and so on, we can conclude that the ORM framework `xorm` does have some negative influence on the performance of the database service comparing to the native library `database/sql` while it's performance is better than our self-define sql template `mysqlt`. 

However, the negative effect is not serious, besides, using the ORM framework can greatly improve coding efficiency and make database more maintainable since it greatly simplifies the **DAO** service of the database. 

Thus in my opinion, we should use ORM framework to simplify our work as long as we don't require much high performance.