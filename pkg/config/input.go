// Copyright (c) 2024 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the Apache-2.0 license, a copy of
// which may be found at https://github.com/silverton-io/kota/blob/main/LICENSE

package config

type Input struct {
	Okta    `json:"okta"`
	Splunk  `json:"splunk"`
	Kafka   `json:"kafka"`
	Kinesis `json:"kinesis"`
	Pubsub  `json:"pubsub"`
}

type Okta struct {
	Eventbridge `json:"eventbridge"`
	Hook        `json:"hook"`
	Api         `json:"api"`
}

type Eventbridge struct {
	Enabled bool   `json:"enabled"`
	Bus     string `json:"bus"`
}

type Hook struct {
	Enabled bool `json:"enabled"`
}

type Api struct {
	Enabled bool `json:"enabled"`
}

type Splunk struct {
	Enabled bool `json:"enabled"`
}

type Kafka struct {
	Enabled bool     `json:"enabled"`
	Topic   string   `json:"topic"`
	Brokers []string `json:"brokers"`
}

type Kinesis struct {
	Enabled bool   `json:"enabled"`
	Stream  string `json:"stream"`
}

type Pubsub struct {
	Enabled bool   `json:"enabled"`
	Topic   string `json:"topic"`
}
