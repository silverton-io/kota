// Copyright (c) 2024 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the Apache-2.0 license, a copy of
// which may be found at https://github.com/silverton-io/kota/blob/main/LICENSE

package config

type Input struct {
	// General
	Group string `json:"group"`
	// Kinesis, Pub/Sub
	Stream string `json:"stream,omitempty"`
	// Eventbridge
	Bus string `json:"bus,omitempty"`
	// Kafka
	Topic   string   `json:"topic,omitempty"`
	Brokers []string `json:"kakfaBrokers,omitempty"`
	// API
	Endpoint string `json:"endpoint,omitempty"`
}
