// Copyright (c) 2024 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the Apache-2.0 license, a copy of
// which may be found at https://github.com/silverton-io/kota/blob/main/LICENSE

package consumer

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/silverton.io/kota/pkg/buffer"
	"github.com/silverton.io/kota/pkg/config"
	"github.com/silverton.io/kota/pkg/envelope"
)

type ApiConsumer struct {
	buffer   *buffer.Buffer
	ctx      context.Context
	shutdown context.CancelFunc
}

func (c *ApiConsumer) Initialize(config *config.Input, buffer *buffer.Buffer) error {
	c.buffer = buffer
	c.ctx, c.shutdown = context.WithCancel(context.Background())
	go func() {
		for {
			// TODO -> Run this wide open but back off according to 429's
			// Log when system has been rate limited
			select {
			case <-c.ctx.Done():
				log.Debug().Msg("shut down api consumer")
				return
			default:
				c.Consume()
			}
			time.Sleep(time.Duration(500) * time.Millisecond) // TODO -> remove this and run wide open but back off.
		}
	}()
	return nil
}

func (c *ApiConsumer) Consume() {
	log.Debug().Msg("polling the system log api")
	// TODO -> actually poll the thing, wrap results in the envelope, and send them along
	fake_envelope := envelope.BuildFakeEnvelope()
	c.buffer.Append([]envelope.KotaEnvelope{fake_envelope})
}

func (c *ApiConsumer) Shutdown() error {
	c.shutdown()
	return nil
}

func NewApiConsumer(config *config.Input, buffer *buffer.Buffer) *ApiConsumer {
	consumer := ApiConsumer{}
	consumer.Initialize(config, buffer)
	return &consumer
}
