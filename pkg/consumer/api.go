// Copyright (c) 2024 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the Apache-2.0 license, a copy of
// which may be found at https://github.com/silverton-io/kota/blob/main/LICENSE

package consumer

import (
	"time"

	"github.com/rs/zerolog/log"
	"github.com/silverton.io/kota/pkg/buffer"
	"github.com/silverton.io/kota/pkg/config"
	"github.com/silverton.io/kota/pkg/envelope"
)

type ApiConsumer struct {
	buffer   *buffer.Buffer
	shutdown chan int
}

func (c *ApiConsumer) Initialize(config *config.Input, buffer *buffer.Buffer) error {
	c.buffer = buffer
	c.shutdown = make(chan int, 1)
	ticker := time.NewTicker(time.Duration(500 * float64(time.Millisecond)))
	go func(shutdown <-chan int) {
		for {
			select {
			case <-ticker.C:
				c.Consume()
			case <-shutdown:
				ticker.Stop()
				log.Debug().Msg("shut down api consumer")
				return
			}
		}
	}(c.shutdown)
	return nil
}

func (c *ApiConsumer) Consume() {
	log.Debug().Msg("polling the system log api")
	// TODO -> actually poll the thing, wrap results in the envelope, and send them along
	fake_envelope := envelope.BuildFakeEnvelope()
	c.buffer.Append([]envelope.KotaEnvelope{fake_envelope})
}

func (c *ApiConsumer) Shutdown() error {
	c.shutdown <- 1
	return nil
}

func NewApiConsumer(config *config.Input, buffer *buffer.Buffer) *ApiConsumer {
	consumer := ApiConsumer{}
	consumer.Initialize(config, buffer)
	return &consumer
}
