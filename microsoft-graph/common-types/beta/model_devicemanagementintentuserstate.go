package beta

import (
	"encoding/json"
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceManagementIntentUserState{}

type DeviceManagementIntentUserState struct {
	// Count of Devices that belongs to a user for an intent
	DeviceCount *int64 `json:"deviceCount,omitempty"`

	// Last modified date time of an intent report
	LastReportedDateTime *string `json:"lastReportedDateTime,omitempty"`

	State *ComplianceStatus `json:"state,omitempty"`

	// The user name that is being reported on a device
	UserName nullable.Type[string] `json:"userName,omitempty"`

	// The user principal name that is being reported on a device
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s DeviceManagementIntentUserState) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementIntentUserState{}

func (s DeviceManagementIntentUserState) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementIntentUserState
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementIntentUserState: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementIntentUserState: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementIntentUserState"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementIntentUserState: %+v", err)
	}

	return encoded, nil
}
