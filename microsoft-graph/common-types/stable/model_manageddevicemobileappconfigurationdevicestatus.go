package stable

import (
	"encoding/json"
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedDeviceMobileAppConfigurationDeviceStatus{}

type ManagedDeviceMobileAppConfigurationDeviceStatus struct {
	// The DateTime when device compliance grace period expires
	ComplianceGracePeriodExpirationDateTime *string `json:"complianceGracePeriodExpirationDateTime,omitempty"`

	// Device name of the DevicePolicyStatus.
	DeviceDisplayName nullable.Type[string] `json:"deviceDisplayName,omitempty"`

	// The device model that is being reported
	DeviceModel nullable.Type[string] `json:"deviceModel,omitempty"`

	// Last modified date time of the policy report.
	LastReportedDateTime *string `json:"lastReportedDateTime,omitempty"`

	Status *ComplianceStatus `json:"status,omitempty"`

	// The User Name that is being reported
	UserName nullable.Type[string] `json:"userName,omitempty"`

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

func (s ManagedDeviceMobileAppConfigurationDeviceStatus) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedDeviceMobileAppConfigurationDeviceStatus{}

func (s ManagedDeviceMobileAppConfigurationDeviceStatus) MarshalJSON() ([]byte, error) {
	type wrapper ManagedDeviceMobileAppConfigurationDeviceStatus
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedDeviceMobileAppConfigurationDeviceStatus: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedDeviceMobileAppConfigurationDeviceStatus: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedDeviceMobileAppConfigurationDeviceStatus"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedDeviceMobileAppConfigurationDeviceStatus: %+v", err)
	}

	return encoded, nil
}
