package beta

import (
	"encoding/json"
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementSettingInstance = DeviceManagementCollectionSettingInstance{}

type DeviceManagementCollectionSettingInstance struct {
	// The collection of values
	Value *[]DeviceManagementSettingInstance `json:"value,omitempty"`

	// Fields inherited from DeviceManagementSettingInstance

	// The ID of the setting definition for this instance
	DefinitionId *string `json:"definitionId,omitempty"`

	// JSON representation of the value
	ValueJson nullable.Type[string] `json:"valueJson,omitempty"`

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

func (s DeviceManagementCollectionSettingInstance) DeviceManagementSettingInstance() BaseDeviceManagementSettingInstanceImpl {
	return BaseDeviceManagementSettingInstanceImpl{
		DefinitionId: s.DefinitionId,
		ValueJson:    s.ValueJson,
		Id:           s.Id,
		ODataId:      s.ODataId,
		ODataType:    s.ODataType,
	}
}

func (s DeviceManagementCollectionSettingInstance) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementCollectionSettingInstance{}

func (s DeviceManagementCollectionSettingInstance) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementCollectionSettingInstance
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementCollectionSettingInstance: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementCollectionSettingInstance: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementCollectionSettingInstance"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementCollectionSettingInstance: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DeviceManagementCollectionSettingInstance{}

func (s *DeviceManagementCollectionSettingInstance) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DefinitionId *string               `json:"definitionId,omitempty"`
		ValueJson    nullable.Type[string] `json:"valueJson,omitempty"`
		Id           *string               `json:"id,omitempty"`
		ODataId      *string               `json:"@odata.id,omitempty"`
		ODataType    *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DefinitionId = decoded.DefinitionId
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.ValueJson = decoded.ValueJson

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceManagementCollectionSettingInstance into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["value"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Value into list []json.RawMessage: %+v", err)
		}

		output := make([]DeviceManagementSettingInstance, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDeviceManagementSettingInstanceImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Value' for 'DeviceManagementCollectionSettingInstance': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Value = &output
	}

	return nil
}
