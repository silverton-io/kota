// Copyright (c) 2024 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the Apache-2.0 license, a copy of
// which may be found at https://github.com/silverton-io/kota/blob/main/LICENSE

package buffer

import (
	"time"

	"github.com/rs/zerolog/log"
	"github.com/silverton.io/kota/pkg/config"
	"github.com/silverton.io/kota/pkg/envelope"
	"github.com/silverton.io/kota/pkg/util"
)

type Buffer struct {
	config        *config.Config
	bufferRecords int
	bufferStart   time.Time
	inputChan     chan envelope.KotaEnvelope
	envelopes     []envelope.KotaEnvelope
	shutdown      chan int
}

func (b *Buffer) Initialize(config *config.Config) error {
	b.inputChan = make(chan envelope.KotaEnvelope, 20000) // TODO -> will this overflow if the persistence to disk is synchronous? Should it be larger? Smaller? idk, think about this later.
	b.shutdown = make(chan int, 1)
	b.config = config
	b.bufferStart = time.Now()
	go func(envelope <-chan envelope.KotaEnvelope, shutdown <-chan int) {
		for {
			select {
			case envelope := <-envelope:
				b.envelopes = append(b.envelopes, envelope) // Fully rewriting on each append is nawwwwt ideal. FIXME.
				b.bufferRecords += 1
				if b.bufferRecords >= b.config.Records {
					b.Purge()
				}
				// TODO -> Add buffer size-based purging
				// TODO -> Add buffer time-based purging
				util.Pprint(b.envelopes) // TODO -> persist to disk, flush, whatever
			case <-shutdown:
				log.Debug().Msg("shutting down buffer")
				// TODO -> do something that is safe on shutdown
				log.Debug().Msg("buffer shut down")
				return
			}
		}
	}(b.inputChan, b.shutdown)
	return nil
}

func (b *Buffer) Append(envelopes []envelope.KotaEnvelope) error {
	for _, envelope := range envelopes {
		log.Debug().Msg("appending envelope to buffer")
		b.inputChan <- envelope
	}
	return nil
}

func (b *Buffer) Purge() error {
	log.Debug().Msg("purging buffer")
	// Actually do something here.
	b.envelopes = []envelope.KotaEnvelope{}
	b.bufferStart = time.Now()
	b.bufferRecords = 0
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
