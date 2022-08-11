package keyvaluestore

import (
	"encoding/json"
	"os"
)

// Save persists the key-value store on disk
func (s *KeyValueStore) Save() error {
	if s.filename == "" {
		return nil
	}

	data, err := s.serialize()
	if err != nil {
		return err
	}

	return os.WriteFile(s.getFileName(), data, 0777)
}

func (s *KeyValueStore) serialize() ([]byte, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	return json.Marshal(s.keyvalue)
}

// Read read the key-value store from disk
func (s *KeyValueStore) Read() error {
	if s.filename == "" {
		return nil
	}

	data, err := os.ReadFile(s.getFileName())
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &s.keyvalue)
}

func (s *KeyValueStore) getFileName() string {
	return os.TempDir() + s.filename
}
