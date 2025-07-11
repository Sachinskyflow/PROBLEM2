package algo

import (
	"time"
)

type TokenBucket struct {
	capacity     int
	tokens       int
	rate         int
	lastWithdraw time.Time
	unit         time.Duration
}

func NewTokenBucket(capacity, rate int, unit time.Duration) *TokenBucket {
	return &TokenBucket{
		capacity:     capacity,
		tokens:       capacity,
		rate:         rate,
		lastWithdraw: time.Now(),
		unit:         unit,
	}
}

func (to *TokenBucket) Refill() {
	now := time.Now()
	elapsed := now.Sub(to.lastWithdraw)
	tokensToAdd := int(elapsed.Seconds() / to.unit.Seconds() * float64(to.rate))
	if tokensToAdd > 0 {
		to.tokens += tokensToAdd
		if to.tokens > to.capacity {
			to.tokens = to.capacity
		}
		to.lastWithdraw = now
	}
}

func (to *TokenBucket) TryConsume(numRequest int) bool {
	to.Refill()
	if to.tokens >= numRequest {
		to.tokens -= numRequest
		return true
	}
	return false
}
