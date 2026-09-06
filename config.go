package quic

import (
	"time"
)

func (c *Config) Clone() *Config { _ = "STUB: not implemented"; return nil }

func (c *Config) handshakeTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c *Config) maxRetryTokenAge() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func validateConfig(config *Config) error { _ = "STUB: not implemented"; return nil }

func populateConfig(config *Config) *Config { _ = "STUB: not implemented"; return nil }
