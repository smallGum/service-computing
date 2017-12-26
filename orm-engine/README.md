# ORM-Engine

*orm-engine* is a small engine used to simplify database operations. Using reflection in golang, it implements both `Insert()` and `Find()` functions to help user insert a new data entry to the database and query all entries from corresponding table.

## Task

### 1. orm rules and table

We use `UserInfo` struct to create database table and test our program, which is defined as follows:

```go
// UserInfo .
type UserInfo struct {
    UID        int   `orm:"id,auto-inc,type=INT(10)"` // tag
    UserName   string
    DepartName string
    CreateAt   *time.Time `orm:"name=created" json:",omitempty"`
}
```

*tags* of each column describes all the attributes of that column. According to the `UserInfo` struct and its tags, we create `userinfo` table in database `test`:

```sql
CREATE DATABASE test;
USE test;

CREATE TABLE `userinfo` (
    `uid` INT(10) NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(64) NULL DEFAULT NULL,
    `departname` VARCHAR(64) NULL DEFAULT NULL,
    `createat` DATE NULL DEFAULT NULL,
    PRIMARY KEY (`uid`)
);
```

After creation, we can see the table in `test` database:

```sql
mysql> DESCRIBE userinfo;

+------------+-------------+------+-----+---------+----------------+
| Field      | Type        | Null | Key | Default | Extra          |
+------------+-------------+------+-----+---------+----------------+
| uid        | int(10)     | NO   | PRI | NULL    | auto_increment |
| username   | varchar(64) | YES  |     | NULL    |                |
| departname | varchar(64) | YES  |     | NULL    |                |
| createat   | date        | YES  |     | NULL    |                |
+------------+-------------+------+-----+---------+----------------+
4 rows in set (0.00 sec)
```

### 2. auto-insertion

We implement function `Insert(o interface{}) (int, error)` to realize database auto-insertion of our *orm-engine*. It should work like this:

```go
user := UserInfo{...}
affected, err := engine.Insert(user)
// INSERT INTO user (name) values (?)
```

And our test code in `main.go` is:

```go
engine := entities.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
t := time.Now()
u1 := entities.UserInfo{
	UserName:   "Jack",
	DepartName: "Software",
	CreateAt:   &t,
}
t = time.Now()
u2 := entities.UserInfo{
	UserName:   "Lucy",
	DepartName: "Management",
	CreateAt:   &t,
}

affected, err := engine.Insert(u1)
if err != nil {
	panic(err)
}
fmt.Printf("%d row(s) affected after inserting %s\n", affected, u1.UserName)
affected, err = engine.Insert(u2)
if err != nil {
	panic(err)
}
fmt.Printf("%d row(s) affected after inserting %s\n", affected, u2.UserName)
```

Run the code, and we can see the result:

```shell
1 row(s) affected after inserting Jack
1 row(s) affected after inserting Lucy
```

```sql
mysql> SELECT * FROM userinfo;

+-----+----------+------------+------------+
| uid | username | departname | createat   |
+-----+----------+------------+------------+
|   1 | Jack     | Software   | 2017-12-26 |
|   2 | Lucy     | Management | 2017-12-26 |
+-----+----------+------------+------------+
2 rows in set (0.00 sec)
```

As we can see, the two new users `Jack` and `Lucy` are inserted into the `userinfo` table successfully.

### 3. query result to object

We implement function `Find(o interface{}) error`, which queries all the entries from the database and transform this entries into underlying object defined by null interface `o`. It should work like this:

```go
pEveryOne := make([]*Userinfo, 0)
err := engine.Find(&pEveryOne)
// SELECT `col-name`,`col-name` ... FROM UserInfo
```

Our test code in `main.go` is:

```go
pEveryOne := make([]*entities.UserInfo, 0)
err = engine.Find(&pEveryOne)
if err != nil {
	panic(err)
}
fmt.Println("all users: ")
for i := 0; i < len(pEveryOne); i++ {
	fmt.Println(*pEveryOne[i])
}
```

Run the code, we can get all users information:

```shell
all users:
{1 Jack Software 2017-12-26 00:00:00 +0000 UTC}
{2 Lucy Management 2017-12-26 00:00:00 +0000 UTC}
```

## Conclusion

*orm-engine* is just a small toy of the **ORM framework**. However, it shows the power of reflection in golang. Using reflection, we can not only get table's name, column's name and value of an objects but also read the table entries from the database and change it back into original objects. Through this project, I have got a deeper insight into [The Laws of Reflection](https://blog.golang.org/laws-of-reflection).