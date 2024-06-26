// Copyright (c) 2024 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the Apache-2.0 license, a copy of
// which may be found at https://github.com/silverton-io/kota/blob/main/LICENSE

package constants

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigPath(t *testing.T) {
	assert.Equal(t, "KOTA_CONFIG_PATH", KOTA_CONFIG_PATH)
	assert.Equal(t, "DEBUG", DEBUG)
}
