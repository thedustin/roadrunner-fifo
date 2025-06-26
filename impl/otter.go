package impl

import (
	"github.com/maypok86/otter/v2"
	"github.com/roadrunner-server/errors"
	"time"
)

type otterImpl struct {
	cache *otter.Cache[string, string]
}

var _ Fifo = (*otterImpl)(nil)

func NewOtterImpl(ttl time.Duration, sizeInBytes int) Fifo {
	return &otterImpl{
		cache: otter.Must(&otter.Options[string, string]{
			MaximumSize:       sizeInBytes,
			ExpiryCalculator:  otter.ExpiryAccessing[string, string](ttl),    // Reset timer on reads/writes
			RefreshCalculator: otter.RefreshWriting[string, string](ttl / 2), // Refresh after writes
			Weigher: func(key string, value string) uint32 {
				return uint32(len(key) + len(value))
			},
			// StatsRecorder: counter, // Attach stats collector
		}),
	}
}

func (o otterImpl) Set(key string, value string) error {
	if _, ok := o.cache.Set(key, value); !ok {
		return errors.E(errors.Op("otter_set"), "failed to set key", key)
	}
	return nil
}

func (o otterImpl) Get(key string) (string, bool) {
	val, ok := o.cache.GetIfPresent(key)

	return val, ok
}

func (o otterImpl) Invalidate(key string) bool {
	if _, ok := o.cache.Invalidate(key); !ok {
		return false
	}
	return true
}
