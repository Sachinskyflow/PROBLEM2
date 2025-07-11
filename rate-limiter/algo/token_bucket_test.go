package algo_test

import (
	"testing"
	"time"

	"problem2/rate-limiter/algo"

	"github.com/stretchr/testify/assert"
)

func TestTokenBucket(t *testing.T) {
	bucket := algo.NewTokenBucket(10, 1, time.Second)

	assert.True(t, bucket.TryConsume(5))
	assert.True(t, bucket.TryConsume(5))
	assert.False(t, bucket.TryConsume(1))

	time.Sleep(10 * time.Second)

	assert.True(t, bucket.TryConsume(10))
}
