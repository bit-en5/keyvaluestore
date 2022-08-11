package keyvaluestore

import (
	"fmt"
	"testing"
)

func TestCache(t *testing.T) {
	want1 := 150
	want2 := 81

	cache := New()
	cache.Set("value1", want1)
	cache.Set("value2", want2)

	has1, _ := cache.Get("value1")
	has2, _ := cache.Get("value2")
	has3, found3 := cache.Get("value3")

	if has1 != want1 {
		t.Errorf("error: expected %d and received %d", want1, has1)
	}

	if has2 != want2 {
		t.Errorf("error: expected %d and received %d", want2, has2)
	}

	if has3 != 0 || found3 {
		t.Error("error: expected 0 and not found")
	}

	cache.Delete("value1")
	v1, found1 := cache.Get("value1")

	if v1 != 0 || found1 {
		t.Error("error: expected 0 and not found after deletion")
	}
}

func TestPersistance(t *testing.T) {
	cache := New()
	path := cache.getFileName()
	fmt.Println(path)
}
