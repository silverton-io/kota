// Copyright (c) 2024 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the Apache-2.0 license, a copy of
// which may be found at https://github.com/silverton-io/kota/blob/main/LICENSE

package handler

import (
	"github.com/gin-gonic/gin"
)

func BuzHandler() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		c.String(200, "🐝")
	}
	return gin.HandlerFunc(fn)
}
