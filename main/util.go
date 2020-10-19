package main

import (
	"database/sql"
	"encoding/json"
	"github.com/go-sql-driver/mysql"
	log "github.com/golang/glog"
	"github.com/gomodule/redigo/redis"
	"net"
	"net/http"
	"strconv"
	"time"
)

// NewRedisPool 实例化Redis连接池
func NewRedisPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     100,
		MaxActive:   500,
		IdleTimeout: 480 * time.Second,
		Dial: func() (redis.Conn, error) {
			timeout := time.Duration(2) * time.Second
			c, err := redis.DialTimeout("tcp", server, timeout, 0, 0)
			if err != nil {
				return nil, err
			}
			if len(password) > 0 {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
	}
}

// NewDB 实例化数据库连接池
func NewDB(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln("open mysql database:", err)
	}
	if err == mysql.ErrInvalidConn {
		log.Warning("error")
	}
	return db
}

// SuccResponse 返回成功
func SuccResponse(rw http.ResponseWriter, data interface{}) {
	rw.Header().Add("Content-Type", "application/json")
	response := map[string]interface{}{
		"code": 0,
		"msg":  "succ",
		"data": data,
	}
	res, err := json.Marshal(response)
	if err != nil {
		log.Info("json marshal:", err)
		return
	}
	_, err = rw.Write(res)
	if err != nil {
		log.Info("write err:", err)
	}
}

// FailResponse 返回失败
func FailResponse(rw http.ResponseWriter) {
	rw.Header().Add("Content-Type", "application/json")
	response := map[string]interface{}{
		"code": 1,
		"msg":  "fail",
	}
	res, err := json.Marshal(response)
	if err != nil {
		log.Info("json marshal:", err)
		return
	}
	_, err = rw.Write(res)
	if err != nil {
		log.Info("write err:", err)
	}
}

// LoggingHandler 打印请求访问日志
type LoggingHandler struct {
	handler http.Handler
}

// ServeHTTP 实现Handler接口
func (h LoggingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Infof("http request:%s %s %s", r.RemoteAddr, r.Method, r.URL)
	h.handler.ServeHTTP(w, r)
}

// HTTPHandler HTTP请求处理器
type HTTPHandler func(http.ResponseWriter, *http.Request)

// StartHTTPServer 启动HTTP服务器
func StartHTTPServer(addr string, path2handler map[string]HTTPHandler) {
	for path, handler := range path2handler {
		http.HandleFunc(path, handler)
	}
	loggingHandler := LoggingHandler{http.DefaultServeMux}
	err := http.ListenAndServe(addr, loggingHandler)
	if err != nil {
		log.Fatal("http server err:", err)
	}
}

// GetQueryParam 获取请求参数
func GetQueryParam(req *http.Request, key string) *string {
	keys, ok := req.URL.Query()[key]
	if !ok || len(keys[0]) < 1 {
		return nil
	}
	// Query()["key"] will return an array of items,
	// we only want the single item.
	val := keys[0]
	return &val
}

// GetQueryIntParam 获取请求整型参数
func GetQueryIntParam(req *http.Request, key string) (v *int) {
	val := GetQueryParam(req, key)
	if val == nil {
		return nil
	}
	intVal, err := strconv.Atoi(*val)
	if err == nil {
		v = new(int)
		*v = intVal
	}
	return
}

// GetLocalIP get local ip
func GetLocalIP() string {
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addresses {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// Contains 是否包含在列表中
func Contains(slice []string, value string) bool {
	if slice == nil || len(slice) == 0 {
		return false
	}
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

// ToJSON 转成JSON字符串
func ToJSON(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		return ""
	} 
	return string(data)
}
