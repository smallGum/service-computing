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
$ ab -n 10000 -c 100 http://127.0.0.1:8080/service/userinfo
```

For *cloudgo-data-orm*:

```
$ ab -n 10000 -c 100 http://127.0.0.1:8080/service/userinfo
```

For *cloudgo-data-template*:

```
$ ab -n 10000 -c 100 http://127.0.0.1:8080/service/userinfo

```