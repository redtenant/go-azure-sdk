package beta

import (
	"encoding/json"
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedTenantsAppPerformance{}

type ManagedTenantsAppPerformance struct {
	AppFriendlyName            nullable.Type[string] `json:"appFriendlyName,omitempty"`
	AppName                    nullable.Type[string] `json:"appName,omitempty"`
	AppPublisher               nullable.Type[string] `json:"appPublisher,omitempty"`
	LastUpdatedDateTime        nullable.Type[string] `json:"lastUpdatedDateTime,omitempty"`
	MeanTimeToFailureInMinutes nullable.Type[int64]  `json:"meanTimeToFailureInMinutes,omitempty"`
	TenantDisplayName          nullable.Type[string] `json:"tenantDisplayName,omitempty"`
	TenantId                   nullable.Type[string] `json:"tenantId,omitempty"`
	TotalActiveDeviceCount     nullable.Type[int64]  `json:"totalActiveDeviceCount,omitempty"`
	TotalAppCrashCount         nullable.Type[int64]  `json:"totalAppCrashCount,omitempty"`
	TotalAppFreezeCount        nullable.Type[int64]  `json:"totalAppFreezeCount,omitempty"`

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

func (s ManagedTenantsAppPerformance) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsAppPerformance{}

func (s ManagedTenantsAppPerformance) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsAppPerformance
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsAppPerformance: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsAppPerformance: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedTenants.appPerformance"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsAppPerformance: %+v", err)
	}

	return encoded, nil
}
