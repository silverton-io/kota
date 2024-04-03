// Copyright (c) 2024 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the Apache-2.0 license, a copy of
// which may be found at https://github.com/silverton-io/kota/blob/main/LICENSE

package consumer

import (
	"github.com/silverton.io/kota/pkg/buffer"
	"github.com/silverton.io/kota/pkg/config"
)

type Consumer interface {
	// Metadata() map[string]interface{} // FIXME -> Get better consumer statistics
	Initialize(config config.Input, buffer *buffer.Buffer) error
	Consume() error
	Shutdown() error
}
