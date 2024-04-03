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
	config              *config.Config
	bufferRecords       int
	bufferFirstAppended time.Time
	inputChan           chan envelope.KotaEnvelope
	envelopes           []envelope.KotaEnvelope
	shutdown            chan int
}

func (b *Buffer) Initialize(config *config.Config) error {
	b.inputChan = make(chan envelope.KotaEnvelope, 20000) // TODO -> will this overflow if the persistence to disk is synchronous? Should it be larger? Smaller? idk, think about this later.
	b.shutdown = make(chan int, 1)
	b.config = config

	ticker := time.NewTicker(time.Duration(b.config.Time) * time.Second)

	// Kick off
	go func(envelope <-chan envelope.KotaEnvelope, shutdown <-chan int) {
		for {
			select {
			case <-ticker.C:
				if !b.bufferFirstAppended.IsZero() && time.Since(b.bufferFirstAppended) > time.Duration(b.config.Time)*time.Second {
					log.Info().Msg("buffer reached max time, purging")
					b.Purge()
				}
			case envelope := <-envelope:
				b.envelopes = append(b.envelopes, envelope) // Fully rewriting on each append is nawwwwt ideal. FIXME.
				// TODO -> If buffer durability is enabled also write this to a local file
				b.bufferRecords += 1
				if b.bufferRecords == 1 {
					log.Debug().Msg("setting buffer first appended time")
					b.bufferFirstAppended = time.Now()
				}
				// Purge the buffer when it reaches the maximum number of records
				if b.bufferRecords >= b.config.Records {
					log.Info().Msg("buffer reached max records, purging")
					b.Purge()
				}
				// util.Pprint(b.envelopes) // TODO -> persist to disk, flush, whatever
			case <-shutdown:
				log.Debug().Msg("shutting down buffer")
				ticker.Stop()
				b.Purge()
				log.Debug().Msg("buffer shut down")
				return
			}
		}
	}(b.inputChan, b.shutdown)

	// TODO -> If any prior buffer files are found, load them into the new buffer and remove the file

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
	util.Pprint(b.envelopes)
	b.envelopes = []envelope.KotaEnvelope{}
	b.bufferRecords = 0
	b.bufferFirstAppended = time.Time{}
	// TODO -> If buffer durability is enabled also remove the local file
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
