package keyvaluestore

import "sync"

const (
	dbname = "data.kvs"
)

type KeyValueStore struct {
	mux      sync.Mutex
	keyvalue map[string]int

	filename string
}

// New constructs and returns a cache service
func New() *KeyValueStore {
	return &KeyValueStore{
		keyvalue: make(map[string]int),
		filename: dbname,
	}
}
