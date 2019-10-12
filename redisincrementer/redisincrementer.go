// Package redisincrementer contains an incrementer that persists to a Redis instance.
package redisincrementer

import "github.com/gomodule/redigo/redis"

type RedisIncrementer struct {
	pool *redis.Pool
}

func (m *RedisIncrementer) Inc() int {
	conn := m.pool.Get()
	defer conn.Close()

	r, err := redis.Int(conn.Do("INCR", "Counter"))
	if err != nil {
		panic(err)
	}

	return r
}

func NewRedisIncrementer(c *Config) *RedisIncrementer {
	ri := &RedisIncrementer{}
	ri.pool = newPool(c.RedisAddr)
	return ri
}

