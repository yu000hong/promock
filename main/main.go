package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("usage: ./promock CONFIG-FILE")
		return
	}

	config := ReadConfig(flag.Args()[0])
	redisPool := NewRedisPool(config.RedisAddress, config.RedisPassword)
	defer func() {
		if redisPool != nil {
			_ = redisPool.Close()
		}
	}()
	db := NewDB(config.MysqlDSN)
	defer func() {
		if db != nil {
			_ = db.Close()
		}
	}()

	// 启动HTTP服务器
	path2handler := make(map[string]HTTPHandler)
	path2handler["/"] = rootHandler
	StartHTTPServer(config.HTTPServerAddress, path2handler)
}

func rootHandler(rw http.ResponseWriter, req * http.Request) {
	ServeProxy(rw, req)
}
