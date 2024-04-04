// Copyright (c) 2024 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the Apache-2.0 license, a copy of
// which may be found at https://github.com/silverton-io/kota/blob/main/LICENSE

package config

type Input struct {
	Okta    `json:"okta"`
	Splunk  `json:"splunk"`
	Kafka   `json:"kafka,omitempty"`
	Kinesis `json:"kinesis,omitempty"`
	Pubsub  `json:"pubsub,omitempty"`
	Sqs     `json:"sqs,omitempty"`
}

type Okta struct {
	Hook `json:"hook"`
	Api  `json:"api"`
}

type Sqs struct {
	Enabled bool   `json:"enabled"`
	Queue   string `json:"queue,omitempty"`
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
	Topic   string   `json:"topic,omitempty"`
	Brokers []string `json:"brokers,omitempty"`
}

type Kinesis struct {
	Enabled bool   `json:"enabled"`
	Stream  string `json:"stream,omitempty"`
}

type Pubsub struct {
	Enabled bool   `json:"enabled"`
	Topic   string `json:"topic,omitempty"`
}
