package main

import (
	"fmt"
	"service-computing/orm-engine/entities"
	"time"
)

func main() {
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

	pEveryOne := make([]*entities.UserInfo, 0)
	err = engine.Find(&pEveryOne)
	if err != nil {
		panic(err)
	}
	fmt.Println("all users: ")
	for i := 0; i < len(pEveryOne); i++ {
		fmt.Println(*pEveryOne[i])
	}
}
