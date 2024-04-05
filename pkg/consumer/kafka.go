// Copyright (c) 2024 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the Apache-2.0 license, a copy of
// which may be found at https://github.com/silverton-io/kota/blob/main/LICENSE

package consumer

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/silverton.io/kota/pkg/buffer"
	"github.com/silverton.io/kota/pkg/config"
	"github.com/silverton.io/kota/pkg/constants"
	"github.com/silverton.io/kota/pkg/envelope"
	"github.com/silverton.io/kota/pkg/util"
	"github.com/twmb/franz-go/pkg/kgo"
)

type KafkaConsumer struct {
	client   *kgo.Client
	topic    string
	group    string
	buffer   *buffer.Buffer
	ctx      context.Context
	shutdown context.CancelFunc
}

func (c *KafkaConsumer) Initialize(config *config.Input, buffer *buffer.Buffer) error {
	ctx, shutdown := context.WithCancel(context.Background())
	log.Debug().Msg("initializing kafka consumer")
	c.topic = config.Kafka.Topic
	c.group = constants.KOTA
	c.ctx = ctx
	c.shutdown = shutdown
	c.buffer = buffer
	client, err := kgo.NewClient(
		kgo.SeedBrokers(config.Kafka.Brokers...),
		kgo.ConsumerGroup(c.group),
		kgo.ConsumeTopics(c.topic),
	)
	c.client = client
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("could not create kafka consumer client")
		return err
	}
	return nil
}

func (c *KafkaConsumer) Consume() {
	log.Debug().Msg("starting kafka consumer from topic: " + c.topic)
	go func(buffer *buffer.Buffer) {
		for {
			var envelopes []envelope.KotaEnvelope
			log.Trace().Msg("polling kafka for new records from topic: " + c.topic)
			fetches := c.client.PollRecords(c.ctx, 1000)
			iter := fetches.RecordIter()
			for !iter.Done() {
				record := iter.Next()
				util.Pprint(record) // TODO -> actually do something useful with this
				// TODO -> Wrap the record in an envelope and pass to buffer
				// This is just a stub for now.
				envelope := envelope.BuildFakeEnvelope()
				envelopes = append(envelopes, envelope)
				buffer.Append(envelopes)
			}
		}
	}(c.buffer)
}

func (c *KafkaConsumer) Shutdown() error {
	log.Debug().Msg("shutting down kafka consumer")
	c.shutdown()
	c.client.Close()
	return nil
}

func NewKafkaConsumer(config *config.Input, buffer *buffer.Buffer) *KafkaConsumer {
	consumer := KafkaConsumer{}
	consumer.Initialize(config, buffer)
	consumer.Consume()
	return &consumer
}
