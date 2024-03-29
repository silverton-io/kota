// Copyright (c) 2024 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the Apache-2.0 license, a copy of
// which may be found at https://github.com/silverton-io/kota/blob/main/LICENSE

package config

type App struct {
	Version       string `json:"version"`
	Name          string `json:"name"`
	Env           string `json:"env"`
	Port          string `json:"port"`
	TrackerDomain string `json:"trackerDomain"`
}
