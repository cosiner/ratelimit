package ratelimit

import "time"

type Token struct {
	rate       int
	rateSecond int

	start     int
	available int
}

func nowSecond() int {
	return int(time.Now().Unix())
}

func NewToken(rate, rateSecond int) Bucket {
	return &Token{
		rate:       rate,
		rateSecond: rateSecond,

		start:     nowSecond(),
		available: rate,
	}
}

func (b *Token) Take(n int) (rate, remain, reset int, taken bool) {
	dur := b.adjust()
	if taken = b.available >= n; taken {
		b.available -= n
	}
	return b.rate, b.available, b.rateSecond - dur, taken
}

func (b *Token) IsFull() bool {
	b.adjust()
	return b.available >= b.rate
}

func (b *Token) IsEmpty() bool {
	b.adjust()
	return b.available <= 0
}

func (b *Token) adjust() int {
	now := nowSecond()

	cycle := (now - b.start) / b.rateSecond
	b.start += cycle * b.rateSecond
	if cycle > 0 {
		b.available = b.rate
	}
	return now - b.start
}

func (b *Token) Reset() {
	b.start = nowSecond()
	b.available = b.rate
}
