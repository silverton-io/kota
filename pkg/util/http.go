// Copyright (c) 2024 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the Apache-2.0 license, a copy of
// which may be found at https://github.com/silverton-io/kota/blob/main/LICENSE

package util

import (
	"github.com/gin-gonic/gin"
)

// HttpHeadersToMap returns a map of http headers,
// but treats single-element headers as strings instead
// of slices of length 1
func HttpHeadersToMap(c *gin.Context) map[string]interface{} {
	headers := make(map[string]interface{})
	for k, v := range c.Request.Header {
		if len(v) == 1 {
			headers[k] = v[0]
		} else {
			headers[k] = v
		}
	}
	return headers
}
