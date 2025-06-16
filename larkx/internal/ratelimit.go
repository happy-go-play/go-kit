package internal

import (
	"sync"
	"time"
)

type RateLimiter struct {
	secondLimit  int           // 每秒限制(5)
	minuteLimit  int           // 每分钟限制(100)
	secondWindow time.Duration // 秒级窗口(1s)
	minuteWindow time.Duration // 分钟级窗口(1m)

	secondRequests []time.Time // 秒级请求记录
	minuteRequests []time.Time // 分钟级请求记录
	mutex          sync.Mutex
}

func NewRateLimiter(secondLimit, minuteLimit int) *RateLimiter {
	return &RateLimiter{
		secondLimit:  secondLimit,
		minuteLimit:  minuteLimit,
		secondWindow: time.Second,
		minuteWindow: time.Minute,
	}
}

func (r *RateLimiter) Allow() bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	now := time.Now()

	// 清理秒级过期的请求
	for len(r.secondRequests) > 0 && now.Sub(r.secondRequests[0]) > r.secondWindow {
		r.secondRequests = r.secondRequests[1:]
	}

	// 清理分钟级过期的请求
	for len(r.minuteRequests) > 0 && now.Sub(r.minuteRequests[0]) > r.minuteWindow {
		r.minuteRequests = r.minuteRequests[1:]
	}

	// 检查秒级限制
	if len(r.secondRequests) >= r.secondLimit {
		return false
	}

	// 检查分钟级限制
	if len(r.minuteRequests) >= r.minuteLimit {
		return false
	}

	// 记录本次请求
	r.secondRequests = append(r.secondRequests, now)
	r.minuteRequests = append(r.minuteRequests, now)
	return true
}
