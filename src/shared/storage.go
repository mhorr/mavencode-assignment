package shared

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

const personindex = "person:timeindex"

// RedisPersonStore is the
// interface between the golang person struct and Redis
type RedisPersonStore interface {
	Store(p Person) error
	Query(param string) ([]Person, error)
}

type redisInfo struct {
	conn redis.Conn
}

// GetRedisPersonStore initializes a connection
// to the redis store and returns an object for dealing with it.
func GetRedisPersonStore() (RedisPersonStore, error) {
	c, err := redis.Dial("tcp", GetConfig().Redis)
	if err != nil {
		return nil, err
	}
	return redisInfo{
		conn: c,
	}, nil
}

// Store writes a person to the redis store.
func (s redisInfo) Store(p Person) error {
	pb, err := json.Marshal(p)
	if err != nil {
		return err
	}
	fullname := p.Firstname + "-" + p.Lastname
	key := "person:" + fullname
	r, err := s.conn.Do("SET", key, pb, "EX", GetConfig().RedisTTL)
	if err != nil {
		log.Printf("%s: %s", err, "")
		return err
	}
	log.Printf("Response from set %s: %#v\n", fullname, r)

	r, err = s.conn.Do("ZADD", personindex, p.Timestamp.Unix(), fullname)
	if err != nil {
		log.Printf("%s: %s", err, "")
		return err
	}
	log.Printf("Response from ZADD %s: %#v\n", fullname, r)

	return nil
}

// Query retrieves person structs from the redis store
func (s redisInfo) Query(param string) ([]Person, error) {
	result := []Person{}
	r, err := redis.Strings(s.conn.Do("ZRANGEBYSCORE", personindex, 0, time.Now().Unix()))
	if err != nil {
		return result, err
	}
	for _, fullname := range r {
		r, err := redis.Bytes(s.conn.Do("GET", "person:"+fullname))
		if err != nil {
			log.Printf("Error on GET for %s: %s\n", "person:"+fullname, err)
			continue
			//return result, err
		}
		var p Person
		err = json.Unmarshal(r, &p)
		if err != nil {
			log.Printf("Error Unmarshalling person object: %s\n", err)
			continue
			//return result, err
		}
		result = append(result, p)
	}
	log.Printf("result: %s\n", r)
	return result, nil
}
