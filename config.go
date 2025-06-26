package fifo

import "time"

const (
	defaultMaxCacheSize    = (64 - 1) * 1024 * 1024 // Default max cache size, 63 MiB (keep some space for overhead)
	defaultExpiring        = 5 * time.Minute        // Default TTL for cache entries
	defaultRefreshInterval = 3 * time.Second        // Default refresh interval for cache entries
)

type Config struct {
	// MaxCacheSize is the maximum size of the cache in megabytes (will be converted to bytes).
	MaxCacheSize int `mapstructure:"max_cache_size"`

	// Expiring is the duration after which cache entries are considered expired (like a ttl, https://maypok86.github.io/otter/user-guide/v2/features/eviction/).
	Expiring time.Duration `mapstructure:"expiring"`

	// RefreshInterval determines how often new entries are loaded into the cache (https://maypok86.github.io/otter/user-guide/v2/features/refresh/).
	RefreshInterval time.Duration `mapstructure:"refresh_interval"`
}

func (c Config) InitDefaults() error {
	if c.MaxCacheSize == 0 {
		c.MaxCacheSize = defaultMaxCacheSize
	} else {
		c.MaxCacheSize *= 1024 * 1024 // Convert to bytes
	}

	if c.Expiring == 0 {
		c.Expiring = defaultExpiring
	}

	if c.RefreshInterval == 0 {
		c.RefreshInterval = defaultRefreshInterval
	}

	return nil
}
