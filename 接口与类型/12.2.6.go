package main

import "fmt"

type IDatabaser interface {
	Connect() error
	Disconnect() error
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
	var redis = Redis{DBName:"teacher"}
	var idb IDatabaser = &redis
	idb.Connect()
	idb.Disconnect()
}