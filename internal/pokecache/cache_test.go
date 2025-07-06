package pokecache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	// Create a cache with 1 second interval
	cache := NewCache(1 * time.Second)

	// Test adding and getting an entry
	key := "test-key"
	value := []byte("test-value")

	cache.Add(key, value)

	// Get the value
	retrieved, found := cache.Get(key)
	if !found {
		t.Error("Expected to find the key in cache")
	}

	if string(retrieved) != string(value) {
		t.Errorf("Expected %s, got %s", string(value), string(retrieved))
	}

	// Test getting non-existent key
	_, found = cache.Get("non-existent")
	if found {
		t.Error("Expected not to find non-existent key")
	}
}

func TestCacheExpiration(t *testing.T) {
	// Create a cache with very short interval for testing
	cache := NewCache(100 * time.Millisecond)

	key := "expire-test"
	value := []byte("will-expire")

	cache.Add(key, value)

	// Should be found immediately
	_, found := cache.Get(key)
	if !found {
		t.Error("Expected to find the key immediately after adding")
	}

	// Wait for expiration
	time.Sleep(200 * time.Millisecond)

	// Should not be found after expiration
	_, found = cache.Get(key)
	if found {
		t.Error("Expected key to be expired and not found")
	}
}
