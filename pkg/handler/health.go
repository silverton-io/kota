// Copyright (c) 2024 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the Apache-2.0 license, a copy of
// which may be found at https://github.com/silverton-io/kota/blob/main/LICENSE

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"silverton.io/kota/pkg/response"
)

func HealthcheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, response.Ok)
}
