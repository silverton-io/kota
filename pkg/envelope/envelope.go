// Copyright (c) 2024 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the Apache-2.0 license, a copy of
// which may be found at https://github.com/silverton-io/kota/blob/main/LICENSE

package envelope

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Payload map[string]interface{}

type KotaEnvelope struct {
	Uuid        uuid.UUID `json:"uuid"`
	Timestamp   time.Time `json:"timestamp"`
	Payload     Payload   `json:"payload"`
	KotaName    string    `json:"kotaName"`
	KotaVersion string    `json:"kotaVersion"`
	KotaEnv     string    `json:"kotaEnv"`
	Source      string    `json:"source"`
}

func (e *KotaEnvelope) AsMap() (map[string]interface{}, error) {
	var m map[string]interface{}
	marshaledEnvelope, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(marshaledEnvelope, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func (e *KotaEnvelope) AsByte() ([]byte, error) {
	eBytes, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	return eBytes, nil
}
