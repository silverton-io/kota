// Copyright (c) 2024 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the Apache-2.0 license, a copy of
// which may be found at https://github.com/silverton-io/kota/blob/main/LICENSE

package config

type Buffer struct {
	Records    int `json:"records"`
	Time       int `json:"time"`
	Size       int `json:"size"`
	Durability `json:"durability"`
}

type Durability struct {
	Enabled   bool   `json:"durability"`
	Directory string `json:"directory"`
}
