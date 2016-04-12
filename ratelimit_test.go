package ratelimit

import (
	"fmt"
	"testing"
	"time"
)

func TestBuckets(t *testing.T) {
	b := NewBuckets(2, 5)
	for i := 0; i < 120; i++ {
		fmt.Println(b.Take(1))
		time.Sleep(time.Second)
	}
}

func TestBulk(t *testing.T) {
	b := NewBulk(2, 5)
	for i := 0; i < 120; i++ {
		fmt.Println(b.Take("test", 1))
		time.Sleep(time.Second)
	}
}
