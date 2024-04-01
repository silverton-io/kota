// Copyright (c) 2024 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the Apache-2.0 license, a copy of
// which may be found at https://github.com/silverton-io/kota/blob/main/LICENSE

package envelope

import (
	"time"

	"github.com/google/uuid"
	"github.com/silverton.io/kota/pkg/config"
)

func BuildEnvelopesFromRequest(conf config.App) []KotaEnvelope {
	fakeEnvelope := KotaEnvelope{
		Uuid:        uuid.New(),
		Timestamp:   time.Now(),
		KotaName:    conf.Name,
		KotaVersion: conf.Version,
		KotaEnv:     conf.Env,
		Payload:     map[string]interface{}{},
	}
	return []KotaEnvelope{fakeEnvelope}
}
