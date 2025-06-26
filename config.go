package fifo

import "time"

const (
	defaultMaxCacheSize = (64 - 1) * 1024 * 1024 // Default max cache size, 63 MiB (keep some space for overhead)
	defaultTtl          = 5 * time.Minute        // Default TTL for cache entries
)

type Config struct {
	// MaxCacheSize is the maximum size of the cache in megabytes (will be converted to bytes).
	MaxCacheSize int `mapstructure:"max_cache_size"`
	// Ttl is the time-to-live for cache entries, default is 5 minutes.
	Ttl time.Duration `mapstructure:"ttl"`
}

func (c Config) InitDefaults() error {
	if c.MaxCacheSize == 0 {
		c.MaxCacheSize = defaultMaxCacheSize
	} else {
		c.MaxCacheSize *= 1024 * 1024 // Convert to bytes
	}

	if c.Ttl == 0 {
		c.Ttl = defaultTtl
	}

	return nil
}
