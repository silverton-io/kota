// Copyright (c) 2024 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the Apache-2.0 license, a copy of
// which may be found at https://github.com/silverton-io/kota/blob/main/LICENSE

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silverton.io/kota/pkg/response"
)

func ArrowFlightHandler(c *gin.Context) {
	// TODO -> Implement me
	// NOTES -> // https://blog.djnavarro.net/posts/2022-10-18_arrow-flight/, https://voltrondata.com/blog/data-transfer-with-apache-arrow-and-golang
	c.JSON(http.StatusOK, response.Ok)
}
