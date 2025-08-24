package pokecache

import (
	"testing"
	"time"
)

func TestAddAndGet(t *testing.T) {
	cache := NewCache(1 * time.Second)
	key := "test-key"
	value := []byte("test-value")

	cache.Add(key, value)

	got, ok := cache.Get(key)
	if !ok {
		t.Errorf("expected key to be present")
	}
	if string(got) != string(value) {
		t.Errorf("expected value %s, got %s", value, got)
	}
}

func TestReapLoopRemovesOldEntries(t *testing.T) {
	// Using a short interval for fast tests
	interval := 10 * time.Millisecond
	cache := NewCache(interval)

	key := "old-key"
	value := []byte("should get removed")

	cache.Add(key, value)

	// Immediately after adding, the key should be present
	if _, ok := cache.Get(key); !ok {
		t.Fatalf("expected key to be present right after Add")
	}

	// Wait for slightly longer than the interval to allow reaping
	time.Sleep(2 * interval)

	// Now, the entry should have been reaped/removed
	if _, ok := cache.Get(key); ok {
		t.Errorf("expected key %s to be removed after reaping", key)
	}
}
