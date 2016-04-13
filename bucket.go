package ratelimit

type Bucket interface {
	Take(n int) (rate, remain, reset int, taken bool)
	IsFull() bool
	IsEmpty() bool
	Reset()
}
