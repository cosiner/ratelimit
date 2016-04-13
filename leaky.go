package ratelimit

type Leaky struct {
	rate       int
	rateSecond int

	start     int
	available int
}

func NewLeaky(rate, rateSecond int) Bucket {
	return &Leaky{
		rate:       rate,
		rateSecond: rateSecond,

		start:     nowSecond(),
		available: 0,
	}
}

func (b *Leaky) Take(n int) (rate, remain, reset int, putted bool) {
	dur := b.adjust()
	curr := b.available - n
	if putted = curr <= b.rate; putted {
		b.available = curr
	}
	return b.rate, b.rate - b.available, b.rateSecond - dur, putted
}

func (b *Leaky) IsFull() bool {
	b.adjust()
	return b.available >= b.rate
}

func (b *Leaky) IsEmpty() bool {
	b.adjust()
	return b.available <= 0
}

func (b *Leaky) adjust() int {
	now := nowSecond()

	cycle := (now - b.start) / b.rateSecond
	b.start += cycle * b.rateSecond
	if cycle > 0 {
		b.available -= b.rate
		if b.available < 0 {
			b.available = 0
		}
	}
	return now - b.start
}

func (b *Leaky) Reset() {
	b.start = nowSecond()
	b.available = 0
}
