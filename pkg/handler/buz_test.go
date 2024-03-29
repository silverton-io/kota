// Copyright (c) 2024 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the Apache-2.0 license, a copy of
// which may be found at https://github.com/silverton-io/kota/blob/main/LICENSE

package handler

import (
	"io"
	"net/http"
	"testing"

	testutil "github.com/silverton-io/buz/pkg/testUtil"
	"github.com/stretchr/testify/assert"
)

func TestBuzHandler(t *testing.T) {
	srv := testutil.BuildTestServer(BuzHandler())

	resp, _ := http.Get(srv.URL + testutil.URL)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf(`got status code %v, want %v`, resp.StatusCode, http.StatusOK)
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "🐝", string(b))
}
