package service

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/toaddyan/redis-slinker/pkg/base62"
	"github.com/toaddyan/redis-slinker/pkg/storage"
)

type database struct {
	pool *redis.Pool
}

func NewPool(host, port, password string) (*redis.Pool, error) {
	return &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			fmt.Println("host and port ", host, port)
			return redis.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
		},
	}, nil
}

func NewService(pool *redis.Pool) storage.Service {
	return &database{pool: pool}
}

func (r *database) isUsed(id uint64) bool {
	conn := r.pool.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", "Shortener:"+strconv.FormatUint(id, 10)))
	if err != nil {
		return false
	}
	return exists
}

func (r *database) Save(url string, expires time.Time) (string, error) {
	// fmt.Println("save is called")
	conn := r.pool.Get()
	defer conn.Close()

	var id uint64

	for used := true; used; used = r.isUsed(id) {
		id = rand.Uint64()
	}
	// fmt.Println("id and url ", id, url)
	shortLink := storage.Item{id, url, expires.Format("2006-01-02 15:04:05.728046 +0300 EEST"), 0}

	// fmt.Println("HMSET performing...")
	_, err := conn.Do("HMSET", redis.Args{"Shortener:" + strconv.FormatUint(id, 10)}.AddFlat(shortLink)...)
	if err != nil {
		return "", err
	}
	// fmt.Println("EXPIRE PERFOMRING... ")
	_, err = conn.Do("EXPIREAT", "Shortener:"+strconv.FormatUint(id, 10), expires.Unix())
	if err != nil {
		return "", err
	}

	return base62.Encode(id), nil
}

func (r *database) Load(code string) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	decodedId, err := base62.Decode(code)
	if err != nil {
		return "", err
	}

	urlString, err := redis.String(conn.Do("HGET", "Shortener:"+strconv.FormatUint(decodedId, 10), "url"))
	if err != nil {
		return "", err
	} else if len(urlString) == 0 {
		return "", errors.New("No link")
	}

	_, err = conn.Do("HINCRBY", "Shortener:"+strconv.FormatUint(decodedId, 10), "visits", 1)

	return urlString, nil
}

func (r *database) isAvailable(id uint64) bool {
	conn := r.pool.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", "Shortener:"+strconv.FormatUint(id, 10)))
	if err != nil {
		return false
	}
	return !exists
}

func (r *database) LoadInfo(code string) (*storage.Item, error) {
	conn := r.pool.Get()
	defer conn.Close()

	decodedId, err := base62.Decode(code)
	if err != nil {
		return nil, err
	}

	values, err := redis.Values(conn.Do("HGETALL", "Shortener:"+strconv.FormatUint(decodedId, 10)))
	if err != nil {
		return nil, err
	} else if len(values) == 0 {
		return nil, errors.New("No link")
	}
	var shortLink storage.Item
	err = redis.ScanStruct(values, &shortLink)
	if err != nil {
		return nil, err
	}

	return &shortLink, nil
}

func (r *database) Close() error {
	return r.pool.Close()
}
