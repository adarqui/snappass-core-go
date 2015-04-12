package snappass_core

import (
	"github.com/garyburd/redigo/redis"
)

type RedisDatabase struct {
	address string
	auth    string
	db      int
	conn    redis.Conn
}

func NewRedisDatabase(address, auth string, db int) (*RedisDatabase, error) {
	redisdb := new(RedisDatabase)
	redisdb.address = address
	redisdb.auth = auth
	redisdb.db = db
	return redisdb, nil
}

func (redisdb *RedisDatabase) connect() error {
	c, err := redis.Dial("tcp", redisdb.address)
	if err != nil {
		return err
	}
	redisdb.conn = c
	if redisdb.auth != "" {
		_, err := redisdb.conn.Do("AUTH", redisdb.auth)
		if err != nil {
			redisdb.conn.Close()
			return err
		}
	}
	if redisdb.db != 0 {
		_, err := redisdb.conn.Do("SELECT", redisdb.db)
		if err != nil {
			redisdb.conn.Close()
			return err
		}
	}
	return nil
}

func (redisdb *RedisDatabase) destroy() error {
	redisdb.conn.Close()
	return nil
}

func (redisdb *RedisDatabase) setAndExpire(key, value []byte, ttl TTL) error {
	_, err := redisdb.conn.Do("SETNX", key, value)
	if err != nil {
		return err
	}
	_, err_expire := redisdb.conn.Do("EXPIRE", key, ttl)
	if err_expire != nil {
		return err
	}
	return nil
}

func (redisdb *RedisDatabase) getAndDelete(key []byte) ([]byte, error) {
	value, err := redisdb.conn.Do("GET", key)
	if err != nil {
		return nil, err
	}
	_, err_del := redisdb.conn.Do("DEL", key)
	if err_del != nil {
		return nil, err_del
	}
	value_bs, err_bytes := redis.Bytes(value, nil)
	return value_bs, err_bytes
}

func (redisdb *RedisDatabase) isKeySet(key []byte) (bool, error) {
	value, err := redisdb.conn.Do("EXISTS", key)
	if err != nil {
		return false, err
	}
	value_bool, err_bool := redis.Bool(value, nil)
	return value_bool, err_bool
}
