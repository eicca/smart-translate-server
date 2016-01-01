package cache

import (
	"testing"
)

type testStruct struct {
	Key string
}

func (s testStruct) CacheKey() string {
	return "key:" + s.Key
}

func TestRedisGetSet(t *testing.T) {
	c := NewRedisClient()
	c.Flush()

	s := testStruct{"foo"}
	actual := "bar"
	err := c.Set(s, actual)
	if err != nil {
		t.Fatal(err)
	}

	var expected string
	err = c.Get(s, &expected)
	if err != nil {
		t.Fatal(err)
	}
	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

func TestNonExistingGet(t *testing.T) {
	c := NewRedisClient()
	c.Flush()

	s := testStruct{"foo"}
	if c.Get(s, nil) == nil {
		t.Error("cache.Get('missing key') should return error")
	}
}
