// Copyright (c) 2024 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the Apache-2.0 license, a copy of
// which may be found at https://github.com/silverton-io/kota/blob/main/LICENSE

package destination

import (
	"github.com/silverton.io/kota/pkg/config"
	"github.com/silverton.io/kota/pkg/envelope"
)

type Destination interface {
	Metadata() map[string]string // FIXME - don't use map[string]interface
	Initialize(config config.Destination) error
	Send(batch []envelope.KotaEnvelope) error
}
