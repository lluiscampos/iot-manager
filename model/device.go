// Copyright 2021 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package model

import "github.com/google/uuid"

// Device represents a device and its available integrations
type Device struct {
	// ID is the device ID
	ID string `json:"id" bson:"_id"`
	// Integrations contains the list of integrations for this device
	IntegrationIDs []uuid.UUID `json:"integration_ids" bson:"integration_ids"`
}

// DeviceIntegration stores the integration details for a specific device
type DeviceIntegration struct {
	// ID is the integration ID
	ID string `json:"id" bson:"id"`
}
