// Copyright (c) 2024 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the Apache-2.0 license, a copy of
// which may be found at https://github.com/silverton-io/kota/blob/main/LICENSE

package response

type Response struct {
	Message string `json:"message"`
}

var Ok = Response{
	Message: "ok",
}

var NotImplemented = Response{
	Message: "not implemented",
}

var InvalidContentType = Response{
	Message: "invalid content type",
}

var BadRequest = Response{
	Message: "bad request",
}

var Timeout = Response{
	Message: "request timed out",
}

var RateLimitExceeded = Response{
	Message: "rate limit exceeded",
}

var ManifoldDistributionError = Response{
	Message: "distribution error",
}

var MissingAuthHeader = Response{
	Message: "missing authorization header",
}

var MissingAuthSchemeOrToken = Response{
	Message: "missing auth scheme or token",
}

var InvalidAuthScheme = Response{
	Message: "invalid scheme",
}

var InvalidAuthToken = Response{
	Message: "invalid token",
}
