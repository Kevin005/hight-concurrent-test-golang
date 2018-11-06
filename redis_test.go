package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"testing"
)

//http://www.runoob.com/redis/redis-benchmarks.html
//测试教程地址

// redis-benchmark -n 10000  -q
//以上实例同时执行 10000 个请求来检测性能

// redis-benchmark -h 127.0.0.1 -p 6379 -t set,lpush -n 10000 -q
// 以上实例中主机为 127.0.0.1，端口号为 6379，执行的命令为 set,lpush，请求数为 10000，通过 -q 参数让结果只显示每秒执行的请求数。

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	_, err = c.Do("SET", "mykey", "superWang")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}

}

func Test_all_mark(t *testing.T) {
	fmt.Println("abcd")
}
