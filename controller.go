package keyvaluestore

// Get returns the value related to the given key and a flag to indicate if the key has been found
func (s *KeyValueStore) Get(key string) (v int, found bool) {
	s.mux.Lock()
	defer s.mux.Unlock()
	v, found = s.keyvalue[key]
	return
}

// Set saves the value related to the given key
func (s *KeyValueStore) Set(key string, value int) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.keyvalue[key] = value
}

// Delete removes the value related to the given key
func (s *KeyValueStore) Delete(key string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	delete(s.keyvalue, key)
}
