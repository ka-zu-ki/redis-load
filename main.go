package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/redis1", connectRedis)
	//http.HandleFunc("/redis2", connectRedis)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Println("server is called")
}

func connectRedis(w http.ResponseWriter, r *http.Request) {
	fmt.Println("redis is called")

	client, _ := newRedis()

	defer client.Close()

	a := client.SetNX(context.Background(), "afjhlsdkgjklsdgjlgla", make([]byte, 88), time.Second*10)
	fmt.Println(a)
	fmt.Println("redis is finished")
}

func newRedis() (*redis.Ring, error) {
	//var address = make(map[string]string)
	//for i, v := range strings.Split("6379", ",") {
	//	address["redis"+strconv.Itoa(i)] = v
	//}
	//
	//fmt.Println(address)

	opt := &redis.RingOptions{
		Addrs: map[string]string{
			//"redis1": "6379",
		},
	}

	cli := redis.NewRing(opt)

	//if err := cli.Ping(context.Background()).Err(); err != nil {
	//	return nil, fmt.Errorf("failed to connect to redis: %w", err)
	//}

	return cli, nil
}
