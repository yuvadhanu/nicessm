package redis

import (
	"fmt"
	"nicessm-api-service/config"

	"os"
	"time"

	"github.com/garyburd/redigo/redis"
)

// RedisCli : ""
type RedisCli struct {
	Pool   *redis.Pool
	Config *config.ViperConfigReader
}

// ConnectionString :
const defaultConnectionString = "localhost:6379"

const redisPassword = ""

func getConnectionString(dcs string) string {

	cs := os.Getenv("REDISCONNECTIONSTRING")
	if len(cs) == 0 {
		cv := config.Config()
		cs = cv.GetString("REDISCONNECTIONSTRING")
	}

	if len(cs) == 0 {
		cs = dcs
	}
	return cs
}

// Redis POOL
func newRedisPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

var pool = newRedisPool(getConnectionString(defaultConnectionString), redisPassword)

//Connect ...
func Connect(configuration *config.ViperConfigReader) *RedisCli {
	return &RedisCli{Pool: pool, Config: configuration}
}

// SetValue :
func (cli *RedisCli) SetValue(key string, value interface{}, expiration ...interface{}) error {

	conn := cli.Pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)

	if err == nil && expiration != nil {
		conn.Do("EXPIRE", key, expiration[0])
	}
	return err
}

//GetValue ..
func (cli *RedisCli) GetValue(key string) interface{} {
	// now := time.Now()
	conn := cli.Pool.Get()
	// fmt.Println("Time taken to get redis conn from the pool: ", time.Now().Sub(now))
	defer conn.Close()

	val, err := conn.Do("GET", key)
	if err != nil {
		fmt.Println("Redis: GetValue: => ", err)
		return ""
	}
	if val == nil {
		return ""
	}
	return fmt.Sprintf("%s", val)
}

// DeleteKey ...
func (cli *RedisCli) DeleteKey(key string) {
	conn := cli.Pool.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", key)
	if err != nil {
		fmt.Println("Redis: DeleteKeyerror: ", err)

	}
}

//GetTTL : get Expiry time for a particluar key
func (cli *RedisCli) GetTTL(key string) interface{} {
	conn := cli.Pool.Get()
	defer conn.Close()

	val, err := conn.Do("TTL", key)
	if err != nil {
		fmt.Println("Redis: GetValue: ", err)
		return ""
	}
	if val == nil {
		return ""
	}
	return fmt.Sprintf("%s", val)

}

// UpdateTTL  :
func (cli *RedisCli) UpdateTTL(key string, ttl int) error {
	conn := cli.Pool.Get()
	defer conn.Close()

	_, err := conn.Do("EXPIRE", key, ttl)
	return err
}

// SETEX :
func (cli *RedisCli) SETEX(key, exp, val interface{}) error {
	conn := cli.Pool.Get()
	defer conn.Close()

	_, err := conn.Do("SETEX", key, exp, val)
	return err
}
