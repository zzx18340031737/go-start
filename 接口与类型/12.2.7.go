package main

import "fmt"

type IDatabaser interface {
	Connect() error
	Disconnect() error
}

type IRediser interface {
	Connect() error
}

type Redis struct {
	DBName string
}

func (redis *Redis)Connect() error {
	fmt.Println("Redis Connect DB => "+redis.DBName)
	//do Connect
	fmt.Println("Redis Connect Success!")
	return nil
}

func (redis *Redis)Disconnect() error{
	//do Disconnect
	fmt.Println("Redis Disconnect Success!")
	return nil
}

func main(){
	var idb IDatabaser = &Redis{DBName:"teacher"}

	var iredis IRediser
	iredis = idb
	iredis.Connect()
}