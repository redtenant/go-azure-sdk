package beta

import (
	"encoding/json"
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IosVppAppAssignedLicense = IosVppAppAssignedDeviceLicense{}

type IosVppAppAssignedDeviceLicense struct {
	// The device name.
	DeviceName nullable.Type[string] `json:"deviceName,omitempty"`

	// The managed device ID.
	ManagedDeviceId nullable.Type[string] `json:"managedDeviceId,omitempty"`

	// Fields inherited from IosVppAppAssignedLicense

	// The user email address.
	UserEmailAddress nullable.Type[string] `json:"userEmailAddress,omitempty"`

	// The user ID.
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// The user name.
	UserName nullable.Type[string] `json:"userName,omitempty"`

	// The user principal name.
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

func (s IosVppAppAssignedDeviceLicense) IosVppAppAssignedLicense() BaseIosVppAppAssignedLicenseImpl {
	return BaseIosVppAppAssignedLicenseImpl{
		UserEmailAddress:  s.UserEmailAddress,
		UserId:            s.UserId,
		UserName:          s.UserName,
		UserPrincipalName: s.UserPrincipalName,
		Id:                s.Id,
		ODataId:           s.ODataId,
		ODataType:         s.ODataType,
	}
}

func (s IosVppAppAssignedDeviceLicense) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IosVppAppAssignedDeviceLicense{}

func (s IosVppAppAssignedDeviceLicense) MarshalJSON() ([]byte, error) {
	type wrapper IosVppAppAssignedDeviceLicense
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IosVppAppAssignedDeviceLicense: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IosVppAppAssignedDeviceLicense: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.iosVppAppAssignedDeviceLicense"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IosVppAppAssignedDeviceLicense: %+v", err)
	}

	return encoded, nil
}
