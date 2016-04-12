package ratelimit

import "time"

type Buckets struct {
	rate       int
	rateSecond int

	start   int
	avaible int
}

func nowSecond() int {
	return int(time.Now().Unix())
}

func NewBuckets(rate, rateSecond int) Buckets {
	return Buckets{
		rate:       rate,
		rateSecond: rateSecond,

		start:   nowSecond(),
		avaible: rate,
	}
}

func (b *Buckets) Take(n int) (rate, remain, reset int, taked bool) {
	dur := b.adjust()
	rate = b.rate
	if taked = b.avaible >= n; taked {
		b.avaible -= n
	}
	remain = b.avaible
	reset = b.rateSecond - dur
	return
}

func (b *Buckets) IsFull() bool {
	b.adjust()
	fulled := b.avaible >= b.rate
	return fulled
}

func (b *Buckets) adjust() int {
	now := nowSecond()

	cycle := (now - b.start) / b.rateSecond
	b.start += cycle * b.rateSecond
	if cycle > 0 {
		b.avaible = b.rate
	}
	return now - b.start
}

func (b *Buckets) Reset() {
	b.start = nowSecond()
	b.avaible = b.rate
}
