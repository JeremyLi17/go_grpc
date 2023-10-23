package limiter

import "golang.org/x/time/rate"

var (
	Limiter *rate.Limiter
)

func InitLimiter() {
	Limiter = rate.NewLimiter(rate.Limit(10), 1)
}
