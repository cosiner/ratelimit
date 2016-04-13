package ratelimit

import (
	"fmt"
	"testing"
	"time"
)

func TestToken(t *testing.T) {
	b := NewToken(2, 5)
	for i := 0; i < 120; i++ {
		b.Take(1)
		//time.Sleep(time.Second)
	}
}

func TestLeaky(t *testing.T) {
	b := NewLeaky(5, 10)
	for i := 0; i < 120; i++ {
		fmt.Println(b.Take(-1))
		time.Sleep(time.Second)
	}
}

func BenchmarkToken(b *testing.B) {
	bkt := NewToken(1200, 3600)
	for i := 0; i < b.N; i++ {
		bkt.Take(1)
	}
}

func BenchmarkBulk(b *testing.B) {
	bkt := NewBulk(1200, 3600, NewToken)
	fmt.Println(b.N)
	for i := 0; i < b.N; i++ {
		bkt.Take("Test", 1)
	}
}

func TestBulk(t *testing.T) {
	b := NewBulk(2, 5, NewToken)
	for i := 0; i < 120; i++ {
		b.Take("test", 1)
	}
}
