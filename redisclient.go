package redisclient

// Connects to Redis DB
// v1.0.3
import (
	"bytes"
	"errors"
	"log"
	"os"

	"github.com/gomodule/redigo/redis"
)

// Client ...
var Client redis.Conn

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger: ", log.Lshortfile+log.Ldate+log.Ltime)
)

/*
RedisConnection ...
*/
func RedisConnection() error {

	log.Print("Redis init ")
	var err error

	host := os.Getenv("REDIS_SERVER")
	if host == "" {
		host = "localhost:6379"
	}

	Client, err = redis.DialURL("redis://" + host)
	if err != nil {
		return err
	}
	log.Print("Redis init done")
	return nil
}

/*
SetEx ... Set with Expiry time
*/
func SetEx(key, value string, time int) (interface{}, error) {
	resp, err := Client.Do("SETEX", key, time, value)
	return resp, err
}

/*
Set ...
*/
func Set(key, value string) (interface{}, error) {

	resp, err := Client.Do("SET", key, value)
	return resp, err
}

/*
Get ...
*/
func Get(key string) (interface{}, error) {
	resp, err := Client.Do("GET", key)
	if err == nil {
		if resp == nil {
			err = errors.New("Unable for find key: " + key)
		}
	}
	return resp, err
}

/*
Del ...
*/
func Del(key string) (interface{}, error) {
	resp, err := Client.Do("DEL", key)
	if err == nil {
		if resp == nil {
			err = errors.New("Unable to delete : " + key)
		}
	}
	return resp, err
}

/*
Auth ...
*/
func Auth() error {
	_, err := redis.String(Client.Do("Auth", "Passw0rd"))
	if err != nil {
		return err
	}
	return nil

}

/*
Ping ...
*/
func Ping() error {
	_, err := redis.String(Client.Do("PING"))
	if err != nil {
		return err
	}
	//fmt.Printf("Ping Response = %s\n", pong)
	return nil
}
