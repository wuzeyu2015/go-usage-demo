package main

import (
	"fmt"
	"reflect"
)

type IRoute interface {
	test()
	test1()
	test2()
}

type Common struct {
}

func (c *Common) test() {
	fmt.Println("test")
}

func (c *Common) test1() {
	fmt.Println("test1")
}

func (c *Common) test2() {
	fmt.Println("test2")
}

type Login struct {
	Common
}

func (l *Login) test() {
	fmt.Println("Login test ---------")
}

type Auth struct {
	Common
}

func (a *Auth) test() {
	fmt.Println("Auth test ------------")
}

func (a *Auth) test1() {
	fmt.Println("Auth test1 -----------")
}

func addroute(route IRoute) {
	route.test()
	route.test1()
	route.test2()
}

var RegisterMessage = make(map[string]interface{})

func init() {
	RegisterMessage["login"] = &Login{}
	RegisterMessage["auth"] = &Auth{}
}

func main() {
	msg := RegisterMessage["login"]
	t := reflect.TypeOf(msg).Elem()
	n := reflect.New(t).Interface().(IRoute)
	addroute(n)

	msg1 := RegisterMessage["auth"]
	t1 := reflect.TypeOf(msg1).Elem()
	n1 := reflect.New(t1).Interface().(IRoute)
	addroute(n1)
}
