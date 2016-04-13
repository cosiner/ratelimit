package ratelimit

import "sync"

type Bulk struct {
	buckets map[string]Bucket
	mu      sync.RWMutex
	pool    sync.Pool
}

func NewBulk(rate, rateSecond int, newBucket func(int, int) Bucket) Bulk {
	return Bulk{
		buckets: make(map[string]Bucket),
		pool: sync.Pool{
			New: func() interface{} {
				return newBucket(rate, rateSecond)
			},
		},
	}
}

func (b *Bulk) Take(key string, n int) (max, remain, reset int, taken bool) {
	b.mu.Lock()
	bkt, has := b.buckets[key]
	if !has {
		bkt = b.pool.Get().(Bucket)
		bkt.Reset()
		b.buckets[key] = bkt
	}
	max, remain, reset, taken = bkt.Take(n)
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
