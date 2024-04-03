// Copyright (c) 2024 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the Apache-2.0 license, a copy of
// which may be found at https://github.com/silverton-io/kota/blob/main/LICENSE

package consumer

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/silverton.io/kota/pkg/config"
	"github.com/twmb/franz-go/pkg/kgo"
)

type KafkaConsumer struct {
	client   *kgo.Client
	stream   string
	group    string
	shutdown chan int
}

func (c *KafkaConsumer) Initialize(config config.Consumer) error {
	ctx := context.Background()
	log.Debug().Msg("initializing kafka client")
	client, err := kgo.NewClient(
		kgo.SeedBrokers(config.Brokers...),
	)
	c.client = client
	if err != nil {
		log.Debug().Stack().Err(err).Msg("could not create kafka sink client")
		return err
	}
	log.Debug().Msg("pinging kafka brokers")
	err = client.Ping(ctx)
	if err != nil {
		log.Debug().Stack().Err(err).Msg("could not ping kafka sink brokers")
		return err
	}
	c.shutdown = make(chan int, 1)
	return nil
}

func (c *KafkaConsumer) Consume() error {
	return nil
}

func (c *KafkaConsumer) Shutdown() error {
	return nil
}
