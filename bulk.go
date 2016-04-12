package ratelimit

import "sync"

type Bulk struct {
	buckets map[string]*Buckets
	mu      sync.RWMutex
	pool    sync.Pool
}

func NewBulk(rate, rateSecond int) Bulk {
	return Bulk{
		buckets: make(map[string]*Buckets),
		pool: sync.Pool{
			New: func() interface{} {
				bkt := NewBuckets(rate, rateSecond)
				return &bkt
			},
		},
	}
}

func (b *Bulk) Take(key string, n int) (max, remain, reset int, taked bool) {
	b.mu.Lock()
	bkt, has := b.buckets[key]
	if !has {
		bkt = b.pool.Get().(*Buckets)
		bkt.Reset()
		b.buckets[key] = bkt
	}
	max, remain, reset, taked = bkt.Take(n)
	b.mu.Unlock()

	return
}

func (b *Bulk) Clean() {
	b.mu.Lock()
	for key, bkt := range b.buckets {
		if bkt.IsFull() {
			delete(b.buckets, key)
			b.pool.Put(bkt)
		}
	}
	b.mu.Unlock()
}
