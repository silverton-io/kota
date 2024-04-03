// Copyright (c) 2024 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the Apache-2.0 license, a copy of
// which may be found at https://github.com/silverton-io/kota/blob/main/LICENSE

package constants

import "time"

// App
const KOTA = "kota"

// HTTP
const DEFAULT_HTTP_TIMEOUT = time.Duration(5) * time.Second
const DEFAULT_SHUTDOWN_TIMEOUT = time.Duration(15) * time.Second

// Default Routes
const DEFAULT_HEALTH_ROUTE = "/health"
const DEFAULT_FLIGHT_ROUTE = "/flight"
const DEFAULT_OKTA_HOOKS_ROUTE = "/okta"
const DEFAULT_SPLUNK_HEC_ROUTE = "/hec"
