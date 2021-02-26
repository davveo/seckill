package rate_limit

import (
	"golang.org/x/time/rate"
	"sync"
)

type IPRateLimiter struct {
	Ips map[string]*rate.Limiter
	mu *sync.RWMutex
	r rate.Limit
	b int
}

func NewIpRateLimiter(r rate.Limit, b int) *IPRateLimiter  {
	return &IPRateLimiter{
		Ips: make(map[string]*rate.Limiter),
		mu: &sync.RWMutex{},
		r: r,
		b: b,
	}
}

func (i *IPRateLimiter) AddIp(ip string) *rate.Limiter  {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(i.r, i.b)
	i.Ips[ip] = limiter
	return limiter
}

func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	limiter, exists := i.Ips[ip]

	if !exists {
		i.mu.Unlock()
		return i.AddIp(ip)
	}

	i.mu.Unlock()

	return limiter
}


