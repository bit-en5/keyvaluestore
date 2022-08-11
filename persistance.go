package keyvaluestore

import (
	"encoding/json"
	"os"
)

// Save persists the key-value store on disk
func (s *KeyValueStore) Save() error {
	data, err := json.Marshal(s.keyvalue)
	if err != nil {
		return err
	}

	return os.WriteFile(s.getFileName(), data, 0777)
}

// Read read the key-value store from disk
func (s *KeyValueStore) Read() error {
	data, err := os.ReadFile(s.getFileName())
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &s.keyvalue)
}

func (s *KeyValueStore) getFileName() string {
	return os.TempDir() + s.filename
}
