package keyvaluestore

import "sync"

type KeyValueStore struct {
	mux      sync.Mutex
	keyvalue map[string]int

	filename string
}

// New constructs and returns a cache service
func New(dbName string) *KeyValueStore {
	return &KeyValueStore{
		keyvalue: make(map[string]int),
		filename: dbName,
	}
}
