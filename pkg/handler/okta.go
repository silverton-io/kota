// Copyright (c) 2024 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the Apache-2.0 license, a copy of
// which may be found at https://github.com/silverton-io/kota/blob/main/LICENSE

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silverton.io/kota/pkg/constants"
	"github.com/silverton.io/kota/pkg/response"
)

type OktaHookResponse struct {
	Message      string  `json:"message"`
	Verification *string `json:"verification,omitempty"`
}

func OktaHookHandler(c *gin.Context) {
	var verification *string
	verification_header := c.GetHeader(constants.OKTA_HOOK_VERIFICATION_HEADER)
	if verification_header != "" {
		verification = &verification_header
	}
	response := OktaHookResponse{
		Message:      response.Ok.Message,
		Verification: verification,
	}
	c.JSON(http.StatusOK, response)
}
