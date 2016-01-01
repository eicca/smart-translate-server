package cache

import (
	"bytes"
	"compress/gzip"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"

	"gopkg.in/redis.v3"
)

// DefaultClient is a top-level reference to a redis client.
var DefaultClient = NewRedisClient()

// Cachable is the object which could be cached by its CacheKey().
type Cachable interface {
	CacheKey() string
}

// Client abstracts the cache storage.
type Client interface {
	Set(obj Cachable, val interface{}) error
	Get(obj Cachable, val interface{}) error
	Flush() error
}

// RedisClient implement cache.Client based on redis.
type RedisClient struct {
	client *redis.Client
}

// NewRedisClient creates a client using `REDIS_HOST` env variable.
func NewRedisClient() Client {
	opts := redis.Options{}
	opts.Addr = os.Getenv("REDIS_HOST")
	return &RedisClient{client: redis.NewClient(&opts)}
}

// Set stores the value transformed into zipped json
// with the key hashed with sha1
func (rs *RedisClient) Set(obj Cachable, val interface{}) error {
	rawVal, err := json.Marshal(val)
	if err != nil {
		return err
	}

	var zipped bytes.Buffer
	w := gzip.NewWriter(&zipped)
	w.Write(rawVal)
	w.Close()

	return rs.client.Set(hash(obj.CacheKey()), zipped.String(), 0).Err()
}

// Get retrieves the value by sha1 hashed key.
// The value is stored into passed `val` parameter (generics...).
func (rs *RedisClient) Get(obj Cachable, val interface{}) error {
	rawVal, err := rs.client.Get(hash(obj.CacheKey())).Result()
	if err != nil {
		return err
	}

	zipped := bytes.NewBufferString(rawVal)
	reader, err := gzip.NewReader(zipped)
	if err != nil {
		return err
	}
	defer reader.Close()

	unzipped, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	return json.Unmarshal(unzipped, val)
}

// Flush wipes all data
func (rs *RedisClient) Flush() error {
	return rs.client.FlushAll().Err()
}

func hash(key string) string {
	hasher := sha1.New()
	io.WriteString(hasher, key)
	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
}
