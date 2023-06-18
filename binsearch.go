package binsearch

import (
	"encoding/binary"
	"github.com/fakefloordiv/binsearch/internal"
)

type entry[T any] struct {
	hash  uint64
	key   string
	value T
}

type Container[T any] struct {
	entries []entry[T]
}

func NewContainer[T any]() *Container[T] {
	return new(Container[T])
}

func (c *Container[T]) SetMap(values map[string]T) {
	for key, value := range values {
		c.Set(key, value)
	}
}

func (c *Container[T]) Set(key string, value T) {
	pair := entry[T]{
		hash:  strhash([]byte(key)),
		key:   key,
		value: value,
	}

	for i, val := range c.entries {
		if pair.hash <= val.hash {
			c.entries = arrInsert(i, pair, c.entries)
			return
		}
	}

	// new hash is the biggest one
	c.entries = append(c.entries, pair)
}

func (c *Container[T]) Get(key string) (val T, found bool) {
	var (
		p       int
		entries = c.entries
		hash    = strhash(internal.S2B(key))
	)

	for a := len(entries); a > 0; a /= 2 {
		for p+a < len(entries) && entries[p+a].hash <= hash {
			p += a
		}
	}

	e := entries[p]

	return e.value, e.key == key
}

func strhash(str []byte) (hash uint64) {
	for i := 8; i < len(str); i += 8 {
		hash += binary.BigEndian.Uint64(str[i-8 : i])
		hash += hash << 10
		hash ^= hash >> 6
	}

	for i := 0; i < len(str)%8; i++ {
		hash += uint64(str[i])
		hash += hash << 10
		hash ^= hash >> 6
	}

	hash += hash << 3
	hash ^= hash >> 11
	hash += hash << 15

	return hash
}

// arrInsert inserts a value in the provided position. It is made by
// shifting the second part after the index by 1. In case current slice
// cannot hold all the new values, a new-one with +1 size will be allocated
func arrInsert[T any](position int, value T, arr []T) []T {
	if len(arr) == 0 {
		return append(arr, value)
	}

	if len(arr) == cap(arr) {
		newArr := make([]T, len(arr), len(arr)+1)
		copy(newArr, arr)
		arr = newArr
	}

	copy(arr[position+1:cap(arr)], arr[position:])
	arr[position] = value
	return arr[:len(arr)+1]
}
