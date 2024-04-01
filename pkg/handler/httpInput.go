// Copyright (c) 2024 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the Apache-2.0 license, a copy of
// which may be found at https://github.com/silverton-io/kota/blob/main/LICENSE

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silverton.io/kota/pkg/buffer"
	"github.com/silverton.io/kota/pkg/config"
	"github.com/silverton.io/kota/pkg/constants"
	"github.com/silverton.io/kota/pkg/envelope"
	"github.com/silverton.io/kota/pkg/response"
)

type ValidatedResponse struct {
	Message      string  `json:"messsage"`
	Verification *string `json:"verification,omitempty"`
}

func HttpInputHandler(b buffer.Buffer, conf config.App) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		// Get the verification token if it is sent
		var verification *string
		verification_header := c.GetHeader(constants.OKTA_HOOK_VERIFICATION_HEADER)
		if verification_header != "" {
			verification = &verification_header
		}
		resp := ValidatedResponse{
			Message:      response.Ok.Message,
			Verification: verification,
		}
		// TODO -> Explicitly check for json content type and whatnot
		envelopes := envelope.BuildEnvelopesFromRequest(conf)
		err := b.Append(envelopes)

		if err != nil {
			c.Header("Retry-After", response.RETRY_AFTER_60)
			c.JSON(http.StatusTooManyRequests, response.ManifoldDistributionError)
		} else {
			c.JSON(http.StatusOK, resp)
		}

	}
	return gin.HandlerFunc(fn)
}
