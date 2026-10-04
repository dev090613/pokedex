package pokecache

import (
	"testing"
	"time"
	"fmt"
)

func TestRoundTrip(t *testing.T) {
	// 선언 c, key, want
	c := NewCache(5 * time.Second)
	key := "test_key"
	want := []byte("test_value")

	// 실행 got := c.Get(key)
	c.Add(key, want)
	got, ok := c.Get(key)
	if !ok {
		t.Fatalf("expected to find key %q", key)
	}

	// 검증: want == got
	if string(want) != string(got) {
		t.Errorf("Error: got %q, want %q", got, want)
	}
	
}

func TestNotRoundTrip(t *testing.T) {
	// 선언 c, key, want
	c := NewCache(5 * time.Second)
	key := "test_key"
	notKey := "test_not_key"
	want := []byte("test_value")

	// 실행
	c.Add(key, want)
	got, ok := c.Get(notKey)
	if ok {
		t.Fatalf("Fatal: got %q, want %q", got, want)
	}
	
}

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}

