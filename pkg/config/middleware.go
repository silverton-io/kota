// Copyright (c) 2024 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the Apache-2.0 license, a copy of
// which may be found at https://github.com/silverton-io/kota/blob/main/LICENSE

package config

type Middleware struct {
	Timeout       `json:"timeout"`
	RateLimiter   `json:"rateLimiter"`
	RequestLogger `json:"requestLogger"`
	Auth          `json:"auth"`
}

type Timeout struct {
	Enabled bool `json:"enabled"`
	Ms      int  `json:"ms"`
}

type RateLimiter struct {
	Enabled bool   `json:"enabled"`
	Period  string `json:"period"`
	Limit   int64  `json:"limit"`
}

type RequestLogger struct {
	Enabled bool `json:"enabled"`
}

type Auth struct {
	Enabled bool     `json:"enabled"`
	Tokens  []string `json:"-"`
}
