// Redis数据库操作
type Redis struct {
	DBName string
}

func (redis *Redis)Connect() error{
fmt.Println("Redis Connect DB =>" +redis.DBName)
//do Connect
fmt.Println("Redis Connect Success!")
return nil
}

func (redis *Redis)Disconnect() error{
	//do DisConnect
	fmt.Println("Redis DisConnect Success!")
	return nil
}
