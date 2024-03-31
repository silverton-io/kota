// Copyright (c) 2024 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the Apache-2.0 license, a copy of
// which may be found at https://github.com/silverton-io/kota/blob/main/LICENSE

package buffer

import (
	"github.com/rs/zerolog/log"
	"github.com/silverton.io/kota/pkg/config"
	"github.com/silverton.io/kota/pkg/envelope"
	"github.com/silverton.io/kota/pkg/util"
)

type Buffer struct {
	config    *config.Config
	inputChan chan envelope.KotaEnvelope
	envelopes []envelope.KotaEnvelope
	shutdown  chan int
}

func (b *Buffer) Initialize(config *config.Config) error {
	b.inputChan = make(chan envelope.KotaEnvelope, 2000)
	b.shutdown = make(chan int, 1)
	b.config = config
	go func(envelope <-chan envelope.KotaEnvelope, shutdown chan int) {
		for {
			select {
			case envelope := <-envelope:
				util.Pprint(envelope) // TODO -> persist to disk, flush, etc
			case <-shutdown:
				log.Debug().Msg("shutting down buffer")
				// TODO -> make me do something that is safe on shutdown
				return
			}
		}
	}(b.inputChan, b.shutdown)
	return nil
}

func (b *Buffer) Append(envelopes []envelope.KotaEnvelope) error {
	for _, envelope := range envelopes {
		log.Debug().Interface("envelope", envelope).Msg("enqueueing envelope")
		b.inputChan <- envelope
	}
	return nil
}

func (b *Buffer) Purge() error {
	log.Debug().Msg("purging buffer")
	b.envelopes = []envelope.KotaEnvelope{}
	return nil
}

func (b *Buffer) Shutdown() error {
	b.shutdown <- 1
	return nil
}

func NewBuffer(config *config.Config) Buffer {
	buffer := Buffer{}
	buffer.Initialize(config)
	return buffer
}
