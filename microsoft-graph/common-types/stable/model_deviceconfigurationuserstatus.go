package stable

import (
	"encoding/json"
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceConfigurationUserStatus{}

type DeviceConfigurationUserStatus struct {
	// Devices count for that user.
	DevicesCount *int64 `json:"devicesCount,omitempty"`

	// Last modified date time of the policy report.
	LastReportedDateTime *string `json:"lastReportedDateTime,omitempty"`

	Status *ComplianceStatus `json:"status,omitempty"`

	// User name of the DevicePolicyStatus.
	UserDisplayName nullable.Type[string] `json:"userDisplayName,omitempty"`

	// UserPrincipalName.
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

func (s DeviceConfigurationUserStatus) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceConfigurationUserStatus{}

func (s DeviceConfigurationUserStatus) MarshalJSON() ([]byte, error) {
	type wrapper DeviceConfigurationUserStatus
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceConfigurationUserStatus: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceConfigurationUserStatus: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceConfigurationUserStatus"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceConfigurationUserStatus: %+v", err)
	}

	return encoded, nil
}
