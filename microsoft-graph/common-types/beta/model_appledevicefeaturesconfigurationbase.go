package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleDeviceFeaturesConfigurationBase interface {
	Entity
	DeviceConfiguration
	AppleDeviceFeaturesConfigurationBase() BaseAppleDeviceFeaturesConfigurationBaseImpl
}

var _ AppleDeviceFeaturesConfigurationBase = BaseAppleDeviceFeaturesConfigurationBaseImpl{}

type BaseAppleDeviceFeaturesConfigurationBaseImpl struct {
	// An array of AirPrint printers that should always be shown. This collection can contain a maximum of 500 elements.
	AirPrintDestinations *[]AirPrintDestination `json:"airPrintDestinations,omitempty"`

	// Fields inherited from DeviceConfiguration

	// The list of assignments for the device configuration profile.
	Assignments *[]DeviceConfigurationAssignment `json:"assignments,omitempty"`

	// DateTime the object was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Admin provided description of the Device Configuration.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The device mode applicability rule for this Policy.
	DeviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`

	// The OS edition applicability for this Policy.
	DeviceManagementApplicabilityRuleOsEdition *DeviceManagementApplicabilityRuleOsEdition `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`

	// The OS version applicability rule for this Policy.
	DeviceManagementApplicabilityRuleOsVersion *DeviceManagementApplicabilityRuleOsVersion `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`

	// Device Configuration Setting State Device Summary
	DeviceSettingStateSummaries *[]SettingStateDeviceSummary `json:"deviceSettingStateSummaries,omitempty"`

	// Device Configuration devices status overview
	DeviceStatusOverview *DeviceConfigurationDeviceOverview `json:"deviceStatusOverview,omitempty"`

	// Device configuration installation status by device.
	DeviceStatuses *[]DeviceConfigurationDeviceStatus `json:"deviceStatuses,omitempty"`

	// Admin provided name of the device configuration.
	DisplayName *string `json:"displayName,omitempty"`

	// The list of group assignments for the device configuration profile.
	GroupAssignments *[]DeviceConfigurationGroupAssignment `json:"groupAssignments,omitempty"`

	// DateTime the object was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// List of Scope Tags for this Entity instance.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// Indicates whether or not the underlying Device Configuration supports the assignment of scope tags. Assigning to the
	// ScopeTags property is not allowed when this value is false and entities will not be visible to scoped users. This
	// occurs for Legacy policies created in Silverlight and can be resolved by deleting and recreating the policy in the
	// Azure Portal. This property is read-only.
	SupportsScopeTags *bool `json:"supportsScopeTags,omitempty"`

	// Device Configuration users status overview
	UserStatusOverview *DeviceConfigurationUserOverview `json:"userStatusOverview,omitempty"`

	// Device configuration installation status by user.
	UserStatuses *[]DeviceConfigurationUserStatus `json:"userStatuses,omitempty"`

	// Version of the device configuration.
	Version *int64 `json:"version,omitempty"`

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

func (s BaseAppleDeviceFeaturesConfigurationBaseImpl) AppleDeviceFeaturesConfigurationBase() BaseAppleDeviceFeaturesConfigurationBaseImpl {
	return s
}

func (s BaseAppleDeviceFeaturesConfigurationBaseImpl) DeviceConfiguration() BaseDeviceConfigurationImpl {
	return BaseDeviceConfigurationImpl{
		Assignments:     s.Assignments,
		CreatedDateTime: s.CreatedDateTime,
		Description:     s.Description,
		DeviceManagementApplicabilityRuleDeviceMode: s.DeviceManagementApplicabilityRuleDeviceMode,
		DeviceManagementApplicabilityRuleOsEdition:  s.DeviceManagementApplicabilityRuleOsEdition,
		DeviceManagementApplicabilityRuleOsVersion:  s.DeviceManagementApplicabilityRuleOsVersion,
		DeviceSettingStateSummaries:                 s.DeviceSettingStateSummaries,
		DeviceStatusOverview:                        s.DeviceStatusOverview,
		DeviceStatuses:                              s.DeviceStatuses,
		DisplayName:                                 s.DisplayName,
		GroupAssignments:                            s.GroupAssignments,
		LastModifiedDateTime:                        s.LastModifiedDateTime,
		RoleScopeTagIds:                             s.RoleScopeTagIds,
		SupportsScopeTags:                           s.SupportsScopeTags,
		UserStatusOverview:                          s.UserStatusOverview,
		UserStatuses:                                s.UserStatuses,
		Version:                                     s.Version,
		Id:                                          s.Id,
		ODataId:                                     s.ODataId,
		ODataType:                                   s.ODataType,
	}
}

func (s BaseAppleDeviceFeaturesConfigurationBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ AppleDeviceFeaturesConfigurationBase = RawAppleDeviceFeaturesConfigurationBaseImpl{}

// RawAppleDeviceFeaturesConfigurationBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAppleDeviceFeaturesConfigurationBaseImpl struct {
	appleDeviceFeaturesConfigurationBase BaseAppleDeviceFeaturesConfigurationBaseImpl
	Type                                 string
	Values                               map[string]interface{}
}

func (s RawAppleDeviceFeaturesConfigurationBaseImpl) AppleDeviceFeaturesConfigurationBase() BaseAppleDeviceFeaturesConfigurationBaseImpl {
	return s.appleDeviceFeaturesConfigurationBase
}

func (s RawAppleDeviceFeaturesConfigurationBaseImpl) DeviceConfiguration() BaseDeviceConfigurationImpl {
	return s.appleDeviceFeaturesConfigurationBase.DeviceConfiguration()
}

func (s RawAppleDeviceFeaturesConfigurationBaseImpl) Entity() BaseEntityImpl {
	return s.appleDeviceFeaturesConfigurationBase.Entity()
}

var _ json.Marshaler = BaseAppleDeviceFeaturesConfigurationBaseImpl{}

func (s BaseAppleDeviceFeaturesConfigurationBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseAppleDeviceFeaturesConfigurationBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseAppleDeviceFeaturesConfigurationBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseAppleDeviceFeaturesConfigurationBaseImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.appleDeviceFeaturesConfigurationBase"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseAppleDeviceFeaturesConfigurationBaseImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalAppleDeviceFeaturesConfigurationBaseImplementation(input []byte) (AppleDeviceFeaturesConfigurationBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AppleDeviceFeaturesConfigurationBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.iosDeviceFeaturesConfiguration") {
		var out IosDeviceFeaturesConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosDeviceFeaturesConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSDeviceFeaturesConfiguration") {
		var out MacOSDeviceFeaturesConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSDeviceFeaturesConfiguration: %+v", err)
		}
		return out, nil
	}

	var parent BaseAppleDeviceFeaturesConfigurationBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAppleDeviceFeaturesConfigurationBaseImpl: %+v", err)
	}

	return RawAppleDeviceFeaturesConfigurationBaseImpl{
		appleDeviceFeaturesConfigurationBase: parent,
		Type:                                 value,
		Values:                               temp,
	}, nil

}
